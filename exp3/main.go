package main

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

func main1() {
	// See http://redis.io/topics/sentinel for instructions how to
	// setup Redis Sentinel.
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{":26379"},
	})
	ctx := context.TODO()
	rdb.Ping(ctx)

}

func main() {
	// clusterSlots returns cluster slots information.
	// It can use service like ZooKeeper to maintain configuration information
	// and Cluster.ReloadState to manually trigger state reloading.
	clusterSlots := func(ctx context.Context) ([]redis.ClusterSlot, error) {
		slots := []redis.ClusterSlot{
			// First node with 1 master and 1 slave.
			/**
			Master[0] -> Slots 0 - 5460
			Master[1] -> Slots 5461 - 10922
			Master[2] -> Slots 10923 - 16383
			Adding replica 127.0.0.1:7004 to 127.0.0.1:7000
			Adding replica 127.0.0.1:7005 to 127.0.0.1:7001
			Adding replica 127.0.0.1:7003 to 127.0.0.1:7002
			*/
			{
				Start: 0,
				End:   5460,
				Nodes: []redis.ClusterNode{{
					Addr: ":7000", // master
				}, {
					Addr: ":7004", // master
				},
				},
			},
			// Second node with 1 master and 1 slave.
			{
				Start: 5461,
				End:   10922,
				Nodes: []redis.ClusterNode{{
					Addr: ":7001", // master
				}, {
					Addr: ":7005", // 1st slave
				}},
			},
			{
				Start: 10923,
				End:   16383,
				Nodes: []redis.ClusterNode{{
					Addr: ":7002", // master
				}, {
					Addr: ":7003", // 1st slave
				}},
			},
		}
		return slots, nil
	}

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots:  clusterSlots,
		RouteRandomly: true,
	})
	ctx := context.TODO()
	fmt.Println(rdb.Ping(ctx))

	// ReloadState reloads cluster state. It calls ClusterSlots func
	// to get cluster slots information.
	rdb.ReloadState(ctx)
}
