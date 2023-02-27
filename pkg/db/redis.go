package db

import (
	"fmt"
	"sync"

	"github.com/richkeyu/gocommons/config"
	"github.com/richkeyu/gocommons/redis"
)

var (
	kmu sync.Mutex

	redisConf map[string]redis.Config
)

func InitRedis() {
	var err error
	if err = config.Load("redis", &redisConf); err != nil {
		panic(fmt.Sprintf("load redis config failed, err: %v", err))
	}
}

//获取KV链接池
func GetKvConn(name string) *redis.Client {

	kmu.Lock()
	defer kmu.Unlock()

	client := redis.GetClient(name)
	if client != nil {
		return client
	}

	if conf, ok := redisConf[name]; ok {
		redis.InitRedisClient(name, conf)
		client = redis.GetClient(name)
	}

	return client
}
