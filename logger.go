package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	fLog *os.File
)

func create(path string) (err error) {
	if sLog, err := os.Stat(path); err == nil {
		if sLog.Size() >= 1048576 {
			os.Remove(path)
		}
	}
	fLog, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return
}
func New(path string) {
	if err := create(path); err != nil {
		fmt.Printf("|ERROR| %s\r\n", err)
		os.Exit(3)
	}
	log.SetOutput(fLog)
}
func Close() {
	fLog.Close()
}
