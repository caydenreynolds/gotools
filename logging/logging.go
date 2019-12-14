package logging

import (
	"log"
	"os"
)

var logFile *os.File

func StartLogger(path string) {
	if logFile != nil {
		Error("Cannot start a new logger. Instance already running")
	}

	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 7777)
	if err != nil {
		panic("Cannot open log file!")
	}
	log.SetOutput(logFile)
}

func CloseLogger() {
	if logFile == nil {
		panic("Cannot close logger. No logger instance is running")
	}
	err := logFile.Close()
	if err != nil {
		panic(err)
	}
	logFile = nil
}

func CheckFatal(e error) {
	if e != nil {
		log.Fatalf("Fatal: %v", e)
	}
}

func Fatal(message string) {
	log.Println("Fatal: " + message)
}

func CheckError(e error) {
	if e != nil {
		log.Fatalf("Error: %v", e)
	}
}

func Error(message string) {
	log.Println("Error: " + message)
}

func Info(message string) {
	log.Println("Info: " + message)
}
