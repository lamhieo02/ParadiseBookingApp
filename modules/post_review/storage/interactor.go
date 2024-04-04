package postreviewstorage

import "gorm.io/gorm"

type postReviewStorage struct {
	db *gorm.DB
}

func NewPostReviewStorage(db *gorm.DB) *postReviewStorage {
	return &postReviewStorage{db: db}
}
