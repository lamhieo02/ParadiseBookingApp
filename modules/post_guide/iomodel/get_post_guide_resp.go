package postguideiomodel

import "time"

type GetPostGuideResp struct {
	ID            int           `json:"id" gorm:"column:id"`
	PostOwnerId   int           `json:"post_owner_id" gorm:"column:post_owner_id"`
	PostOwner     OwnerResp     `json:"post_owner"`
	TopicID       int           `json:"topic_id" gorm:"column:topic_id"`
	TopicName     string        `json:"topic_name"`
	Title         string        `json:"title" gorm:"column:title"`
	Description   string        `json:"description" gorm:"column:description"`
	Cover         string        `json:"cover" gorm:"column:cover"`
	Lat           float64       `json:"lat" gorm:"column:lat"`
	Lng           float64       `json:"lng" gorm:"column:lng"`
	Location      Location      `json:"location"`
	Address       string        `json:"address" gorm:"column:address"`
	RatingAverage float64       `json:"rating_average" form:"rating_average"`
	Languages     []string      `json:"languages" gorm:"column:languages"`
	CreatedAt     time.Time     `json:"created_at" gorm:"column:created_at"`
	Schedule      string        `json:"schedule" gorm:"column:schedule"`
	PlaceRelates  []PlaceRelate `json:"place_relates"`
}

type PlaceRelate struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	PricePerNight float64 `json:"price_per_night"`
	Address       string  `json:"address"`
	Cover         string  `json:"cover"`
	Country       string  `json:"country"`
	State         string  `json:"state"`
	District      string  `json:"district"`
}

type Location struct {
	Country  string `json:"country"`
	State    string `json:"state"`
	District string `json:"district"`
}

type OwnerResp struct {
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
