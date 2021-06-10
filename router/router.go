package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	router := gin.New()
	//启用 日志和 recover中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//注册api 路由
	api := router.Group("api")
	registerApi(api)

	return router
}