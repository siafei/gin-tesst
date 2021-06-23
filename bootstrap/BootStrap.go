package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"siafei/gin-test/app/config"
	"siafei/gin-test/internal/cache"
	"siafei/gin-test/internal/db"
	"siafei/gin-test/router"
	"sync"
	"time"
)

type App struct {
	server *http.Server
	group  sync.WaitGroup
}

var (
	Db *gorm.DB
	Redis *redis.Pool
)

func New() *App {
	return &App{}
}

func (app *App) Run() {
	app.initConfig()

	app.initServer()
	//服务启动
	app.group.Add(1)
	go app.start()

	//监听停止信号 停止服务
	app.group.Add(1)
	go app.listen()

	app.group.Wait()
}

//启动服务
func (app *App) start() {
	defer app.group.Done()
	if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

//监听停止信号
func (app *App) listen() {
	defer app.group.Done()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func (app *App) initConfig() {
	//初始化 引入 配置文件
	config.InitConfig()

	//初始化数据库
	conn,err := db.New()
	if err != nil {
		log.Fatal("Mysql Init err:", err)
	}
	Db = conn
	Redis = cache.New()
}

func (app *App) initServer() {
	gin.SetMode("release") //debug	release test
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
