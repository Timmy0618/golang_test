package word

import (
	"myapp/controller/classification/word"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func GetRoute(route *iris.Application, db *gorm.DB) *iris.Application {

	wordController := word.New(db)
	wordAPI := route.Party("/classification/word")
	{
		wordAPI.Use(iris.Compression)

		// GET: http://localhost:8080/words
		wordAPI.Post("/", wordController.Create)
		wordAPI.Get("/{page}", wordController.List)
		wordAPI.Patch("/{id}", wordController.Update)
		wordAPI.Delete("/{id}", wordController.Delete)
	}

	return route
}
