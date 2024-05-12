package bookingguiderstorage

import "gorm.io/gorm"

type bookingGuiderSto struct {
	db *gorm.DB
}

func NewBookingGuiderStorage(db *gorm.DB) *bookingGuiderSto {
	return &bookingGuiderSto{db}
}
