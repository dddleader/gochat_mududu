package logic

import (
	"fmt"
	"gochat_mududu/config"
	"runtime"

	"github.com/google/uuid"
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
	logicConfig := config.Conf.Logic

	runtime.GOMAXPROCS(logicConfig.LogicBase.CpuNum)
	logic.ServerId = fmt.Sprintf("logic-%s", uuid.New().String())
	//redis client for publishing msg
	err := logic.InitPublishRedisClient()
	if err != nil {
		logrus.Panicf("logic init publishRedisClient fail,err:%s", err.Error())
	}
	//rpc server for rpc call
	//初始化RPC服务器
	if err := logic.InitRpcServer(); err != nil {
		logrus.Panicf("logic init rpc server fail")
	}
}
