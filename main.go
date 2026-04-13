package main

import (
	"os"
	"os/signal"
	"syscall"

	"goravel/bootstrap"
	serverfasthttp "goravel/server/fasthttp"
	servergin "goravel/server/gin"
	serverhertz "goravel/server/hertz"
)

func main() {
	app := bootstrap.Boot()
	app.Start()

	go servergin.Start()
	go serverhertz.Start()
	go serverfasthttp.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
