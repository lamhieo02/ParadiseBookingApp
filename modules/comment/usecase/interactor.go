package commentusecase

import "context"

type CommentStorage interface {
	DeleteByID(ctx context.Context, id int) error
}

type commentUseCase struct {
	commentStore CommentStorage
}

func NewCommentUseCase(commentSto CommentStorage) *commentUseCase {
	return &commentUseCase{commentStore: commentSto}
}
