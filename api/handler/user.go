package handler

import (
	"gochat_mududu/api/rpc"
	"gochat_mududu/proto"
	"gochat_mududu/tools"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 由前端发来的json form
type FormLogin struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"passWord" json:"passWord" binding:"required"`
}

func Login(c *gin.Context) {
	var formLogin FormLogin
	//解析json结构体
	if err := c.ShouldBindBodyWith(&formLogin, binding.JSON); err != nil {
		tools.FailWithMsg(c, err.Error())
		return
	}
	//向logic层发出登录rpc
	req := &proto.LoginRequest{
		Name:     formLogin.UserName,
		Password: tools.Sha1(formLogin.Password),
	}
	code, authToken, msg := rpc.RpcLogicObj.Login(req)
	if code == tools.CodeFail || authToken == "" {
		tools.FailWithMsg(c, msg)
		return
	}
	tools.SuccessWithMsg(c, "login success", authToken)
}

type FormRegister struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"passWord" json:"passWord" binding:"required"`
}

func Register(c *gin.Context) {
	var formRegister FormRegister
	if err := c.ShouldBindBodyWith(&formRegister, binding.JSON); err != nil {
		tools.FailWithMsg(c, err.Error())
		return
	}
	req := &proto.RegisterRequest{
		Name:     formRegister.UserName,
		Password: tools.Sha1(formRegister.Password),
	}
	code, authToken, msg := rpc.RpcLogicObj.Register(req)
	if code == tools.CodeFail || authToken == "" {
		tools.FailWithMsg(c, msg)
		return
	}
	tools.SuccessWithMsg(c, "register success", authToken)
}
