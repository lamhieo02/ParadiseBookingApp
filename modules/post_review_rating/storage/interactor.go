package postreviewratingstorage

import "gorm.io/gorm"

type postReviewRatingStorage struct {
	db *gorm.DB
}

func NewPostReviewRatingStorage(db *gorm.DB) *postReviewRatingStorage {
	return &postReviewRatingStorage{db: db}
}
