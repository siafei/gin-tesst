package bootstrap

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/siafei/gin-test/config"
	"github.com/siafei/gin-test/global"
)

func NewDBEngine(databaseSetting *config.DatabaseSetting) (*gorm.DB,error)  {
	db,err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.Port,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err!=nil {
		return nil, err
	}
	if global.ServerSetting.RunModel == "debug" {
		//设置日志级别
		db.LogMode(true)
	}
	db.SingularTable(true)	//使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)	//设置最大闲置数量
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)	//设置最大打开连接数
	return db,nil
}

func SetupDBEngine() error {
	var err error
	global.DBEngine,err = NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
