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

type DataPlace struct {
	Id               int        `json:"id"`
	VendorID         int        `json:"vendor_id" gorm:"column:vendor_id"`
	Name             string     `json:"name" gorm:"column:name"`
	Description      string     `json:"description" gorm:"column:description"`
	PricePerNight    float64    `json:"price_per_night" gorm:"column:price_per_night"`
	Address          string     `json:"address" gorm:"column:address"`
	Images           []string   `json:"images" gorm:"column:images"`
	Lat              float64    `json:"lat" gorm:"column:lat"`
	Lng              float64    `json:"lng" gorm:"column:lng"`
	Country          string     `json:"country" gorm:"column:country"`
	State            string     `json:"state" gorm:"column:state"`
	District         string     `json:"district" gorm:"column:district"`
	MaxGuest         int        `json:"max_guest" gorm:"column:max_guest"`
	NumBed           int        `json:"num_bed" gorm:"column:num_bed"`
	BedRoom          int        `json:"bed_room" gorm:"column:bed_room"`
	NumPlaceOriginal int        `json:"num_place_original" gorm:"column:num_place_original"`
	CreatedAt        *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type DataListBooking struct {
	Id              int        `json:"id" gorm:"column:id"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	PlaceId         int        `json:"place_id"`
	Place           *DataPlace `json:"place"`
	StatusId        int        `json:"status_id"`
	ChekoutDate     string     `json:"checkout_date"`
	CheckInDate     string     `json:"checkin_date"`
	GuestName       string     `json:"guest_name"`
	TotalPrice      float64    `json:"total_price"`
	ContentToVendor string     `json:"content_to_vendor"`
	NumberOfGuest   int        `json:"number_of_guest"`
	PaymentMethod   int        `json:"payment_method"`
}
