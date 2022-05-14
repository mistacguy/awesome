package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zshf.private/api"
	"zshf.private/global"
	"zshf.private/middleware"
)

func Router() {
	engine := gin.Default()

	//跨域
	engine.Use(middleware.Cors())

	//接口初始化
	web := engine.Group("/admin")
	{
		web.GET("/test", func(context *gin.Context) {
			context.String(200, "run")
		})
		// 用户登录API
		web.GET("/captcha", api.WebGetCaptcha)
		web.POST("/login", api.WebUserLogin)

		// 开启JWT认证,以下接口需要认证成功才能访问
		web.Use(middleware.JwtAuth())
	}

	// 启动、监听端口
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}
