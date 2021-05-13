package service

import (
	"gin_go/common"
	"gin_go/entity"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BookService struct {
}

type DetailInput struct {
	BookId int `form:"book_id" json:"book_id" comment:"书的id" validate:"required" example:"10"`
}

func (params *DetailInput) BindingValidParams(c *gin.Context) error {
	return common.DefaultGetValidParams(c, params)
}

func (this *BookService) LoadBookDetail(ctx *gin.Context, tx *gorm.DB, id int) (*entity.Books, error) {
	books := &entity.Books{}

	err := tx.Where("book_id = ?", id).Find(books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (this *BookService) LoadBookLists(ctx *gin.Context, tx *gorm.DB) (*entity.BookList, error) {
	bookLists := &entity.BookList{}

	err := tx.Where("book_id in (1,2,3)").Find(bookLists).Error
	if err != nil {
		return nil, err
	}

	return bookLists, nil
}
