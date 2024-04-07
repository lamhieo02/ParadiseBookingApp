package likepostreviewiomodel

type LikePostReviewReq struct {
	AccountID    int64 `json:"account_id"`
	PostReviewID int   `json:"post_review_id"`
	Type         int   `json:"type"`
}
