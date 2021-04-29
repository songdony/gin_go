package router

import (
	"gin_go/controller"
	"gin_go/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// demo
	v1 := router.Group("/demo")
	v1.Use(middleware.Tracing())
	v1.Use(middleware.RecoveryMiddleware(), middleware.RequestLog(), middleware.IPAuthMiddleware(), middleware.ValidatorMiddleware())
	{
		controller.DemoRegister(v1)
	}

	return router
}
