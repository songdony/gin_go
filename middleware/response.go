package middleware

import (
	"encoding/json"
	"fmt"
	"gin_go/lib"
	"github.com/gin-gonic/gin"
	"strings"
)

type ResponseCode int

//1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode ResponseCode = iota
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode

	InvalidRequestErrorCode ResponseCode = 401
	CustomizeCode           ResponseCode = 1000

	GROUPALL_SAVE_FLOWERROR ResponseCode = 2001
)

type Response struct {
	ErrorCode ResponseCode `json:"errno"`
	ErrorMsg  string       `json:"errmsg"`
	Data      interface{}  `json:"data"`
	TraceId   interface{}  `json:"trace_id"`
	Stack     interface{}  `json:"stack"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {

	stack := ""
	if c.Query("is_debug") == "1" || lib.GetConfEnv() == "dev" {
		stack = strings.Replace(fmt.Sprintf("%+v", err), err.Error()+"\n", "", -1)
	}

	resp := &Response{ErrorCode: code, ErrorMsg: err.Error(), Data: "",Stack: stack}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{ErrorCode: SuccessCode, ErrorMsg: "", Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

