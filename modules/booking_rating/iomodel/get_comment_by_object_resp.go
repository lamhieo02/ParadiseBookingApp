package bookingratingiomodel

import "paradise-booking/entities"

type GetCommentByObjectResp struct {
	ListRating    []GetCommentRespByObject
	DataPlace     *DataPlace     `json:"place"`
	DataPostGuide *DataPostGuide `json:"post_guide"`
}

type GetCommentRespByObject struct {
	DataRating DataBookingRating
	DataUser   entities.Account `json:"user"`
}
