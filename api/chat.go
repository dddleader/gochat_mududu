package api

import (
	"context"
	"fmt"
	"gochat_mududu/api/router"
	"gochat_mududu/api/rpc"
	"gochat_mududu/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 将api层相关服务逻辑封装至Chat之中
type Chat struct {
}

func New() *Chat {
	return &Chat{}
}

// api server handle request from user
// send rpc request to other layers
func (c *Chat) Run() {
	rpc.InitLogicRpcServer()
	r := router.Register()
	//default as debug
	runMode := config.GetGinRunMode()
	logrus.Info("server start , now run mode is ", runMode)
	gin.SetMode(runMode)
	port := config.Conf.Api.ApiBase.ListenPort
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
	//start http function
	go func() {
		//listen to port 7070
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("start listen : %s\n", err)
		}
	}()
	//set chan to receive quit signal
	quit := make(chan os.Signal)
	//receive挂起、中断、终止、退出
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	logrus.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("Server Shutdown:", err)
	}
	logrus.Infof("Server exiting")
	os.Exit(0)
}
