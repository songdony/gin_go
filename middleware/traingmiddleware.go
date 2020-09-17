package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := GetURI(c.Request)
		if list := strings.Split(url, "?"); len(list) > 0 {
			url = list[0]
		}

		span, ctx := opentracing.StartSpanFromContext(c, url)
		c.Request.WithContext(ctx)

		c.Next()

		defer span.Finish()
	}
}

func GetURI(r *http.Request) string {
	reqURI := r.RequestURI
	if reqURI == "" {
		reqURI = r.URL.RequestURI()
	}
	return reqURI
}
