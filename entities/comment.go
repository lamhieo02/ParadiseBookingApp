package entities

import "paradise-booking/common"

type Comment struct {
	common.SQLModel
	Content   string `json:"content" gorm:"column:content"`
	Image     string `json:"image" gorm:"column:image"`
	Videos    string `json:"videos" gorm:"column:videos"`
	AccountID int64  `json:"account_id" gorm:"column:account_id"`
}

func (Comment) TableName() string {
	return "comments"
}
