package config

import "time"

//增加配置的话 需要新增 type 然后在 initSection 中新增

type rabbitMq struct {
	UserName string
	Password string
	Port     string
	Addr     string
}

type server struct {
	RunMode           string
	Port              string
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
}

type mysql struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	Port         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type redis struct {
	Host string
	Password string
	Port string
	Index int
	MaxActive int
	MaxIdle int
	IdleTimeout time.Duration
}

var (
	RabbitMqConf *rabbitMq
	ServerConf   *server
	Mysql        *mysql
	Redis		*redis
)

var section map[string]interface{}

func InitSection() {
	section = make(map[string]interface{})
	section["Server"] = &ServerConf
	section["RabbitMq"] = &RabbitMqConf
	section["Mysql"] = &Mysql
	section["Redis"] = &Redis
}
