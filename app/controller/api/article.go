package api

import (
	"github.com/gin-gonic/gin"
	param "github.com/siafei/gin-test/app/param/tag"
	article_service "github.com/siafei/gin-test/app/service/article"
	"github.com/siafei/gin-test/bootstrap"
	"github.com/siafei/gin-test/global"
)

type Article struct {

}

func NewArticle() Article {
	return Article{}
}

var (
	service_article article_service.Article
)
/**
*文章详情
 */
func (A Article) Detail (c *gin.Context)  {
	article,err := service_article.Detail(c)
	if err !=nil {
		bootstrap.NewResponse(c).ToErrorResponse(global.OptionError(err.Error()))
		return
	}
	bootstrap.NewResponse(c).ToResponse(gin.H{"detail":article})
	return
}

/**
*文章列表
 */
type SeList struct {
	Id int
	TagId int
}
func (A Article) List (c *gin.Context)  {
	var param param.TagList
	if err := c.ShouldBind(&param);err != nil {
		bootstrap.NewResponse(c).ToErrorResponse(global.OptionError(err.Error()))
		return
	}
	var article []SeList
	bootstrap.DBEngine.Table("blog_article").Where("tag_id=?",2).Select("blog_article.id as id,tag_id").Joins("JOIN blog_article_tag ON blog_article_tag.artilce_id = blog_article.id").Scan(&article)
	c.JSON(200,article)
}

/**
*添加文章
 */

func (A Article) Post (c *gin.Context)  {
	var param param.TagPost
	if err := c.ShouldBind(&param);err != nil {
		bootstrap.NewResponse(c).ToErrorResponse(global.OptionError(err.Error()))
		return
	}
	c.JSON(200,gin.H{"article":"post"})
}

/**
*修改文章
 */
func (A Article) Put (c *gin.Context)  {
	c.JSON(200,gin.H{"article":"put"})
}

/**
*删除文章
 */
func (A Article) Delete (c *gin.Context)  {
	c.JSON(200,gin.H{"article":"delete"})
}