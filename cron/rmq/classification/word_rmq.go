package main

import (
	"encoding/json"
	"log"
	"myapp/config"
	wordModel "myapp/model/classification/word"
	"myapp/pkg/redis"

	"myapp/pkg/rmq"
	"myapp/services/classification"
	wordService "myapp/services/classification/word"
	redisService "myapp/services/redis"
	rmqServices "myapp/services/rmq"

	"reflect"
)

func main() {
	config.Default()

	rmq := rmq.New()

	//redis連線
	var wordList []wordModel.Word
	rdb := redis.Default()

	//rmq
	msgs := rmqServices.Pop(rmq)

	var forever chan struct{}

	go func() {
		for d := range msgs {

			//檢查wordList 有沒有被更新
			if !redisService.Scan(rdb, "wordList") {
				redisService.Set(rdb, "wordList", wordService.GetWordList())
			}

			err := json.Unmarshal(redisService.Get(rdb, "wordList"), &wordList)
			if err != nil {
				panic(err)
			}

			log.Printf("Received a message: %s", reflect.TypeOf(d.Body))
			classification.Classify(d.Body, wordList)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
