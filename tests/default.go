package tests

import (
	"log"
	"myapp/config"
	redisPkg "myapp/pkg/redis"
	"myapp/pkg/rmq"
	"myapp/router"

	gormPkg "myapp/pkg/gorm"

	"github.com/go-redis/redis/v9"
	"github.com/kataras/iris/v12"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type test struct {
	route *iris.Application
	db    *gorm.DB
	rmq   *amqp.Connection
	rdb   *redis.Client
}

func New() *test {
	//讀取config
	config.Default()

	route := router.Default()

	// //新增logger檔案
	// f := logger.NewLogFile()
	// defer f.Close()

	// //設定logger
	// route.Logger().SetOutput(f)

	//設定rmq
	rmq := rmq.New()
	defer rmq.Close()

	//設定redis
	rdb := redisPkg.Default()
	defer rdb.Close()

	//db連線
	db, err := gormPkg.New()
	if err != nil {
		log.Fatal(err)
	}

	return &test{route, db, rmq, rdb}
}
