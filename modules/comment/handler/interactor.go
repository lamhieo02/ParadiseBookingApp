package commenthandler

import (
	"context"
	commentiomodel "paradise-booking/modules/comment/iomodel"
)

type CommentUseCase interface {
	DeleteCommentByID(ctx context.Context, id int) error
	GetCommentByPostReviewID(ctx context.Context, id int) ([]*commentiomodel.CommentResp, error)
}

type commentHandler struct {
	commentUC CommentUseCase
}

func NewCommentHandler(commentUseCase CommentUseCase) *commentHandler {
	return &commentHandler{commentUC: commentUseCase}
}
