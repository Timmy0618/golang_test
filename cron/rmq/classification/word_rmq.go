package main

import (
	"fmt"
	"log"
	"myapp/config"
	wordModel "myapp/model/classification/word"
	"myapp/pkg/gorm"
	"myapp/pkg/rmq"
	"myapp/services/classification"

	"reflect"
)

func main() {
	config.Default()

	conn := rmq.New()
	db, _ := gorm.New()

	var wordList []wordModel.Word
	result := db.Limit(10).Offset(0).Find(&wordList)
	if result.Error != nil {
		fmt.Println("List fail")
		return
	}
	fmt.Println(wordList)

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"sentence_queue", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", reflect.TypeOf(d.Body))
			classification.Classify(d.Body, wordList)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
