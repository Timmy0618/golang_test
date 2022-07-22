package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func Set(rdb *redis.Client, key string, value interface{}) {
	value, _ = json.Marshal(value)
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func Get(rdb *redis.Client, key string) []byte {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return []byte(val)
	// var result interface{}
	// json.Unmarshal([]byte(val), &result)
	// fmt.Println(reflect.TypeOf(result))
	// return result
}

func Del(rdb *redis.Client, key string) {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}
}

func Scan(rdb *redis.Client, key string) bool {
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = rdb.Scan(ctx, cursor, key, 0).Result()
		if err != nil {
			panic(err)
		}

		if len(keys) == 0 {
			return false
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}
	return true
}
