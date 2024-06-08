package commentiomodel

import "paradise-booking/entities"

type UpdateCommentReq struct {
	Content string `json:"content"`
	Image   string `json:"image"`
}

func (c *UpdateCommentReq) ToEntity() *entities.Comment {
	return &entities.Comment{
		Content: c.Content,
		Image:   c.Image,
	}
}
