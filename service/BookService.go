package service

import (
	"gin_go/entity"
	"gin_go/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BookService struct {
}

type DetailInput struct {
	BookId int    `form:"book_id" json:"book_id" comment:"书的id" validate:"required" example:"10"`
}

func (params *DetailInput) BindingValidParams(c *gin.Context) error {
	return common.DefaultGetValidParams(c, params)
}


func(this *BookService) LoadBookList(ctx *gin.Context,tx *gorm.DB, id int) (*entity.Books,error) {
	panic("sss")
	books:=&entity.Books{}

	err := tx.Where("book_id = ?", id).Find(books).Error
	if err != nil {
		return nil, err
	}

	return books,nil
}
