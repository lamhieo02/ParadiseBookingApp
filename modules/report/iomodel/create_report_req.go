package reportiomodel

import (
	"paradise-booking/entities"
	"strings"
)

type CreateReportReq struct {
	ObjectID    int      `json:"object_id"`
	ObjectType  int      `json:"object_type"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	StatusID    int      `json:"status_id"`
	Videos      []string `json:"videos"`
	Images      []string `json:"images"`
	UserID      int      `json:"user_id"`
}

func (c *CreateReportReq) ToEntity() *entities.Report {
	return &entities.Report{
		ObjectID:    c.ObjectID,
		ObjectType:  c.ObjectType,
		UserID:      c.UserID,
		Type:        c.Type,
		Description: c.Description,
		StatusID:    c.StatusID,
		Videos:      strings.Join(c.Videos, ","),
		Images:      strings.Join(c.Images, ","),
	}
}
