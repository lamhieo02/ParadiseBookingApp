package commenthandler

import (
	"context"
)

type CommentUseCase interface {
	DeleteCommentByID(ctx context.Context, id int) error
}

type commentHandler struct {
	commentUC CommentUseCase
}

func NewCommentHandler(commentUseCase CommentUseCase) *commentHandler {
	return &commentHandler{commentUC: commentUseCase}
}
