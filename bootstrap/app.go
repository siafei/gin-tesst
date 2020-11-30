package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/siafei/gin-test/global"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func NewResponse(ctx *gin.Context) *Response{
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data interface{})  {
	if data ==nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK,data)
}

func (r *Response) ToResponseList(list interface{},total int)  {
	r.Ctx.JSON(http.StatusOK,gin.H{
		"list":list,
		"pager":Pager{
			Page: GetPage(r.Ctx),
			Limit: GetLimit(r.Ctx),
			Total: total,
		},
	})
}

func (r *Response) ToErrorResponse(err *global.Error)  {
	response := gin.H{"code":err.Code,"msg":err.Msg}
	r.Ctx.JSON(err.StatusCode,response)
}