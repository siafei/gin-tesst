package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"siafei/gin-test/app/config"
)

func New() (*gorm.DB,error) {
	db, err := gorm.Open(config.Mysql.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.DBName,
		config.Mysql.Charset,
		config.Mysql.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if config.ServerConf.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)	//让grom转义struct名字的时候不用加上s
	db.DB().SetMaxIdleConns(config.Mysql.MaxIdleConns)	//设置最大空闲时间的连接数
	db.DB().SetMaxOpenConns(config.Mysql.MaxOpenConns)	//设置最大连接数
	return db, nil
}
