package entities

import "paradise-booking/common"

type BookingRating struct {
	common.SQLModel
	UserId    int    `json:"user_id" gorm:"column:user_id"`
	BookingId int    `json:"booking_id" gorm:"column:booking_id"`
	Title     string `json:"title" gorm:"column:title"`
	Content   string `json:"content" gorm:"column:content"`
	Rating    int    `json:"rating" gorm:"column:rating"`
}

func (BookingRating) TableName() string {
	return "booking_rating"
}
