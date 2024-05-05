package calendarguideriomodel

type GetCalendarGuiderResp struct {
	ID          int `json:"id"`
	PostGuideID int `json:"post_guide_id"`
	GuiderID    int `json:"guider_id"`
	// Guider      Guider `json:"guider"`
	Note     string `json:"note"`
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
	Price    int    `json:"price"`
	Status   bool   `json:"status"`
}

type Guider struct {
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
