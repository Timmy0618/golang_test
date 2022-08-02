package main

import (
	"log"
	"myapp/config"
	logger "myapp/pkg/log"
	"myapp/pkg/redis"
	"myapp/pkg/rmq"
	"myapp/router"
	"myapp/router/classification/group"
	"myapp/router/classification/word"
	"path/filepath"
	"runtime"

	"myapp/pkg/gorm"

	"github.com/spf13/viper"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	viper.SetDefault("application.root", basePath)

	//讀取config
	config.Default()

	route := router.Default()

	//新增logger檔案
	f := logger.NewLogFile()
	defer f.Close()

	//設定logger
	route.Logger().SetOutput(f)

	//設定rmq
	rmq := rmq.New()
	defer rmq.Close()

	//設定redis
	rdb := redis.Default()
	defer rdb.Close()

	//db連線
	db, err := gorm.New()
	if err != nil {
		log.Fatal(err)
		return
	}

	route = word.GetRoute(route, db, rmq, rdb)
	route = group.GetRoute(route, db)
	route.Listen(":8080")
}
