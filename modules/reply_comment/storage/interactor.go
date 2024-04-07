package replycommentstorage

import "gorm.io/gorm"

type replyCommentStorage struct {
	db *gorm.DB
}

func NewReplyCommentStorage(db *gorm.DB) *replyCommentStorage {
	return &replyCommentStorage{db: db}
}
