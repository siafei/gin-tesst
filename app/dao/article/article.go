package dao

import (
	"github.com/siafei/gin-test/app/model"
	"github.com/siafei/gin-test/bootstrap"
	"github.com/siafei/gin-test/global"
)

type Article struct {

}

func (a Article) Create(data map[string]interface{}) (interface{},error) {
	return nil,nil
}

func (a Article) Update(data map[string]interface{}) (interface{},error) {
	return nil,nil
}

func (a Article) Detail(id uint32) (interface{},error)  {
	article := model.Article{}
	if bootstrap.DBEngine.Where("id=?",id).First(&article).Error != nil {
		return nil,&global.Error{
			StatusCode: 400,
			Msg: "不存在的文章",
			Code: 400,
		}
	}
	return &article,nil
}

func (a Article) List(username string,password string) ([]interface{},error)  {
	return nil,nil
}

func (a Article) Delete(username string,password string) error  {
	return nil
}