package placewishlistiomodel

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
