package api

import (
	"github.com/gin-gonic/gin"
	param "github.com/siafei/gin-test/app/param/auth"
	service "github.com/siafei/gin-test/app/service/auth"
	"github.com/siafei/gin-test/bootstrap"
	"github.com/siafei/gin-test/global"
)

type Auth struct {

}

var (
	auth_service service.Auth
)

func (A Auth) Register(c *gin.Context)  {
	var param param.Login
	if err := c.ShouldBind(&param);err != nil {
		bootstrap.NewResponse(c).ToErrorResponse(global.OptionError(err.Error()))
		return
	}
	token,err := auth_service.Register(c)
	if err !=nil {
		bootstrap.NewResponse(c).ToErrorResponse(global.OptionError(err.Error()))
		return
	}
	bootstrap.NewResponse(c).ToResponse(gin.H{"token":token})
	return
}


func (A Auth) Login(c *gin.Context)  {
	var param param.Login
	if err := c.ShouldBind(&param);err != nil {
		bootstrap.NewResponse(c).ToErrorResponse(global.OptionError(err.Error()))
		return
	}
	token,err := auth_service.Login(c)
	if err !=nil {
		bootstrap.NewResponse(c).ToErrorResponse(global.OptionError(err.Error()))
		return
	}
	bootstrap.NewResponse(c).ToResponse(gin.H{"token":token})
	return
}