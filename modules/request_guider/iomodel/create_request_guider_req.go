package requestguideriomodel

import (
	"paradise-booking/entities"
	"strings"
)

type CreateRequestGuiderReq struct {
	UserID        int      `json:"user_id" form:"user_id"`
	FullName      string   `json:"full_name" form:"full_name"`
	Username      string   `json:"username" form:"username"`
	Email         string   `json:"email" form:"email"`
	Phone         string   `json:"phone" form:"phone"`
	DOB           string   `json:"dob" form:"dob"`
	Address       string   `json:"address" form:"address"`
	Description   string   `json:"description" form:"description"`
	Experience    string   `json:"experience" form:"experience"`
	Reason        string   `json:"reason" form:"reason"`
	GoalsOfTravel []string `json:"goals_of_travel" form:"goals_of_travel"`
	Languages     []string `json:"languages" form:"languages"`
}

func (req *CreateRequestGuiderReq) ToEntity() *entities.RequestGuider {
	return &entities.RequestGuider{
		UserId:       req.UserID,
		FullName:     req.FullName,
		Username:     req.Username,
		Email:        req.Email,
		Phone:        req.Phone,
		DOB:          req.DOB,
		Address:      req.Address,
		Description:  req.Description,
		Experience:   req.Experience,
		Reason:       req.Reason,
		GoalOfTravel: strings.Join(req.GoalsOfTravel, ","),
		Languages:    strings.Join(req.Languages, ","),
	}
}
