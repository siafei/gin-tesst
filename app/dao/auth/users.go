package dao

import (
	"fmt"
	"github.com/siafei/gin-test/app/model"
	"github.com/siafei/gin-test/bootstrap"
	"github.com/siafei/gin-test/global"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	
}

func (u Users) Create(data map[string]interface{}) error {
	password , _ := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), bcrypt.DefaultCost)
	bootstrap.DBEngine.AutoMigrate(&model.Users{})
	user := model.Users{
		Username:       data["user_name"].(string),
		Password:     string(password),
		NickName:         data["nick_name"].(string),
	}
	if err := bootstrap.DBEngine.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u Users) GetUserByUserNamePass(username string,password string) (*model.Users,error)  {
	user := model.Users{}
	if bootstrap.DBEngine.Where("username=?",username).First(&user).Error != nil {
		return nil,&global.Error{
			StatusCode: 400,
			Msg: "不存在的用户",
			Code: 400,
		}
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		return nil,&global.Error{
			StatusCode: 400,
			Msg: "不存在的用户",
			Code: 400,
		}
	}
	return &user,nil
}