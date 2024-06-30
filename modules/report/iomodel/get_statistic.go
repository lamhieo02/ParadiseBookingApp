package reportiomodel

type GetStatisticPlaceReq struct {
	DateFrom string `json:"date_from" form:"date_from" binding:"required"`
	DateTo   string `json:"date_to" form:"date_to" binding:"required"`
	Type     int    `json:"type" form:"type" binding:"required"`
	PlaceID  string `json:"place_id" form:"place_id"`
}

type StatisticPlaceResp struct {
	TotalRevenue        float64            `json:"total_revenue"`
	TotalBookingSuccess int                `json:"total_booking_success"`
	TotalBookingCancel  int                `json:"total_booking_cancel"`
	StatisticBooking    []StatisticBooking `json:"statistic_booking"`
	StatisticRevenue    []StatisticRevenue `json:"statistic_revenue"`
}

type StatisticBooking struct {
	BookingSuccess int    `json:"booking_success"`
	BookingCancel  int    `json:"booking_cancel"`
	ColumnName     string `json:"column_name"`
}

type StatisticRevenue struct {
	Revenue    float64 `json:"revenue"`
	ColumnName string  `json:"column_name"`
}
