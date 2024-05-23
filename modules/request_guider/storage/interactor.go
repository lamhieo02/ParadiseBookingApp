package requestguiderstorage

import "gorm.io/gorm"

type RequestGuiderSto struct {
	db *gorm.DB
}

func NewRequestGuiderStorage(db *gorm.DB) *RequestGuiderSto {
	return &RequestGuiderSto{
		db: db,
	}
}
