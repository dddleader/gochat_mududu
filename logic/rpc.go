package logic

import (
	"context"
	"errors"
	"gochat_mududu/config"
	"gochat_mududu/logic/dao"
	"gochat_mududu/proto"
	"gochat_mududu/tools"
	"time"

	"github.com/sirupsen/logrus"
)

type RpcLogic struct {
}

// 接收api层登录rpc
func (rpc *RpcLogic) Login(ctx context.Context, args *proto.LoginRequest, reply *proto.LoginResponse) (err error) {
	//initialize to be fail code
	reply.Code = config.FailReplyCode
	u := new(dao.User)
	userName := args.Name
	passWord := args.Password
	//db act as persistent storage
	data := u.GetUserByUserName(userName)
	if (data.Id == 0) || (passWord != data.Password) {
		return errors.New("no this user or password error!")
	}
	//存储登陆状态
	loginSessionId := tools.GetSessionIdByUserId(data.Id)
	//set token
	//err = redis.HMSet(auth, userData)
	randToken := tools.GetRandomToken(32)
	sessionId := tools.CreateSessionId(randToken)
	userData := make(map[string]interface{})
	userData["userId"] = data.Id
	userData["userName"] = data.UserName
	//check is login
	//如果一集登录，则删除旧会话令牌
	token, _ := RedisSessClient.Get(loginSessionId).Result()
	if token != "" {
		//logout already login user session
		oldSession := tools.CreateSessionId(token) //如果重复，删除用户数据
		err := RedisSessClient.Del(oldSession).Err()
		if err != nil {
			return errors.New("logout user fail!token is:" + token)
		}
	}
	//设置redis会话
	RedisSessClient.Do("MULTI")
	//使用新的authtoken生成的sessionId存储userdata
	RedisSessClient.HMSet(sessionId, userData) //同一时间，一个token对应存储一个hash表
	RedisSessClient.Expire(sessionId, 86400*time.Second)
	RedisSessClient.Set(loginSessionId, randToken, 86400*time.Second) //同一时间，一个Id对应一个unique token
	err = RedisSessClient.Do("EXEC").Err()
	//err = RedisSessClient.Set(authToken, data.Id, 86400*time.Second).Err()
	if err != nil {
		logrus.Infof("register set redis token fail!")
		return err
	}
	reply.Code = config.SuccessReplyCode
	reply.AuthToken = randToken
	return
}
func (rpc *RpcLogic) Register(ctx context.Context, args *proto.RegisterRequest, reply *proto.RegisterReply) (err error) {
	reply.Code = config.FailReplyCode
	u := new(dao.User)
	uData := u.CheckHaveUserName(args.Name)
	if uData.Id > 0 {
		return errors.New("this user name already have , please login !!!")
	}
	u.UserName = args.Name
	u.Password = args.Password
	//create new account
	userId, err := u.Add()
	if err != nil {
		logrus.Infof("register err:%s", err.Error())
		return err
	}
	if userId == 0 {
		return errors.New("register userId empty!")
	}
	//set token
	randToken := tools.GetRandomToken(32)
	sessionId := tools.CreateSessionId(randToken)
	userData := make(map[string]interface{})
	userData["userId"] = userId
	userData["userName"] = args.Name
	RedisSessClient.Do("MULTI")
	RedisSessClient.HMSet(sessionId, userData)
	RedisSessClient.Expire(sessionId, 86400*time.Second)
	err = RedisSessClient.Do("EXEC").Err()
	if err != nil {
		logrus.Infof("register set redis token fail!")
		return err
	}
	//客户端保有randToken,随时可生成sessionId
	reply.Code = config.SuccessReplyCode
	reply.AuthToken = randToken
	return
}
