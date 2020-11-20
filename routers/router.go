package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine  {
	r := gin.New()

	/**
	*引用中间件
	 */
	r.Use(gin.Logger())	//输出请求日志，并标准化日志格式
	r.Use(gin.Recovery())	//异常捕获 防止因为出现 panic 导致服务崩溃 同时将异常日志格式标准化


	api := r.Group("api")	//引入api 文件分组
	registerApi(api)	//引入api 文件分组
	return  r
}