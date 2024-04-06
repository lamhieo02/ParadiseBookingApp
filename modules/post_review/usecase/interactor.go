package postreviewusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

type PostReviewStore interface {
	Create(ctx context.Context, data *entities.PostReview) error
	UpdateWithMap(ctx context.Context, data *entities.PostReview, props map[string]interface{}) error
	GetByID(ctx context.Context, postReviewID int) (*entities.PostReview, error)
	UpdateByID(ctx context.Context, id int, data *entities.PostReview) error
	ListPostReviewByAccountID(ctx context.Context, accountID int, paging *common.Paging) ([]*entities.PostReview, error)
	DeleteByID(ctx context.Context, postReviewID int) error
}

type postReviewUsecase struct {
	postReviewStore PostReviewStore
}

func NewPostReviewUseCase(PostReviewStore PostReviewStore) *postReviewUsecase {
	return &postReviewUsecase{postReviewStore: PostReviewStore}
}
