package iomodel

import (
	"paradise-booking/entities"
	"time"
)

type ListBookingResp struct {
	UserId   int               `json:"user_id"`
	User     entities.Account  `json:"user"`
	ListData []DataListBooking `json:"data"`
}

type DataListBooking struct {
	Id          int            `json:"id" gorm:"column:id"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	PlaceId     int            `json:"place_id"`
	Place       entities.Place `json:"place"`
	StatusId    int            `json:"status_id"`
	ChekoutDate string         `json:"checkout_date"`
	CheckInDate string         `json:"checkin_date"`
}
