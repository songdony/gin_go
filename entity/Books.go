package entity

type Books struct{
	BookID int `gorm:"column:book_id;AUTO_INCREMENT;PRIMARY_KEY"`
	BookName string `gorm:"column:book_name;type:varchar(50)"`
	BookIntr string `gorm:"column:book_intr;type:text"`
	BookPrice1 float64 `gorm:"column:book_price1;type:decimal"`
	BookPrice2 float64 `gorm:"column:book_price2;type:decimal"`
	BookAuthor string `gorm:"column:book_author;type:varchar(50)"`
	BookPress string `gorm:"column:book_press;type:varchar(50)"`
	BookDate string `gorm:"column:book_date;type:varchar(50)"`
	BookKind int `gorm:"column:book_kind;type:int"`
}
type BookList []*Books
