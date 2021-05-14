package entity

type Books struct {
	BookID     int     `json:"book_id" gorm:"column:book_id;AUTO_INCREMENT;PRIMARY_KEY"`
	BookName   string  `json:"book_name" gorm:"column:book_name;type:varchar(50)"`
	BookIntr   string  `json:"book_intr" gorm:"column:book_intr;type:text"`
	BookPrice1 float64 `json:"book_price1" gorm:"column:book_price1;type:decimal"`
	BookPrice2 float64 `json:"book_price2" gorm:"column:book_price2;type:decimal"`
	BookAuthor string  `json:"book_author" gorm:"column:book_author;type:varchar(50)"`
	BookPress  string  `json:"book_press" gorm:"column:book_press;type:varchar(50)"`
	BookDate   string  `json:"book_date" gorm:"column:book_date;type:varchar(50)"`
	BookKind   int     `json:"book_kind" gorm:"column:book_kind;type:int"`
}
type BookList []*Books
