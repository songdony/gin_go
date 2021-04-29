package middleware

import (
	"errors"
	"fmt"
	"gin_go/lib"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

// RecoveryMiddleware捕获所有panic，并且返回错误信息  兜底
func RecoveryMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func(){
			if err := recover(); err != nil {
				fmt.Println(string(debug.Stack()))

				if lib.ConfBase.DebugMode != "debug" {
					ResponseError(context, 500, errors.New("内部错误"))
					return
				} else {
					ResponseError(context, 500, errors.New(fmt.Sprint(err)))
					return
				}

			}
		}()
		context.Next()
	}
}
