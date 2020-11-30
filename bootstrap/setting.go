package bootstrap

import (
	"github.com/spf13/viper"
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

func (s *Setting) ReadSection(k string,v interface{}) error  {
	err := s.vp.UnmarshalKey(k,v)
	if err != nil {
		return  err
	}
	return nil
}
