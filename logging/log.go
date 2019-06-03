package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Filename is ...
var Filename = ""

// SetFilename is ...
func SetFilename(filename string) {
	Filename = filename
}

// Errorf is ...
func Errorf(msg string, a ...interface{}) {
	//internalLog ..
	internalLog(msg, a...)
}

// Logf is ...
func Logf(msg string, a ...interface{}) {
	internalLog(msg, a...)
}

// internalLog is ...
func internalLog(msg string, a ...interface{}) {

	//prepare the message
	outputMsg := fmt.Sprintf(msg, a...)
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")

	// prepare the message
	outputMsg = fmt.Sprintf("%s %s", timestamp, outputMsg)

	// print to screen and append to log file
	fmt.Println(outputMsg)

	appendToLogFile(outputMsg)
}

// appendToLogFile is ...
func appendToLogFile(msg string) {
	if Filename == "" {
		return
	}

	// append log to file
	f, err := os.OpenFile(Filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Failed opening log file %v", err)
	}
	defer f.Close()
	fmt.Fprintln(f, msg)
}
