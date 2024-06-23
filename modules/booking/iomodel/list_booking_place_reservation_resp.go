package iomodel

type ListBookingPlaceReservationResp struct {
	Data []BookingPlaceResp `json:"data"`
}

type BookingPlaceResp struct {
	*DataPlace
	IsBooked bool `json:"is_booked"`
}
