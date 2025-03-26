package main

import (
	"file_downloader/app"
	"file_downloader/config"
	"file_downloader/logger"
	"log"
)

func main() {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal("Ошибка инициализации логгера:", err)
	}

	err = logger.InitLogger("downloads.log")
	if err != nil {
		log.Fatal(err)
	}

	downloader := app.NewDownloader(*config)

	tasks := []app.DownloadTask{
		{"https://images.wallpaperscraft.ru/image/single/listia_asfalt_osen_1382185_3840x2160.jpg", "фото-1.jpg"},
		{"https://images.wallpaperscraft.ru/image/single/listia_asfalt_osen_1382185_3840x2160.jpg", "фото-2.jpg"},
		{"https://images.wallpaperscraft.ru/image/single/listia_asfalt_osen_1382185_3840x2160.jpg", "фото-3.jpg"},
		{"https://images.wallpapercraft.ru/image/single/listia_asfalt_osen_1382185_3840x2160.jpg", "фото-4.jpg"},
		{"https://images.wallpaperscraft.ru/image/single/listia_asfalt_osen_1382185_3840x2160.jpg", "фото-5.jpg"},
		{"https://images.wallpaperscraft.ru/image/single/listia_asfalt_osen_1382185_3840x2160.jpg", "фото-6.jpg"},
		{"https://images.wallpaperscraft.ru/image/single/listia_asfalt_osen_1382185_3840x2160.jpg", "фото-7.jpg"},
	}

	go downloader.TrackProgress(len(tasks))
	downloader.WorkerPool(tasks)
}
