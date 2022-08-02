package rmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func New() *amqp.Connection {
	const config string = "amqp://%s:%s@%s:%s/"

	sources := fmt.Sprintf(config,
		viper.GetString("rmq.USERNAME"),
		viper.GetString("rmq.PASSWORD"),
		viper.GetString("rmq.HOST"),
		viper.GetString("rmq.PORT"),
	)
	// conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	fmt.Println(sources)

	conn, err := amqp.Dial(sources)
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	fmt.Println("RMQ 連線成功")

	return conn
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
