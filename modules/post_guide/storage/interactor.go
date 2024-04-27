package postguidestorage

import "gorm.io/gorm"

type postGuideStorage struct {
	db *gorm.DB
}

func NewPostGuideStorage(db *gorm.DB) *postGuideStorage {
	return &postGuideStorage{db: db}
}
