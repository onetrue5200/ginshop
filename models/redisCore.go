package models

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var redisCoreCtx = context.Background()

var RedisDB *redis.Client

func init() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisDB.Ping(redisCoreCtx).Result()
	if err != nil {
		fmt.Println(err)
	}
}
