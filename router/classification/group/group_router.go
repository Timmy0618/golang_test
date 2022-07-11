package group

import (
	"myapp/controller/classification/group"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func GetRoute(route *iris.Application, db *gorm.DB) *iris.Application {

	groupController := group.New(db)
	groupAPI := route.Party("/classification/group")
	{
		groupAPI.Use(iris.Compression)

		// GET: http://localhost:8080/words
		groupAPI.Post("", groupController.Create)
		groupAPI.Get("/{page}", groupController.List)
		groupAPI.Patch("/{id}", groupController.Update)
		groupAPI.Delete("/{id}", groupController.Delete)
	}

	return route
}
