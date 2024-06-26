package entities

import "paradise-booking/common"

type PostReview struct {
	common.SQLModel
	PostOwnerId int     `json:"post_owner_id" gorm:"column:post_owner_id"`
	Title       string  `json:"title" gorm:"column:title"`
	Topic       int     `json:"topic" gorm:"column:topic"`
	Content     string  `json:"content" gorm:"column:content"`
	Image       string  `json:"image" gorm:"column:img"`
	Videos      string  `json:"videos" gorm:"column:videos"`
	Lat         float64 `json:"lat" gorm:"column:lat"`
	Lng         float64 `json:"lng" gorm:"column:lng"`
	Country     string  `json:"country" gorm:"column:country"`
	State       string  `json:"state" gorm:"column:state"`
	District    string  `json:"district" gorm:"column:district"`
}

func (PostReview) TableName() string {
	return "post_review"
}
