package log

import (
	"os"
	"time"
)

func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func NewLogFile() *os.File {
	filename := todayFilename()
	f, err := os.OpenFile("./pkg/log/"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
