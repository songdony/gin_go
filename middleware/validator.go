package middleware

import (
	"gin_scaffold/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

// 使用的是v9过滤器
func ValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		val := validator.New()

		val.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("comment")
		})

		//自定义验证方法
		//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
		val.RegisterValidation("is-validuser", func(fl validator.FieldLevel) bool {
			return fl.Field().String() == "admin"
		})

		c.Set(common.ValidatorKey, val)

		c.Next()
	}


}
