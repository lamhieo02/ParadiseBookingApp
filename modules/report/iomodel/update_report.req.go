package reportiomodel

import (
	"paradise-booking/entities"
	"strings"
)

type UpdateReportReq struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	StatusID    int      `json:"status_id"`
	Videos      []string `json:"videos"`
	Images      []string `json:"images"`
}

func (c *UpdateReportReq) ToEntity() *entities.Report {
	return &entities.Report{
		Type:        c.Type,
		Description: c.Description,
		StatusID:    c.StatusID,
		Videos:      strings.Join(c.Videos, ","),
		Images:      strings.Join(c.Images, ","),
	}
}
