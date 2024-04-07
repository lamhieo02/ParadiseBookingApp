package replycommenthandler

import (
	"context"
	replycommentiomodel "paradise-booking/modules/reply_comment/iomodel"
)

type ReplyCommentUseCase interface {
	ReplySourceComment(ctx context.Context, data *replycommentiomodel.ReplyCommentReq) error
}

type replyCommentHandler struct {
	replyCommentUC ReplyCommentUseCase
}

func NewReplyCommentHandler(replyCommentUC ReplyCommentUseCase) *replyCommentHandler {
	return &replyCommentHandler{replyCommentUC: replyCommentUC}
}
