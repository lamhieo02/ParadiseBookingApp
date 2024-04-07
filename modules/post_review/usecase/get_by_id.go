package postreviewusecase

import (
	"context"
	"paradise-booking/modules/post_review/convert"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
)

func (postReviewUsecase *postReviewUsecase) GetPostReviewByID(ctx context.Context, postReviewID int) (*postreviewiomodel.PostReviewResp, error) {
	postReview, err := postReviewUsecase.postReviewStore.GetByID(ctx, postReviewID)
	if err != nil {
		return nil, err
	}

	comments, err := postReviewUsecase.commentStore.GetByPostReviewID(ctx, postReviewID)
	if err != nil {
		return nil, err
	}

	result := convert.ConvertPostReviewEntityToModelDetail(postReview, comments)

	return result, nil
}
