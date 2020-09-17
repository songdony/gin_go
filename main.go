package main

import (
	"gin_scaffold/lib"
	"gin_scaffold/router"
	"os"
	"os/signal"
	"syscall"
)

var TimeFormat = "2006-01-02 15:04:05"

func main() {

	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})

	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()

}
