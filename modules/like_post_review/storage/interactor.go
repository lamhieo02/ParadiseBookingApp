package likepostreviewstorage

import "gorm.io/gorm"

type likePostReviewStorage struct {
	db *gorm.DB
}

func NewLikePostReviewStorage(db *gorm.DB) *likePostReviewStorage {
	return &likePostReviewStorage{db: db}
}
