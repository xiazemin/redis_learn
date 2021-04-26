package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type userCredentialLoginCount struct {
	ID                string // 用户id也可能是用户名
	account_pwd_count int    // 账号或密码错误次数
	token_count       int    // 令牌错误次数
}

func main() {
	//建立连接
	sf := &redis.FailoverOptions{
		// The master name.
		MasterName: "mymaster",
		// A seed list of host:port addresses of sentinel nodes.
		SentinelAddrs: []string{"127.0.0.1:26379", "127.0.0.1:26380", "127.0.0.1:26381"},

		// Following options are copied from Options struct.
		Password: "",
		DB:       2,
	}
	//ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)
	//user_credential_login_count := &userCredentialLoginCount{
	//	ID:             "hast_users",
	//	AcountPwdCount: 1,
	//	TokenCount:     2,
	//}
	datas := map[string]interface{}{
		//"id":  "hast_users",
		"tel": 1,
		"age": 2,
	}
	rdb := redis.NewFailoverClient(sf)
	ctx := context.Background()
	fmt.Println("client:", rdb)
	err1 := rdb.Ping(ctx).Err()
	if err1 != nil {
		fmt.Println("Ping:", err1)
	}
	// 删除某个key 以及对应的值
	// rdb.Del(ctx,"hast_users")

	bExist, err := rdb.HExists(ctx, "hash_tests", "tel").Result()
	log.Println(bExist, err)
	bRet, err := rdb.HSet(ctx, "hash_tests", "tel", 3).Result()
	log.Println(bRet, err)

	rdb.HSet(ctx, "hash_tests", "token_count", 2).Result()
	rdb.HSet(ctx, "hash_tests", "account_pwd_count", 3).Result()
	//bExist, err = rdb.HExists(ctx, "hash_tests", "tel").Result()
	//log.Println(bExist, err)
	//resq := rdb.HGet(ctx, "hash_tests", "tel")
	//log.Println(resq.Int())
	hmgetall := rdb.HGetAll(ctx, "hash_tests")
	if hmgetall.Err() != nil {
		log.Printf("###### HMGet hmgetall err", hmgetall.Err())
	}
	log.Printf("###### HMGet hmgetall Val", hmgetall.Val())

	log.Printf("###### HMSet  #################")
	rst := rdb.HMSet(ctx, "hast_users", datas)
	log.Printf("###### HMSet rst err", rst.Err())
	log.Printf("###### HMSet rst Val", rst.Val())
	// 如果要对key设置过期时间
	rdb.Expire(ctx, "hast_users", 10*time.Minute)

	log.Printf("###### HMGet  #################")

	log.Printf("###### HMGet  tel#################")
	hmget := rdb.HMGet(ctx, "hast_users", "tel")
	if hmget.Err() != nil {
		log.Printf("###### HMGet hmget err", hmget.Err())
	}
	log.Printf("###### HMGet hmget Val", hmget.Val())

	log.Printf("###### HMGet  token_count #################")
	hmget = rdb.HMGet(ctx, "hast_users", "age")
	if hmget.Err() != nil {
		log.Printf("###### HMGet hmget err", hmget.Err())
	}
	log.Printf("###### HMGet hmget Val", hmget.Val())

	bRet, err = rdb.HSet(ctx, "hast_users", "age", 3).Result()
	log.Println(bRet, err)
	log.Printf("###### HMGetALL  #################")
	hmgetall = rdb.HGetAll(ctx, "hast_users")
	if hmgetall.Err() != nil {
		log.Printf("###### HMGet hmgetall err", hmgetall.Err())
	}

	log.Printf("###### HMGet hmgetall Val", hmgetall.Val())

}
