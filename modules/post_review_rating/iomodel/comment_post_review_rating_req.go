package postreviewratingiomodel

type CommentPostReviewRatingReq struct {
	AccountID    int64  `json:"account_id"`
	PostReviewID int    `json:"post_review_id"`
	Comment      string `json:"comment"`
	Image        string `json:"image"`
	Videos       string `json:"video"`
}
