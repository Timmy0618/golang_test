package main

import (
	"myapp/router"
	"myapp/router/book"
)

func main() {
	route := router.Default()
	route = book.GetRoute(route)
	route.Listen(":8080")
}
