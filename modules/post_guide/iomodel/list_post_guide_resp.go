package postguideiomodel

import "paradise-booking/common"

type ListPostGuideResp struct {
	Data   []GetPostGuideResp `json:"data"`
	Paging *common.Paging     `json:"paging"`
}
