package entities

import (
	"paradise-booking/common"
	"strconv"
)

type LikePostReview struct {
	common.SQLModel
	PostReviewId int `json:"post_review_id" gorm:"column:post_review_id"`
	AccountId    int `json:"account_id" gorm:"column:account_id"`
	Status       int `json:"status" gorm:"column:status"`
}

func (LikePostReview) TableName() string {
	return "like_post_review"
}

func (l *LikePostReview) CacheKeyNumLikePostReview() string {
	return "numlike:" + strconv.Itoa(l.PostReviewId)
}
