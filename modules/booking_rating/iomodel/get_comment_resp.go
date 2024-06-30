package bookingratingiomodel

import "paradise-booking/entities"

type GetCommentResp struct {
	DataRating    DataBookingRating
	DataUser      entities.Account `json:"user"`
	DataPlace     *DataPlace       `json:"place"`
	DataPostGuide *DataPostGuide   `json:"post_guide"`
}
