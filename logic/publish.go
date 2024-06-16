package logic

import (
	"gochat_mududu/config"
	"gochat_mududu/tools"

	"github.com/sirupsen/logrus"
)

// use redis as msg queue
func (logic *Logic) InitPublishRedisClient() (err error) {
	redisOpt := tools.RedisOption{
		Address:  config.Conf.Common.CommonRedis.RedisAddress,
		Password: config.Conf.Common.CommonRedis.RedisPassword,
		Db:       config.Conf.Common.CommonRedis.Db,
	}
	RedisClient := tools.GetRedisInstance(redisOpt)
	//ping pong to test server
	if pong, err := RedisClient.Ping().Result(); err != nil {
		logrus.Infof("RedisCli Ping Result pong: %s,  err: %s", pong, err)
	}
	return err
}
