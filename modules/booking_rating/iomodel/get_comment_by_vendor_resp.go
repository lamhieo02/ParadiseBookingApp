package bookingratingiomodel

import "paradise-booking/entities"

type GetCommentByVendorResp struct {
	ListRating []GetCommentUserByVendor
	// DataVendor *entities.Account `json:"vendor"`
}

type GetCommentUserByVendor struct {
	DataRating    DataBookingRating
	DataPlace     *DataPlace       `json:"place"`
	DataPostGuide *DataPostGuide   `json:"post_guide"`
	DataUser      entities.Account `json:"user"`
}

type DataBookingRating struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	BookingId  int    `json:"booking_id" `
	ObjectId   int    `json:"object_id" `
	ObjectType int    `json:"object_type"`
	Title      string `json:"title" `
	Content    string `json:"content" `
	Rating     int    `json:"rating"`
}

type DataPlace struct {
	ID               int      `json:"id" form:"id"`
	VendorID         int      `json:"vendor_id" gorm:"column:vendor_id"`
	Name             string   `json:"name" gorm:"column:name"`
	Description      string   `json:"description" gorm:"column:description"`
	PricePerNight    float64  `json:"price_per_night" gorm:"column:price_per_night"`
	Address          string   `json:"address" gorm:"column:address"`
	Images           []string `json:"images" gorm:"-"`
	Lat              float64  `json:"lat" gorm:"column:lat"`
	Lng              float64  `json:"lng" gorm:"column:lng"`
	Country          string   `json:"country" gorm:"column:country"`
	State            string   `json:"state" gorm:"column:state"`
	District         string   `json:"district" gorm:"column:district"`
	MaxGuest         int      `json:"max_guest" gorm:"column:max_guest"`
	NumBed           int      `json:"num_bed" gorm:"column:num_bed"`
	BedRoom          int      `json:"bed_room" gorm:"column:bed_room"`
	NumPlaceOriginal int      `json:"num_place_original" gorm:"column:num_place_original"`
}

type DataPostGuide struct {
	ID          int      `json:"id" form:"id"`
	PostOwnerId int      `json:"post_owner_id" gorm:"column:post_owner_id"`
	TopicID     int      `json:"topic_id" gorm:"column:topic_id"`
	Title       string   `json:"title" gorm:"column:title"`
	Description string   `json:"description" gorm:"column:description"`
	Images      []string `json:"images" gorm:"-"`
	Lat         float64  `json:"lat" gorm:"column:lat"`
	Lng         float64  `json:"lng" gorm:"column:lng"`
	Country     string   `json:"country" gorm:"column:country"`
	State       string   `json:"state" gorm:"column:state"`
	District    string   `json:"district" gorm:"column:district"`
	Address     string   `json:"address" gorm:"column:address"`
	Languages   string   `json:"languages" gorm:"column:languages"`
	Schedule    string   `json:"schedule" gorm:"column:schedule"`
}
