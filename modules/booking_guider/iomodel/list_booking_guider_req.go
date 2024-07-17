package bookingguideriomodel

type Filter struct {
	Statuses    []int  `json:"statuses"`
	DateFrom    string `json:"date_from"`
	DateTo      string `json:"date_to"`
	PostGuideID int    `json:"post_guide_id"`
	GuiderID    int    `json:"guider_id"`
	UserID      int    `json:"user_id"`
}
