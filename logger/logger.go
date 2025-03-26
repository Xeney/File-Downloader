package logger

import (
	"log"
	"os"
)

// Добавляем в начало файла
var (
	logFile *os.File
	logger  *log.Logger
)

// Инициализация логгера (вызовите это в начале main())
func InitLogger(logPath string) error {
	var err error
	logFile, err = os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	logger = log.New(logFile, "", log.LstdFlags)
	return nil
}

// Простая функция логирования
func LogMessage(message string) {
	if logger != nil {
		logger.Println(message)
	}
}
