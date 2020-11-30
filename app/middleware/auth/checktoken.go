package middleware

import (
	"github.com/gin-gonic/gin"
	service "github.com/siafei/gin-test/app/service/auth"
	"github.com/siafei/gin-test/bootstrap"
	"github.com/siafei/gin-test/global"
)

//登录检测
func CheckToken() gin.HandlerFunc  {
	return func(c *gin.Context) {
		var auth_service service.Auth
		_,err := auth_service.CheckToken(c)
		if err !=nil {
			bootstrap.NewResponse(c).ToErrorResponse(global.OptionError(err.Error()))
			c.Abort()
			return
		}
		c.Next()
	}
}