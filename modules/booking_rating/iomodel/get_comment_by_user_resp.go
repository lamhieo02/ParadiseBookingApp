package iomodel

import "paradise-booking/entities"

type GetCommentByUserResp struct {
	DataRating []entities.BookingRating
	DataUser   entities.Account `json:"user"`
}
