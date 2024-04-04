package postreviewiomodel

type CreatePostReviewReq struct {
	AccountID int64   `json:"account_id"`
	Title     string  `json:"title"`
	Topic     string  `json:"topic"`
	Content   string  `json:"content"`
	Image     string  `json:"image"`
	Videos    string  `json:"videos"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
}
