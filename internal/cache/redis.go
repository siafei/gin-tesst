package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"siafei/gin-test/app/config"
)

func New() *redis.Pool {
	redis_addr := fmt.Sprintf("%s:%s",
		config.Redis.Host,
		config.Redis.Port,
	)
	setdb := redis.DialDatabase(config.Redis.Index)    // 设置db
	setPasswd := redis.DialPassword(config.Redis.Password) // 设置redis连接密码
	return &redis.Pool{
		MaxIdle: config.Redis.MaxIdle,
		MaxActive: config.Redis.MaxActive,
		IdleTimeout: config.Redis.IdleTimeout,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) { return redis.Dial("tcp", redis_addr,setdb,setPasswd) },
	}
}
