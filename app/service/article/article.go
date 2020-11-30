package service

import (
	"github.com/gin-gonic/gin"
	dao_article "github.com/siafei/gin-test/app/dao/article"
	"github.com/siafei/gin-test/bootstrap"
)

type Article struct {

}
func NewAuth() Article {
	return Article{}
}

var (
	article_dao dao_article.Article
)

func (a Article) Create(data map[string]interface{}) (interface{},error) {
	return nil,nil
}

func (a Article) Update(data map[string]interface{}) (interface{},error) {
	return nil,nil
}

func (a Article) Detail(c *gin.Context) (interface{},error)  {
	id := bootstrap.StrTo(c.Param("id")).MustUInt32()
	article,err := article_dao.Detail(id)
	if err != nil {
		return nil, err
	}
	return article,nil
}

func (a Article) List(username string,password string) ([]interface{},error)  {
	return nil,nil
}

func (a Article) Delete(username string,password string) error  {
	return nil
}