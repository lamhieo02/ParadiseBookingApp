package bookingratingiomodel

import "paradise-booking/entities"

type GetCommentByUserResp struct {
	ListRating []GetCommentRespByUser
	DataUser   entities.Account `json:"user"`
}

type GetCommentRespByUser struct {
	DataRating    *DataBookingRating
	DataPlace     *DataPlace     `json:"place"`
	DataPostGuide *DataPostGuide `json:"post_guide"`
}
