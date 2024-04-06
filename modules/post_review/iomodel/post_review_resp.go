package postreviewiomodel

import "time"

type PostReviewResp struct {
	ID          int64      `json:"id"`
	PostOwnerID int64      `json:"post_owner_id"`
	Title       string     `json:"title"`
	Topic       string     `json:"topic"`
	Content     string     `json:"content"`
	Image       string     `json:"image"`
	Lat         float64    `json:"lat"`
	Lng         float64    `json:"lng"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
