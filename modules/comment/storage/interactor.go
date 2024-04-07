package commentstorage

import "gorm.io/gorm"

type commentStorage struct {
	db *gorm.DB
}

func NewCommentStorage(db *gorm.DB) *commentStorage {
	return &commentStorage{db: db}
}
