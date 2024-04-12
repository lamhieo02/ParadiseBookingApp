package postreviewiomodel

import "time"

type Filter struct {
	TopicID  int     `json:"topic_id"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Country  string  `json:"country"`
	State    string  `json:"state"`
	District string  `json:"district"`
	// IsLatest  bool       `json:"is_latest"`
	// IsPopular bool       `json:"is_popular"`
	DateFrom *time.Time `json:"date_from"`
	DateTo   *time.Time `json:"date_to"`
}
