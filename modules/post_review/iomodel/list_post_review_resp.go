package postreviewiomodel

import "paradise-booking/common"

type ListPostReviewResp struct {
	Data   []PostReviewResp `json:"data"`
	Paging *common.Paging   `json:"paging"`
}
