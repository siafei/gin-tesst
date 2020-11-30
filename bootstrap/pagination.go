package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/siafei/gin-test/global"
)

func GetPage(c *gin.Context) int  {
	page := StrTo(c.Query("page")).MustInt()
	if page <=0 {
		return 1
	}
	return page
}

func GetLimit(c *gin.Context) int  {
	limit := StrTo(c.Query("limit")).MustInt()
	if limit<=0 {
		return global.AppSetting.DefaultPageSize
	}
	if limit >global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return limit
}

func GetPageOffset(page,limit int) int{
	result := 0
	if page>0 {
		result = (page-1) * limit
	}
	return result
}

