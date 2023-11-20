package iomodel

type UpdatePlaceReq struct {
	Name          string  `json:"name" form:"name"`
	Description   string  `json:"description" form:"description"`
	PricePerNight float64 `json:"price_per_night" form:"price_per_night"`
	Address       string  `json:"address" form:"address"`
	Capacity      int     `json:"capacity" form:"capacity"`
	Cover         string  `json:"cover" form:"cover"`
	Lat           float32 `json:"lat" form:"lat"`
	Lng           float32 `json:"lng" form:"lng"`
	Country       string  `json:"country" form:"country"`
	State         string  `json:"state" form:"state"`
	City          string  `json:"city" form:"city"`
}