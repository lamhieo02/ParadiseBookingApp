package postreviewiomodel

type UpdatePostReviewReq struct {
	PostReviewID int64    `json:"post_review_id"`
	AccountID    int64    `json:"account_id"`
	Title        string   `json:"title"`
	Topic        int      `json:"topic"`
	Content      string   `json:"content"`
	Images       []string `json:"images"`
	Videos       []string `json:"videos"`
	Lat          float64  `json:"lat"`
	Lng          float64  `json:"lng"`
	Country      string   `json:"country"`
	State        string   `json:"state"`
	District     string   `json:"district"`
}
