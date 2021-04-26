package main

import (
	"context"
	"fmt"
	"net"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"127.0.0.1:26379", "127.0.0.1:26380", "127.0.0.1:26381"},
	})
	ctx := context.TODO()
	fmt.Println("flush err:", client.FlushDB(ctx).Err())
	c, err := client.SlaveForKey(ctx, "test")
	fmt.Println(err, c)

	sentinel := redis.NewSentinelClient(&redis.Options{
		Addr:       ":26379",
		MaxRetries: -1,
	})

	addr, err := sentinel.GetMasterAddrByName(ctx, "mymaster").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(addr)

	master := redis.NewClient(&redis.Options{
		Addr:       net.JoinHostPort(addr[0], addr[1]),
		MaxRetries: -1,
	})
	masterPort := addr[1]

	fmt.Println(master, masterPort)
	// Wait until slaves are picked up by sentinel.

}
