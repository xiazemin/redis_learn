package main

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	ctx := context.TODO()
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
}
