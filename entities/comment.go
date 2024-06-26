package entities

import (
	"paradise-booking/common"
	"strconv"
)

type Comment struct {
	common.SQLModel
	Content      string `json:"content" gorm:"column:content"`
	Image        string `json:"image" gorm:"column:image"`
	Videos       string `json:"videos" gorm:"column:videos"`
	AccountID    int64  `json:"account_id" gorm:"column:account_id"`
	PostReviewID int64  `json:"post_review_id" gorm:"column:post_review_id"`
}

func (Comment) TableName() string {
	return "comments"
}

func (c *Comment) CacheKeyNumCommentByPostReview() string {
	return "num_cmt" + strconv.Itoa(int(c.PostReviewID))
}
