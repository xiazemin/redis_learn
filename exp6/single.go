package main

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.TODO()
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	err = client.Set(ctx, "feekey", "examples", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "feekey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("feekey", val)

	val2, err := client.Get(ctx, "feekey2").Result()
	if err == redis.Nil {
		fmt.Println("feekey does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("feekey", val2)
	}
}

func main() {
	ExampleClient()
}
