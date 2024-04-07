package likepostreviewusecase

import (
	"context"
	"paradise-booking/entities"
)

type LikePostReviewStorage interface {
	UpdateWithMap(ctx context.Context, data *entities.LikePostReview, props map[string]interface{}) error
	Create(ctx context.Context, data *entities.LikePostReview) error
	FindDataByCondition(ctx context.Context, condition map[string]interface{}) ([]*entities.LikePostReview, error)
}

type likePostReviewUsecase struct {
	likePostReviewStore LikePostReviewStorage
}

func NewLikePostReviewUseCase(likePostReviewSto LikePostReviewStorage) *likePostReviewUsecase {
	return &likePostReviewUsecase{likePostReviewStore: likePostReviewSto}
}
