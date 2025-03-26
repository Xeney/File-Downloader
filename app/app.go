package app

import (
	"file_downloader/config"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type DownloadTask struct {
	URL      string
	FilePath string
}

type Downloader struct {
	MaxThreads int
	Timeout    time.Duration
	Client     *http.Client
	ProgressCh chan int
	activeTask sync.Map
}

func NewDownloader(config config.Config) *Downloader {
	return &Downloader{
		MaxThreads: config.MaxThreads,
		Timeout:    time.Duration(config.MaxThreads) * time.Second,
		Client: &http.Client{
			Timeout: time.Duration(config.TimeoutSeconds) * time.Second,
		},
		ProgressCh: make(chan int, 100),
	}
}

func (d *Downloader) DownloadFile(task DownloadTask) error {
	resp, err := d.Client.Get(task.URL)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned: %v", err)
	}

	file, err := os.Create("downloads/" + task.FilePath)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func (d *Downloader) WorkerPool(tasks []DownloadTask) error {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, d.MaxThreads)

	for _, task := range tasks {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(t DownloadTask) {
			defer wg.Done()
			defer func() { <-semaphore }()

			err := d.DownloadFile(t)
			if err != nil {
				fmt.Printf("Ошибка загрузки: %v\n", err)
			} else {
				d.ProgressCh <- 1
			}
		}(task)
	}

	wg.Wait()                          // Ждём завершения всех загрузок
	time.Sleep(100 * time.Millisecond) // Даём время на обработку последнего значения
	close(d.ProgressCh)
	return nil
}

func (d *Downloader) TrackProgress(total int) {
	completed := 0
	for range d.ProgressCh {
		completed++
		fmt.Printf("\rПрогресс: %d/%d (%.1f%%)", completed, total, float64(completed)/float64(total)*100)
	}
	fmt.Println("\nЗагрузка завершена!")
}
