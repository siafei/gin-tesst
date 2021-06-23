package config

import (
	"github.com/spf13/viper"
	"time"
)

var vp *viper.Viper

/**
初始化config
*/
func InitConfig() {
	vp = viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("./")
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		panic(err)
	}
	InitSection()
	readSection()
	initTimeConf()
}

/**
将变量赋值到config下的变量中
*/
func readSection() {
	for key, val := range section {
		_ = vp.UnmarshalKey(key, val)
	}
}

/**
时间的配置另外初始化
*/
func initTimeConf() {
	ServerConf.ReadHeaderTimeout *= time.Second
	ServerConf.WriteTimeout *= time.Second
	Redis.IdleTimeout *= time.Second
}
