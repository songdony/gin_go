package controller

import (
	"fmt"
	"gin_go/entity"
	"gin_go/lib"
	"gin_go/middleware"
	"gin_go/service"
	"github.com/gin-gonic/gin"
)

type DemoController struct {
}



func DemoRegister(router *gin.RouterGroup) {
	demo := DemoController{}
	router.GET("/index", demo.Index)
	router.GET("/detail", demo.Detail)
}

func (demo *DemoController) Index(c *gin.Context) {
	middleware.ResponseSuccess(c, "hello")
	return
}


func (demo *DemoController) Detail(c *gin.Context) {
	params := &service.DetailInput{}
	ret := &entity.Books{}
	//params.BookId = 220

	if err := params.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	fmt.Println("params=",params)
	ret,err = (&service.BookService{}).LoadBookList(c,tx,params.BookId)

	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}

	middleware.ResponseSuccess(c, ret)
	return
}