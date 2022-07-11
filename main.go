package main

import (
	"log"
	"myapp/config"
	logger "myapp/pkg/log"
	"myapp/router"
	"myapp/router/book"
	"myapp/router/classification/group"
	"myapp/router/classification/word"

	"myapp/pkg/gorm"
)

func main() {
	config.Default()

	route := router.Default()

	//新增logger檔案
	f := logger.NewLogFile()
	defer f.Close()

	//設定logger
	route.Logger().SetOutput(f)

	db, err := gorm.New()
	if err != nil {
		log.Fatal(err)
		return
	}

	route = book.GetRoute(route, db)
	route = word.GetRoute(route, db)
	route = group.GetRoute(route, db)
	route.Listen(":8080")
}
