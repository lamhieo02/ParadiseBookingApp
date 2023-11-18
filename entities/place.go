package entities

import "paradise-booking/common"

type Place struct {
	common.SQLModel
	VendorID      int     `json:"vendor_id" gorm:"column:vendor_id"`
	Name          string  `json:"name" gorm:"column:name"`
	Description   string  `json:"description" gorm:"column:description"`
	PricePerNight float64 `json:"price_per_night" gorm:"column:price_per_night"`
	Address       string  `json:"address" gorm:"column:address"`
	Capacity      int     `json:"capacity" gorm:"column:capacity"`
	Cover         string  `json:"cover" gorm:"column:cover"`
	Lat           float32 `json:"lat" gorm:"column:lat"`
	Lng           float32 `json:"lng" gorm:"column:lng"`
	Country       string  `json:"country" gorm:"column:country"`
	State         string  `json:"state" gorm:"column:state"`
	City          string  `json:"city" gorm:"column:city"`
}

func (Place) TableName() string {
	return "places"
}
