package common_utils

import (
	"log"
	"os"
	"path/filepath"
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	debugLogger   *log.Logger
)

func InitLogger() {
	logDir := "log"

	// Create the log directory if it doesn't exist
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatal("Error creating log directory:", err)
	}

	infoLogFile, err := os.OpenFile(filepath.Join(logDir, "info.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening info.log file:", err)
	}

	warningLogFile, err := os.OpenFile(filepath.Join(logDir, "warning.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening warning.log file:", err)
	}

	errorLogFile, err := os.OpenFile(filepath.Join(logDir, "error.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening error.log file:", err)
	}

	debugLogFile, err := os.OpenFile(filepath.Join(logDir, "debug.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening debug.log file:", err)
	}

	infoLogger = log.New(infoLogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(warningLogFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(errorLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(debugLogFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string) {
	infoLogger.Println(message)
}

func Warning(message string) {
	warningLogger.Println(message)
}

func Error(message string) {
	errorLogger.Println(message)
}

func Debug(message string) {
	debugLogger.Println(message)
}
