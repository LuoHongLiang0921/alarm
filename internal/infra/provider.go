package infra

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Provider struct {
	DB *gorm.DB
}

func InitProvider() (*Provider, func()) {
	db, bizErr := connectMySQL()
	if bizErr != nil {
		return nil, func() {}
	}
	return &Provider{
			DB: db,
		}, func() {
			db.Close()
		}
}

func connectMySQL() (*gorm.DB, error) {
	// 初始化数据库连接
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test01?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return nil, err
	}
	// defer db.Close()
	return db, nil
}
