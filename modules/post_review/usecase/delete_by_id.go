package postreviewusecase

import (
	"context"
)

func (postReviewUsecase *postReviewUsecase) DeletePostReviewByID(ctx context.Context, postReviewID int) error {
	// old data to check exist
	_, err := postReviewUsecase.postReviewStore.GetByID(ctx, postReviewID)
	if err != nil {
		return err
	}

	// delete data
	if err := postReviewUsecase.postReviewStore.DeleteByID(ctx, postReviewID); err != nil {
		return err
	}

	return nil
}
