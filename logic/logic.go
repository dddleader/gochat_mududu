package logic

import (
	"gochat_mududu/config"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Logic struct {
	ServerId string
}

func New() *Logic {
	return new(Logic)
}

//logic层接收api层消息并push至消息队列之中
//接收登录消息，返回AuthToken

func (logic *Logic) Run() {
	runtime.GOMAXPROCS(config.Conf.Logic.LogicBase.CpuNum)
	//redis client for publishing msg
	err := logic.InitPublishRedisClient()
	if err != nil {
		logrus.Error("err in InitPublishRedisClient")
	}
	//rpc server for rpc call

}
