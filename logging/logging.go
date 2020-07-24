package logging

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"
)

type LogLevel string

const (
	INFO  LogLevel = "INFO"
	ERROR LogLevel = "ERROR"
	FATAL LogLevel = "FATAL"
)

//Singleton pattern logging
var singletonLogDir *string
var singletonLogFile *os.File
var singletonNewFileTicker *time.Ticker
var singletonClose chan bool
var singletonFileLock = sync.Mutex{}

func StartLogger(logDir *string) error {
	if singletonLogFile != nil {
		Error("Cannot start a new logger. Instance already running")
		return errors.New("cannot start a new logger! Instance already running") // Hopefully, this will never happen
	} else {
		// Open a logfile
		singletonLogDir = logDir
		var err error
		singletonLogFile, err = ioutil.TempFile("", "goLoggingStartupFile")
		if err != nil {
			return err
		}
		tempLogFile := *singletonLogFile
		openFile()
		if err = tempLogFile.Close(); err == nil { // If the file has not been closed, we couldn't open a new logfile
			return errors.New("could not open new file for logging")
		}
		if err = os.Remove(tempLogFile.Name()); err != nil {
			return err
		}

		// Change the file every day at midnight, so we store a month's worth of logs
		singletonClose = make(chan bool)
		year, month, day := time.Now().UTC().Date()
		firstFileOpen := time.Date(year, month, day, 0, 1, 0, 0, time.UTC)
		firstFileOpen = firstFileOpen.AddDate(0, 0, 1)
		time.AfterFunc(firstFileOpen.Sub(time.Now().UTC()), openNewFileDaily)
		return nil
	}
}

func CloseLogger() {
	if singletonLogFile == nil {
		panic("Cannot close logger. No logger instance is running")
	}
	singletonFileLock.Lock()
	defer singletonFileLock.Unlock()
	singletonClose <- true // Shut down our logging goroutines
	close(singletonClose)
	if singletonNewFileTicker != nil {
		singletonNewFileTicker.Stop()
	}
	err := singletonLogFile.Close()
	if err != nil {
		panic(err)
	}
	singletonLogFile = nil
}

//We ignore errors in this function...hopefully they've been handled before now
func writeToFile(level LogLevel, message string, values ...interface{}) {
	now := time.Now().UTC()
	message = fmt.Sprintf(message, values...)
	hours, minutes, seconds := now.Clock()
	message = fmt.Sprintf("[%d:%d:%d] {%s} %s\n", hours, minutes, seconds, level, message)
	singletonFileLock.Lock()
	defer singletonFileLock.Unlock()
	_, _ = singletonLogFile.WriteString(message)
}

func Fatal(message string) {
	writeToFile(FATAL, message)
	os.Exit(1)
}

func CheckFatal(e error) {
	if e != nil {
		writeToFile(FATAL, e.Error())
		os.Exit(1)
	}
}

func CheckFatalMessage(e error, message string) {
	if e != nil {
		writeToFile(FATAL, "%s: %s", e.Error(), message)
		os.Exit(1)
	}
}

func CheckFatalMessagef(e error, message string, values ...interface{}) {
	message = fmt.Sprintf(message, values...)
	if e != nil {
		writeToFile(FATAL, "%s: %s", e.Error(), message)
		os.Exit(1)
	}
}

func Error(message string) {
	go writeToFile(ERROR, message)
}

func CheckError(e error) bool {
	if e != nil {
		go writeToFile(ERROR, e.Error())
		return true
	}
	return false
}

func CheckErrorMessage(e error, message string) bool {
	if e != nil {
		go writeToFile(ERROR, "%s: %s", e.Error(), message)
		return true
	}
	return false
}

func CheckErrorMessagef(e error, message string, values ...interface{}) bool {
	message = fmt.Sprintf(message, values...)
	if e != nil {
		go writeToFile(FATAL, "%s: %s", e.Error(), message)
		return true
	}
	return false
}

func Info(message string) {
	go writeToFile(INFO, message)
}

func Infof(message string, values ...interface{}) {
	message = fmt.Sprintf(message, values...)
	go writeToFile(INFO, message)
}

func openNewFileDaily() {
	// Change the file every day, so we store a month's worth of logs
	dailyDuration, err := time.ParseDuration("5m")
	singletonNewFileTicker = time.NewTicker(dailyDuration)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case <-singletonClose:
			return
		case _ = <-singletonNewFileTicker.C:
			openFile()
		}
	}
}

func openFile() {
	singletonFileLock.Lock()
	defer singletonFileLock.Unlock()
	day := time.Now().UTC().Day()
	logFilePath := path.Join(*singletonLogDir, fmt.Sprintf("%d.log", day))
	newLogFile, err := os.Create(logFilePath)
	if err != nil {
		Error("Cannot open new log file")
	} else {
		_ = singletonLogFile.Close()
		singletonLogFile = newLogFile
	}
}
