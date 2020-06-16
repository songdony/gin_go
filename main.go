package main

import (
	"fmt"
	"gin_scaffold/lib"
	"gin_scaffold/router"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var TimeFormat = "2006-01-02 15:04:05"

func main(){

	// 解析配置文件目录
	if err := lib.ParseConfPath("conf/dev/base"); err != nil {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " ParseConfPath:"+err.Error())
	}

	//初始化配置文件
	if err := lib.InitViperConf(); err != nil {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitViperConf:"+err.Error())
	}

	if err := lib.InitBaseConf(lib.GetConfPath("conf/dev/base")); err != nil {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitBaseConf:"+err.Error())
	}

	if err := lib.InitDBPool(lib.GetConfPath("conf/dev/mysql")); err != nil {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitDBPool:"+err.Error())
	}

	if err := lib.InitRedisConf(lib.GetConfPath("conf/dev/redis")); err != nil {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitRedis:"+err.Error())
	}


	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()

}