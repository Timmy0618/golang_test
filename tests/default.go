package tests

import (
	"log"
	"myapp/config"
	rmq "myapp/pkg/Rmq"
	redisPkg "myapp/pkg/redis"
	"myapp/router"
	"path/filepath"
	"runtime"
	"testing"

	gormPkg "myapp/pkg/gorm"
	logger "myapp/pkg/log"

	"github.com/go-redis/redis/v9"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type test struct {
	Route *iris.Application
	DB    *gorm.DB
	Rmq   *amqp.Connection
	Rdb   *redis.Client
}

func New() *test {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	viper.SetDefault("application.root", basePath+"/../")

	//讀取config
	config.Default()

	Route := router.Default()

	// //新增logger檔案
	f := logger.NewLogFile()
	// defer f.Close()

	// //設定logger
	Route.Logger().SetOutput(f)

	//設定rmq
	Rmq := rmq.New()
	// defer Rmq.Close()

	//設定redis
	Rdb := redisPkg.Default()
	// defer Rdb.Close()

	//db連線
	DB, err := gormPkg.New()
	if err != nil {
		log.Fatal(err)
	}

	return &test{Route, DB, Rmq, Rdb}
}

func IrisTester(Route *iris.Application, t *testing.T) *httptest.Expect {
	handler := Route
	return httptest.New(t, handler)
}
