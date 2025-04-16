package dao

import "github.com/jinzhu/gorm"

var (
	IndexDAO IIndexDAO
)

func Init(db *gorm.DB) {
	IndexDAO = NewIndexDAO(db)
}
