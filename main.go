package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siafei/gin-test/global"
	"github.com/siafei/gin-test/routers"
	"log"
	"net/http"
)

func init()  {
	err := SetupSetting()  //引入配置文件
	if err != nil {
		log.Fatalf("init setupSetting err: %v",err)
	}
	err = SetupDBEngine()
	if err != nil {
		log.Fatalf("init DBEngine err: %v",err)
	}
	err = SetupLogger()
	if err != nil {
		log.Fatalf("init Logger err: %v",err)
	}
}


func main()  {
	gin.SetMode(global.ServerSetting.RunModel)   //设置运行模式
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":"+global.ServerSetting.HttpPort,
		Handler: router,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeOut,
		WriteTimeout: global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1<<20,
	}
	s.ListenAndServe()
}





