package logging

import (
	"log"
	"os"
)

//Singleton pattern logging
var logFile *os.File

func StartLogger(path string) {
	if logFile != nil {
		Error("Cannot start a new logger. Instance already running")
	}

	logFile, err := os.Create(path)
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
		log.Fatalf("Fatal: %v\n", e)
	}
}

func CheckFatalMessage(e error, message string) {
	if e != nil {
		log.Fatalf("Fatal: %v. %s\n", e, message)
	}
}

func Fatal(message string) {
	log.Fatal("Fatal: " + message)
}

func CheckError(e error) bool {
	if e != nil {
		log.Printf("Error: %v\n", e)
	}
	return e != nil
}

func CheckErrorMessage(e error, message string) bool {
	if e != nil {
		log.Printf("Error: %v. %s\n", e, message)
	}
	return e != nil
}

func Error(message string) {
	log.Println("Error: " + message)
}

func Info(message string) {
	log.Println("Info: " + message)
}
