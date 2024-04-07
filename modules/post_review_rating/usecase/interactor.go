package postreviewratingusecase

import (
	"context"
	"paradise-booking/entities"
)

type PostReviewRatingStore interface {
	Create(ctx context.Context, data *entities.PostReviewRating) error
	UpdateByID(ctx context.Context, id int, data *entities.PostReviewRating) error
	DeleteByID(ctx context.Context, id int) error
	UpdateWithMap(ctx context.Context, data *entities.PostReviewRating, props map[string]interface{}) error
	GetByID(ctx context.Context, id int) (*entities.PostReviewRating, error)
}

type CommentStorage interface {
	Create(ctx context.Context, data *entities.Comment) error
	UpdateByID(ctx context.Context, id int, data *entities.Comment) error
	DeleteByID(ctx context.Context, id int) error
}

type AccountStorage interface {
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type postReviewRatingUsecase struct {
	postReviewRatingStore PostReviewRatingStore
	commentStore          CommentStorage
	accountSto            AccountStorage
}

func NewPostReviewRatingUseCase(PostReviewRatingStore PostReviewRatingStore, commentStorage CommentStorage, accountSto AccountStorage) *postReviewRatingUsecase {
	return &postReviewRatingUsecase{postReviewRatingStore: PostReviewRatingStore, commentStore: commentStorage, accountSto: accountSto}
}
