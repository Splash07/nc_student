package utils

import (
	"fmt"
	"log"
	"os"
)

var GeneralLogger *log.Logger
var ErrorLogger *log.Logger

func init() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error reading given path:", err)
		os.Exit(1)
	}
	// absPath, err := filepath.Abs("D:/Study/github.com/Splash07/nc_student/log")
	// if err != nil {
	// 	fmt.Println("Error reading given path:", err)
	// }

	generalLog, err := os.OpenFile(path+"/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	GeneralLogger = log.New(generalLog, "General:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(generalLog, "Error:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
