package replycommenthandler

import (
	"context"
	"paradise-booking/entities"
	replycommentiomodel "paradise-booking/modules/reply_comment/iomodel"
)

type ReplyCommentUseCase interface {
	ReplySourceComment(ctx context.Context, data *replycommentiomodel.ReplyCommentReq) error
	DeleteByID(ctx context.Context, id int) error
	EditReplyCommentByID(ctx context.Context, id int, data *entities.ReplyComment) error
}

type replyCommentHandler struct {
	replyCommentUC ReplyCommentUseCase
}

func NewReplyCommentHandler(replyCommentUC ReplyCommentUseCase) *replyCommentHandler {
	return &replyCommentHandler{replyCommentUC: replyCommentUC}
}
