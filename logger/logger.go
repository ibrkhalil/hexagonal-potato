package logger

import (
	"log"
	"os"
)

var AppLogger *log.Logger

func Init(logFile string) error {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	AppLogger = log.New(file, "APP_LOG: ", log.LstdFlags|log.Lshortfile)
	return nil
}
