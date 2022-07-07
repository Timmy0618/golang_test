package main

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/router"
	"myapp/router/book"

	"myapp/pkg/gorm"
)

func main() {
	config.Default()
	db, err := gorm.New()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(db)

	route := router.Default()
	route = book.GetRoute(route, db)
	route.Listen(":8080")
}
