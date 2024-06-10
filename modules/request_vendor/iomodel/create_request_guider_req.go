package requestvendoriomodel

import (
	"paradise-booking/entities"
)

type CreateRequestVendorReq struct {
	UserID      int    `json:"user_id" form:"user_id"`
	FullName    string `json:"full_name" form:"full_name"`
	Username    string `json:"username" form:"username"`
	Email       string `json:"email" form:"email"`
	Phone       string `json:"phone" form:"phone"`
	DOB         string `json:"dob" form:"dob"`
	Address     string `json:"address" form:"address"`
	Description string `json:"description" form:"description"`
	Experience  string `json:"experience" form:"experience"`
}

func (req *CreateRequestVendorReq) ToEntity() *entities.RequestVendor {
	return &entities.RequestVendor{
		UserId:      req.UserID,
		FullName:    req.FullName,
		Username:    req.Username,
		Email:       req.Email,
		Phone:       req.Phone,
		DOB:         req.DOB,
		Address:     req.Address,
		Description: req.Description,
		Experience:  req.Experience,
	}
}
