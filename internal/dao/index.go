package dao

import (
	"alarm/internal/model"
	"github.com/jinzhu/gorm"
)

type IIndexDAO interface {
	GetIndex() (*model.Index, error)
}

type indexDAO struct {
	db *gorm.DB
}

func NewIndexDAO(db *gorm.DB) IIndexDAO {
	return &indexDAO{db: db}
}

func (d *indexDAO) GetIndex() (*model.Index, error) {
	var index model.Index
	if err := d.db.First(&index).Error; err != nil {
		return nil, err
	}
	return &index, nil
}