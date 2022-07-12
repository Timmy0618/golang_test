package rmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func New() *amqp.Connection {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()
	fmt.Println("RMQ 連線成功")

	return conn
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
