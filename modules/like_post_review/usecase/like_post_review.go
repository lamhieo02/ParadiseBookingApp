package likepostreviewusecase

import (
	"context"
	"paradise-booking/constant"
	"paradise-booking/entities"
	likepostreviewiomodel "paradise-booking/modules/like_post_review/iomodel"
)

func (likePostReviewUsecase *likePostReviewUsecase) LikePostReview(ctx context.Context, req *likepostreviewiomodel.LikePostReviewReq) error {
	// check if already exist
	likePostReviews, err := likePostReviewUsecase.likePostReviewStore.FindDataByCondition(ctx, map[string]interface{}{
		"account_id":     req.AccountID,
		"post_review_id": req.PostReviewID,
	})
	if err != nil {
		return err
	}

	if len(likePostReviews) == 0 {
		// create new like
		if err := likePostReviewUsecase.likePostReviewStore.Create(ctx, &entities.LikePostReview{
			AccountId:    int(req.AccountID),
			PostReviewId: int(req.PostReviewID),
			Status:       constant.LIKE_POST_REVIEW,
		}); err != nil {
			return err
		}
		return nil
	}
	// because we only have 1 account_id and post_review_id
	likePostReview := likePostReviews[0]
	status := 0
	if req.Type == constant.LIKE_POST_REVIEW {
		status = constant.LIKE_POST_REVIEW
	} else if req.Type == constant.UNLIKE_POST_REVIEW {
		status = constant.UNLIKE_POST_REVIEW
	}

	if err := likePostReviewUsecase.likePostReviewStore.UpdateWithMap(ctx, likePostReview, map[string]interface{}{
		"status": status,
	}); err != nil {
		return err
	}

	return nil
}
