package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

// GeneralLogger exported
var GeneralLogger *log.Logger

// ErrorLogger exported
var ErrorLogger *log.Logger

func init() {

	generalLog, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	mw := io.MultiWriter(os.Stdout, generalLog)

	GeneralLogger = log.New(mw, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(mw, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
