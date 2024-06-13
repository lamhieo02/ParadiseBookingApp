package iomodel

type GetPlaceResp struct {
	ID               int               `json:"id" form:"id"`
	VendorID         int               `json:"vendor_id" form:"vendor_id"`
	Name             string            `json:"name" form:"name"`
	Description      string            `json:"description" form:"description"`
	PricePerNight    float64           `json:"price_per_night" form:"price_per_night"`
	Address          string            `json:"address" form:"address"`
	Cover            string            `json:"cover" form:"cover"`
	Lat              float64           `json:"lat" form:"lat"`
	Lng              float64           `json:"lng" form:"lng"`
	Country          string            `json:"country" form:"country"`
	State            string            `json:"state" form:"state"`
	District         string            `json:"district" form:"district"`
	MaxGuest         int               `json:"max_guest" form:"max_guest"`
	Numbed           int               `json:"num_bed" form:"num_bed"`
	IsFree           bool              `json:"is_free" form:"is_free"`
	RatingAverage    float64           `json:"rating_average" form:"rating_average"`
	BedRoom          int               `json:"bed_room" form:"bed_room"`
	PostGuideRelates []PostGuideRelate `json:"post_guide_relates" form:"post_guide_relates"`
}

type PostGuideRelate struct {
	ID          int    `json:"id" form:"id"`
	TopicID     int    `json:"topic_id" form:"topic_id"`
	TopicName   string `json:"topic_name" form:"topic_name"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Cover       string `json:"cover" form:"cover"`
	Country     string `json:"country" form:"country"`
	State       string `json:"state" form:"state"`
	District    string `json:"district" form:"district"`
	Address     string `json:"address" form:"address"`
}
