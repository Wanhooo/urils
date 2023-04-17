package initialize

import (
	"github.com/gin-gonic/gin"
	"urils/application/middleware"

	"urils/application/router"
)

func InitRouter() *gin.Engine {
	// 1. 创建路由
	Router := gin.Default()
	// 注册中间件
	Router.Use(middleware.Cors(), middleware.ExceptionMiddleware, middleware.GinLogger(), middleware.GinRecovery(true))
	// 2. Api路由分组
	ApiGroup := Router.Group("/api")
	// 3. 初始化用户相关路由
	router.InitUserRouter(ApiGroup)
	return Router
}
