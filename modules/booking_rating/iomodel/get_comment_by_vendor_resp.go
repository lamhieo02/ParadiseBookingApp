package iomodel

import "paradise-booking/entities"

type GetCommentByVendorResp struct {
	ListRating []GetCommentUserByVendor
}

type GetCommentUserByVendor struct {
	DataRating entities.BookingRating
	DataPlace  entities.Place   `json:"place"`
	DataUser   entities.Account `json:"user"`
}
