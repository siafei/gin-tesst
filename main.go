package main

import "siafei/gin-test/internal/rabbitmq"

func main()  {
	//gin.SetMode("debug") //debug	release test
	//r := router.New()
	//s := &http.Server{
	//	Addr: ":8080",
	//	Handler: r,
	//	ReadHeaderTimeout: 10*time.Second,
	//	WriteTimeout: 10*time.Second,
	//	MaxHeaderBytes: 1<<20,
	//}
	//s.ListenAndServe()
	rabbitmq.Test()
}





