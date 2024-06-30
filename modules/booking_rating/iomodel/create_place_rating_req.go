package bookingratingiomodel

type CreateBookingRatingReq struct {
	BookingID int `json:"booking_id"`
	// PlaceID   int     `json:"place_id"`
	ObjectID   int     `json:"object_id"`
	ObjectType int     `json:"object_type"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Rating     float64 `json:"rating"`
}
