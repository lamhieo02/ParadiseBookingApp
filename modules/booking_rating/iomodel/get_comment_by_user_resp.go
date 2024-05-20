package iomodel

import "paradise-booking/entities"

type GetCommentByUserResp struct {
	ListRating []GetCommentRespByUser
	DataUser   entities.Account `json:"user"`
}

type GetCommentRespByUser struct {
	DataRating    *entities.BookingRating
	DataPlace     *entities.Place     `json:"place"`
	DataPostGuide *entities.PostGuide `json:"post_guide"`
}
