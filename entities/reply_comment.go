package entities

import "paradise-booking/common"

type ReplyComment struct {
	common.SQLModel
	SourcecommentID int    `json:"source_comment_id" gorm:"column:source_comment_id;"`
	Content         string `json:"content" gorm:"column:content;"`
	Image           string `json:"image" gorm:"column:image;"`
	Videos          string `json:"videos" gorm:"column:videos;"`
	AccountID       int    `json:"account_id" gorm:"column:account_id;"`
}

func (ReplyComment) TableName() string {
	return "reply_comments"
}
