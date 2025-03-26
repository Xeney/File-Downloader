package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	MaxThreads     int    `json:"max_threads"`
	DownloadDir    string `json:"download_dir"`
	TimeoutSeconds int    `json:"timeout_seconds"`
	EnableLogging  bool   `json:"enable_logging"`
	LogFile        string `json:"log_file"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %v", err)
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	// Валидация
	if config.MaxThreads <= 0 {
		return nil, fmt.Errorf("max_threads должен быть > 0")
	}
	if config.TimeoutSeconds <= 0 {
		return nil, fmt.Errorf("timeout_seconds должен быть > 0")
	}

	// Создание директории для загрузок, если её нет
	if err := os.MkdirAll(config.DownloadDir, 0755); err != nil {
		return nil, fmt.Errorf("не удалось создать директорию: %v", err)
	}

	return &config, nil
}
