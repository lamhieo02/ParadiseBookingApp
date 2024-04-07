package replycommentiomodel

type ReplyCommentReq struct {
	SourceCommentID int    `json:"source_comment_id"`
	Content         string `json:"content"`
	Image           string `json:"image"`
	Videos          string `json:"videos"`
	AccountID       int    `json:"account_id"`
}
