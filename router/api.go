package router

import (
	"github.com/gin-gonic/gin"
	"siafei/gin-test/app/controller/api"
)

func registerApi(router *gin.RouterGroup)  {

	//需要登录验证的 接口
	//router.Use(middleware.CheckToken())
	//{
		//文章分组
		router_article := router.Group("/article")
		{
			article := api.NewArticle() //文章控制器
			router_article.GET("/:id",article.Detail)    //文章详情
			router_article.GET("",article.List)    //文章列表
			router_article.POST("",article.Post)    //文章创建
			router_article.PUT("/:id",article.Put)    //文章修改
			router_article.DELETE("/:id",article.Delete)    //文章删除
		}


	//}


}
