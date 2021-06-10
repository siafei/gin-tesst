package api

import "github.com/gin-gonic/gin"

type Article struct {

}

func NewArticle() Article {
	return Article{}
}

func (A Article) Detail (c *gin.Context)  {
	c.JSON(200,gin.H{"type":"detail"})
}

func (A Article) List (c *gin.Context)  {
	c.JSON(200,gin.H{"type":"List"})
}

func (A Article) Post (c *gin.Context)  {
	c.JSON(200,gin.H{"type":"Post"})
}

func (A Article) Put (c *gin.Context)  {
	c.JSON(200,gin.H{"article":"put"})
}

func (A Article) Delete (c *gin.Context)  {
	c.JSON(200,gin.H{"article":"delete"})
}