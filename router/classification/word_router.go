package classification

import (
	"myapp/controller/classification"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func GetRoute(route *iris.Application, db *gorm.DB) *iris.Application {

	classificationController := classification.New(db)
	classificationsAPI := route.Party("/classification")
	{
		classificationsAPI.Use(iris.Compression)

		// GET: http://localhost:8080/words
		classificationsAPI.Post("/", classificationController.Create)
		classificationsAPI.Get("/{page}", classificationController.List)
		classificationsAPI.Patch("/{id}", classificationController.Update)
		classificationsAPI.Delete("/{id}", classificationController.Delete)
	}

	return route
}
