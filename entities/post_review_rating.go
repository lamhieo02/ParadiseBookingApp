package entities

import "paradise-booking/common"

type PostReviewRating struct {
	common.SQLModel
	PostReviewId int  `json:"post_review_id" gorm:"column:post_review_id"`
	IsLike       bool `json:"is_like" gorm:"column:is_like"`
	CommentID    int  `json:"comment_id" gorm:"column:comment_id"`
}

func (PostReviewRating) TableName() string {
	return "post_review_rating"
}
