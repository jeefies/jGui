package jgui

import (
	"log"
	"os"
)

var logger *log.Logger
var loggerFile *os.File

func init() {
	var err error
	loggerFile, err = os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0644)
	check(err)

	logger = log.New(loggerFile,  "Jgui logger:", log.Lmsgprefix |  log.Lmicroseconds | log.Lshortfile)

	logger.SetOutput(os.Stdout)
}
