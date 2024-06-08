package replycommentiomodel

import "paradise-booking/entities"

type UpdateReplyCommentReq struct {
	Content string `json:"content"`
	Image   string `json:"image"`
}

func (c *UpdateReplyCommentReq) ToEntity() *entities.ReplyComment {
	return &entities.ReplyComment{
		Content: c.Content,
		Image:   c.Image,
	}
}
