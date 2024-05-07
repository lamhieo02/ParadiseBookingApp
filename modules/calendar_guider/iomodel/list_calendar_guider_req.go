package calendarguideriomodel

type Filter struct {
	PostGuideID    int     `json:"post_guide_id" form:"post_guide_id"`
	GuiderID       int     `json:"guider_id" form:"guider_id"`
	DateFrom       string  `json:"date_from" form:"date_from"`
	DateTo         string  `json:"date_to" form:"date_to"`
	PricePerPerson float64 `json:"price_per_person" form:"price_per_person"`
	Status         *bool   `json:"status" form:"status"`
}
