package main

import (
	"siafei/gin-test/bootstrap"
)

func main() {
	app := bootstrap.New()
	app.Run()
}