package main

import (
	"github.com/siafei/gin-test/routers"
	"net/http"
	"time"
)

func main()  {
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadHeaderTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		MaxHeaderBytes: 1<<20,
	}
	s.ListenAndServe()
}


