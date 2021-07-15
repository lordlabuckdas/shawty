package utils

import (
	"log"
	"os"

	"github.com/lordlabuckdas/shawty/config"
)

var (
	InfoLogger *log.Logger
	ErrorLogger *log.Logger
)

func LoggerInit() error {
	lf, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return err
	}
	InfoLogger = log.New(lf, "[INFO] ", log.Ldate|log.Ltime|log.Llongfile)
	ErrorLogger = log.New(lf, "[ERROR] ", log.Ldate|log.Ltime|log.Llongfile)
	return nil
}
