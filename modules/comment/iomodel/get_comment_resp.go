package commentiomodel

import "time"

type CommentResp struct {
	ID            int64               `json:"id"`
	Content       string              `json:"content"`
	Image         string              `json:"image"`
	Videos        string              `json:"videos"`
	AccountID     int64               `json:"account_id"`
	Owner         OwnerResp           `json:"owner"`
	DateComment   *time.Time          `json:"date_comment"`
	ReplyComments []*ReplyCommentResp `json:"reply_comments"`
}

type ReplyCommentResp struct {
	ID          int64      `json:"id"`
	Content     string     `json:"content"`
	Image       string     `json:"image"`
	Videos      string     `json:"videos"`
	AccountID   int64      `json:"account_id"`
	Owner       OwnerResp  `json:"owner"`
	DateComment *time.Time `json:"date_comment"`
}

type OwnerResp struct {
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
