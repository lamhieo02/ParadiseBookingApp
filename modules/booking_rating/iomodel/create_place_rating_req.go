package iomodel

type CreateBookingRatingReq struct {
	BookingID int     `json:"booking_id"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Rating    float64 `json:"rating"`
}
