package iomodel

import "paradise-booking/entities"

type GetCommentByObjectResp struct {
	ListRating    []GetCommentRespByObject
	DataPlace     *entities.Place     `json:"place"`
	DataPostGuide *entities.PostGuide `json:"post_guide"`
}

type GetCommentRespByObject struct {
	DataRating entities.BookingRating
	DataUser   entities.Account `json:"user"`
}
