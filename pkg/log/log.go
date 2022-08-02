package log

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func NewLogFile() *os.File {
	basePath := viper.GetString("application.root")

	filename := todayFilename()
	f, err := os.OpenFile(basePath+"./pkg/log/"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
