package commentusecase

import (
	"context"
	"paradise-booking/entities"
)

type CommentStorage interface {
	DeleteByID(ctx context.Context, id int) error
	GetByPostReviewID(ctx context.Context, postReviewID int) ([]*entities.Comment, error)
	UpdateByID(ctx context.Context, id int, data *entities.Comment) error
}

type ReplyCommentStorage interface {
	GetBySourceCommentID(ctx context.Context, sourceCommentID int) ([]entities.ReplyComment, error)
}

type AccountStorage interface {
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type commentUseCase struct {
	commentStore CommentStorage
	replyComment ReplyCommentStorage
	accountStore AccountStorage
}

func NewCommentUseCase(commentSto CommentStorage, replyCommentSto ReplyCommentStorage, accountSto AccountStorage) *commentUseCase {
	return &commentUseCase{commentStore: commentSto, replyComment: replyCommentSto, accountStore: accountSto}
}
