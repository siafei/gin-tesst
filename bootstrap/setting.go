package bootstrap

import (
	"github.com/siafei/gin-test/global"
	"github.com/spf13/viper"
	"time"
)

type Setting struct {
	vp *viper.Viper
}

//引入配置文件
func NewSetting() (*Setting,error)  {
	vp := viper.New()
	vp.SetConfigName("config")	//设置配置文件名字
	vp.AddConfigPath("config/")	//设置配置文件的路径
	vp.SetConfigType("yaml")	//设置配置文件类型
	err := vp.ReadInConfig()	// 搜索路径，并读取配置数据
	if err != nil {
		return nil, err
	}
	return &Setting{vp},nil
}

//将配置赋值到全局变量中
func SetupSetting() error  {
	setting , err := NewSetting()
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

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}

func (s *Setting) ReadSection(k string,v interface{}) error  {
	err := s.vp.UnmarshalKey(k,v)
	if err != nil {
		return  err
	}
	return nil
}
