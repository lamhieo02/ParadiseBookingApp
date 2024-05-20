package entities

import "paradise-booking/common"

type BookingRating struct {
	common.SQLModel
	UserId     int    `json:"user_id" gorm:"column:user_id"`
	BookingId  int    `json:"booking_id" gorm:"column:booking_id"`
	ObjectId   int    `json:"object_id" gorm:"column:object_id"`
	ObjectType int    `json:"object_type" gorm:"column:object_type"`
	Title      string `json:"title" gorm:"column:title"`
	Content    string `json:"content" gorm:"column:content"`
	Rating     int    `json:"rating" gorm:"column:rating"`
}

func (BookingRating) TableName() string {
	return "booking_rating"
}

type StatisticResp struct {
	Rating float64 `gorm:"column:rating" json:"rating"`
	Count  int64   `gorm:"column:count" json:"count"`
}
