package redis

import (
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

func Default() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + viper.GetString("redis.REDIS_PORT"),
		Password: viper.GetString("redis.REDIS_PASSWORD"),
		DB:       0, // use default DB
	})
	fmt.Println("redis 連線成功")

	return rdb
}
