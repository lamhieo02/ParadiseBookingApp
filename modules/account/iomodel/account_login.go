package iomodel

type AccountLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     int    `json:"type"`
	FullName string `json:"full_name"`
	Avatar   string `json:"avatar"`
}
