package postreviewusecase

import (
	"context"
	"paradise-booking/entities"
)

type PostReviewStore interface {
	Create(ctx context.Context, data *entities.PostReview) error
}

type postReviewUsecase struct {
	postReviewStore PostReviewStore
}

func NewPostReviewUseCase(PostReviewStore PostReviewStore) *postReviewUsecase {
	return &postReviewUsecase{postReviewStore: PostReviewStore}
}
