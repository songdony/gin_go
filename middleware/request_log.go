package middleware

import "github.com/gin-gonic/gin"

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
	    // 这里可以生成traceid来实现链路追踪，将traceid放入context传递即可
		c.Next()
	}
}
