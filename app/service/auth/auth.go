package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	dao "github.com/siafei/gin-test/app/dao/auth"
	"github.com/siafei/gin-test/app/service"
	"github.com/siafei/gin-test/global"
	"time"
)

type Auth struct {
	
}
func NewAuth() Auth {
	return Auth{}
}

var (
	dao_users dao.Users
)

func (A Auth) Login(c *gin.Context) (string,error) {
	_,err := dao_users.GetUserByUserNamePass(c.Query("user_name"),c.Query("password"))
	if err != nil {
		return "",global.OptionError("账号或密码错误")
	}
	token,err := service.GenerateToken(c.Query("user_name"),c.Query("password"))
	if err != nil {
		return "" ,global.OptionError("生成token失败,请稍后再试")
	}
	return token,nil
}

/**
*注册
 */
func (A Auth) Register(c *gin.Context) (string,error)  {
	user_create := map[string]interface{}{
		"user_name"		:	string(c.Query("user_name")),
		"password"	:	string(c.Query("password")),
		"nick_name"	: 	string(c.Query("nick_name")),
	}
	err := dao_users.Create(user_create)
	if err != nil {
		return "" ,&global.Error{
			StatusCode: 500,
			Msg: "注册失败，请稍后再试",
			Code: 400,
		}
	}
	token,err := service.GenerateToken(c.Query("user_name"),c.Query("password"))
	if err != nil {
		return "" ,&global.Error{
			StatusCode: 400,
			Msg: err.Error(),
			Code: 400,
		}
	}
	return token,nil
}

func (A Auth) CheckToken(c *gin.Context) (interface{},error)  {
	token,err := service.GetToken(c)
	if err != nil {
		return nil,err
	}
	claims,err := service.ParseToken(token)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	if claims.VerifyExpiresAt(time.Now().Unix(),true) == false {
		return nil,&global.Error{
			StatusCode: 403,
			Msg: "登录已经过期，请重新登录",
			Code: 403,
		}
	}
	user,err := dao_users.GetUserByUserNamePass(claims.Username,claims.Password)
	if err != nil {
		return nil,&global.Error{
			StatusCode: 403,
			Msg: "不存在的用户，请重新登录",
			Code: 403,
		}
	}
	return user,err
}