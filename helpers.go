package main

import (
	"github.com/siafei/gin-test/bootstrap"
	"github.com/siafei/gin-test/global"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"time"
)



func SetupDBEngine() error {
	var err error
	bootstrap.DBEngine,err = bootstrap.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func SetupLogger() error {
	bootstrap.Log = bootstrap.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath+"/"+global.AppSetting.LogFileName+global.AppSetting.LogFileExt,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	},"",log.LstdFlags).WithCaller(2)
	return nil
}

//将配置赋值到全局变量中
func SetupSetting() error  {
	setting , err := bootstrap.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server",&global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App",&global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database",&global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT",&global.JwtSetting)
	if err != nil {
		return err
	}
	global.JwtSetting.Express *= time.Second
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}