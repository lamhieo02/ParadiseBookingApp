package entities

import "paradise-booking/common"

type PostReview struct {
	common.SQLModel
	PostOwnerId int     `json:"post_owner_id" gorm:"column:post_owner_id"`
	Title       string  `json:"title" gorm:"column:title"`
	Topic       string  `json:"topic" gorm:"column:topic"`
	Content     string  `json:"content" gorm:"column:content"`
	Img         string  `json:"img" gorm:"column:img"`
	Lat         float64 `json:"lat" gorm:"column:lat"`
	Lng         float64 `json:"lng" gorm:"column:lng"`
}

func (PostReview) TableName() string {
	return "post_review"
}
