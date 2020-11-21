package config

import (
	"time"
)

//服务器设置
type ServerSettings struct {
	RunModel string
	HttpPort string
	ReadTimeOut time.Duration
	WriteTimeOut time.Duration
}

//应用设置
type AppSettings struct {
	DefaultPageSize int
	MaxPageSize int
	LogSavePath string
	LogFileName string
	LogFileExt string
}

type DatabaseSetting struct {
	DBType string
	UserName string
	Password string
	Host string
	Port string
	DBName string
	TablePrefix string
	Charset string
	ParseTime bool
	MaxIdleConns int
	MaxOpenConns int
}
