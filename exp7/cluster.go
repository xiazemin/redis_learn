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
			"127.0.0.1:7003",
			"127.0.0.1:7004",
			"127.0.0.1:7005"}, //set redis cluster url
		Password: "", //set password
	})

	ctx := context.TODO()
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	fmt.Println("pool state init state:", client.PoolStats())
	for i := 0; i < 1000; i++ {
		k := fmt.Sprintf("key:%d", i)
		v := k
		val, err := client.Set(ctx, k, v, 60*time.Second).Result()
		if err != nil {
			fmt.Println(val)
			panic(err)
		}

		val, err = client.Get(ctx, k).Result()
		if err != nil {
			fmt.Println(val)
			panic(err)
		}
		//fmt.Println("key:", val)
	}
	fmt.Println("pool state final state:", client.PoolStats()) //获取客户端连接池相关信息

}
