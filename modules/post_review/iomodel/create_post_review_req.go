package postreviewiomodel

type CreatePostReviewReq struct {
	AccountID int64    `json:"account_id"`
	Title     string   `json:"title"`
	Topic     int      `json:"topic"`
	Content   string   `json:"content"`
	Images    []string `json:"images"`
	Videos    []string `json:"videos"`
	Lat       float64  `json:"lat"`
	Lng       float64  `json:"lng"`
}
