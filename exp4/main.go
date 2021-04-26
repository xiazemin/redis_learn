package main

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"shard1": ":7000",
			"shard2": ":7001",
			"shard3": ":7002",
		},
	})
	ctx := context.TODO()
	fmt.Println(rdb.Ping(ctx))
}
