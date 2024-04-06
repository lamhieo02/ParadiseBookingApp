package postreviewusecase

import (
	"context"
	"paradise-booking/entities"
)

func (postReviewUsecase *postReviewUsecase) GetPostReviewByID(ctx context.Context, postReviewID int) (*entities.PostReview, error) {
	result, err := postReviewUsecase.postReviewStore.GetByID(ctx, postReviewID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
