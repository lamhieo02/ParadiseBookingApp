package postreviewusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (postReviewUsecase *postReviewUsecase) ListPostReviewByAccountID(ctx context.Context, accountID int, paging *common.Paging) ([]*entities.PostReview, error) {

	result, err := postReviewUsecase.postReviewStore.ListPostReviewByAccountID(ctx, accountID, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}
