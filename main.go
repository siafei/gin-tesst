package main

import (
	"github.com/gin-gonic/gin"
)




func main()  {
	r:= gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200,gin.H{"message":"hello"})
	})
	r.GET("/error", func(context *gin.Context) {
		panic("error")
	})
	r.Run()
}





