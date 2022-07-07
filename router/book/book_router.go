package book

import (
	"myapp/controller/book"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func GetRoute(route *iris.Application, db *gorm.DB) *iris.Application {

	booksAPI := route.Party("/books")
	{
		booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/", book.List)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", book.Create)
	}

	return route
}
