package entities

import (
	"paradise-booking/common"
	"strconv"
)

type PostGuide struct {
	common.SQLModel
	PostOwnerId int     `json:"post_owner_id" gorm:"column:post_owner_id"`
	TopicID     int     `json:"topic_id" gorm:"column:topic_id"`
	Title       string  `json:"title" gorm:"column:title"`
	Description string  `json:"description" gorm:"column:description"`
	Cover       string  `json:"cover" gorm:"column:cover"`
	Lat         float64 `json:"lat" gorm:"column:lat"`
	Lng         float64 `json:"lng" gorm:"column:lng"`
	Country     string  `json:"country" gorm:"column:country"`
	State       string  `json:"state" gorm:"column:state"`
	District    string  `json:"district" gorm:"column:district"`
	Address     string  `json:"address" gorm:"column:address"`
	Languages   string  `json:"languages" gorm:"column:languages"`
}

func (PostGuide) TableName() string {
	return "post_guide"
}

func (p PostGuide) CacheKey() string {
	return "post_guide:" + strconv.Itoa(p.Id)
}

func (p PostGuide) CacheKeyGuideRating() string {
	return "post_guide_rating:" + strconv.Itoa(p.Id)
}
