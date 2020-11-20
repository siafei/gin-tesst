package api

import "github.com/gin-gonic/gin"

type Article struct {

}

func NewArticle() Article {
	return Article{}
}


/**
*文章详情
 */
func (A Article) Detail (c *gin.Context)  {
	c.JSON(200,gin.H{"article":"detail"})
}

/**
*文章列表
 */
func (A Article) List (c *gin.Context)  {
	c.JSON(200,gin.H{"article":"list"})
}

/**
*添加文章
 */

func (A Article) Post (c *gin.Context)  {
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