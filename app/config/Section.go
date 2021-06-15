package config

import "time"

//增加配置的话 需要新增 type 然后在 initSection 中新增

type rabbitMq struct {
	UserName string
	Password string
	Port     int
	Addr     string
}

type server struct {
	Port              string
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
}

var (
	RabbitMqConf *rabbitMq
	ServerConf   *server
)

var section map[string]interface{}

func InitSection() {
	section = make(map[string]interface{})
	section["Server"] = &ServerConf
	section["RabbitMq"] = &RabbitMqConf
}
