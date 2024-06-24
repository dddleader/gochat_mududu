package router

import (
	"gochat_mududu/api/handler"
	"gochat_mududu/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用gin框架进行分组、路由
func Register() *gin.Engine {
	r := gin.Default()
	r.Use(CorsMiddleware())
	initUserRouter(r)

	return r
}

// router for user to call
func initUserRouter(r *gin.Engine) {
	//set group
	userGroup := r.Group("/user")
	userGroup.POST("/login", handler.Login)
	userGroup.POST("/register", handler.Register)
	r.NoRoute(func(c *gin.Context) {
		tools.FailWithMsg(c, "please check request url !")
	})
}

func CorsMiddleware() gin.HandlerFunc {
	//middleware 对context进行处理
	return func(c *gin.Context) {
		method := c.Request.Method
		var openCorsFlag = true
		if openCorsFlag {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONs" {
			c.JSON(http.StatusOK, nil)
		}
		c.Next()
	}
}
