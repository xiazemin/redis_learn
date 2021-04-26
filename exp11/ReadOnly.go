package main

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:7000",
			"127.0.0.1:7001",
			"127.0.0.1:7002",
		}, //set redis cluster url
		Password: "", //set password
	})

	readonlyClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:7003",
			"127.0.0.1:7004",
			"127.0.0.1:7005"}, //set redis cluster url
		Password: "", //set password
		ReadOnly: true,
	})

	ctx := context.TODO()
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	fmt.Println(client.Set(ctx, "hello", 1234, 1*time.Minute))
	fmt.Println(client.Get(ctx, "hello"))
	fmt.Println(readonlyClient.Set(ctx, "hello1", 1234, 1*time.Minute))
	fmt.Println(readonlyClient.Get(ctx, "hello1"))
	fmt.Println(readonlyClient.Get(ctx, "hello"))
	fmt.Println(client.Get(ctx, "hello1"))
	fmt.Println(readonlyClient.ReadOnly(ctx))
	fmt.Println(readonlyClient.Get(ctx, "hello1"))
	fmt.Println(readonlyClient.Get(ctx, "hello"))
}
