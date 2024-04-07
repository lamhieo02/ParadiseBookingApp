package replycommentusecase

import (
	"context"
	"paradise-booking/entities"
)

type ReplyCommentStorage interface {
	Create(ctx context.Context, data *entities.ReplyComment) error
	DeleteByID(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*entities.ReplyComment, error)
	UpdateWithMap(ctx context.Context, data *entities.ReplyComment, props map[string]interface{}) error
	GetBySourceCommentID(ctx context.Context, sourceCommentID int) ([]entities.ReplyComment, error)
}

type CommentStorage interface {
	GetByID(ctx context.Context, id int) (*entities.Comment, error)
}

type replyCommentUsecase struct {
	replyCommentSto ReplyCommentStorage
	commentSto      CommentStorage
}

func NewReplyCommentUsecase(replySto ReplyCommentStorage, commentSto CommentStorage) *replyCommentUsecase {
	return &replyCommentUsecase{replyCommentSto: replySto, commentSto: commentSto}
}
