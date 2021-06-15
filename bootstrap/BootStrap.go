package bootstrap

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"siafei/gin-test/app/config"
	"siafei/gin-test/router"
)

type App struct {
	server *http.Server
}

func New() *App {
	return &App{}
}

func (app *App) Run() {
	app.InitConfig()
	app.InitServer()
	app.server.ListenAndServe()
}

func (app *App) InitConfig() {
	//初始化 引入 配置文件
	config.InitConfig()
}

func (app *App) InitServer() {
	gin.SetMode("debug") //debug	release test
	r := router.New()
	s := &http.Server{
		Addr:              ":" + config.ServerConf.Port,
		Handler:           r,
		ReadHeaderTimeout: config.ServerConf.ReadHeaderTimeout,
		WriteTimeout:      config.ServerConf.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}
	app.server = s
}
