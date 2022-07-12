package word

import (
	"myapp/controller/classification/word"

	"github.com/kataras/iris/v12"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func GetRoute(route *iris.Application, db *gorm.DB, rmq *amqp.Connection) *iris.Application {

	wordController := word.New(db, rmq)
	wordAPI := route.Party("/classification/word")
	{
		wordAPI.Use(iris.Compression)

		// GET: http://localhost:8080/words
		wordAPI.Post("/", wordController.Create)
		wordAPI.Get("/{page}", wordController.List)
		wordAPI.Patch("/{id}", wordController.Update)
		wordAPI.Delete("/{id}", wordController.Delete)
		wordAPI.Post("/rmq", wordController.RmqAdd)
	}

	return route
}
