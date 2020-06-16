package middleware

import (
	"github.com/gin-gonic/gin"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 黑白名单的判断，黑名单不让进
		c.Next()
	}
}
