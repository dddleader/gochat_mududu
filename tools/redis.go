package tools

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// 保管同redis client的链接
var RedisClientMap = map[string]*redis.Client{}
var syncLock sync.Mutex

// 工具方法作用一：提供数据结构
type RedisOption struct {
	Address  string
	Password string
	Db       int
}

func GetRedisInstance(redisOpt RedisOption) *redis.Client {
	syncLock.Lock()
	if redisCli, ok := RedisClientMap[redisOpt.Address]; ok {
		return redisCli
	}
	client := redis.NewClient(&redis.Options{
		Addr:       redisOpt.Address,
		Password:   redisOpt.Password, // 密码
		DB:         redisOpt.Db,       // 数据库
		MaxConnAge: 20 * time.Second,
	})
	RedisClientMap[redisOpt.Address] = client
	syncLock.Unlock()
	return RedisClientMap[redisOpt.Address]
}
