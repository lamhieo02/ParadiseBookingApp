package postreviewusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/modules/post_review/convert"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
)

func (postReviewUsecase *postReviewUsecase) ListPostReviewByAccountID(ctx context.Context, accountID int, paging *common.Paging) (*postreviewiomodel.ListPostReviewResp, error) {

	data, err := postReviewUsecase.postReviewStore.ListPostReviewByAccountID(ctx, accountID, paging)
	if err != nil {
		return nil, err
	}

	result := convert.ConvertListPostReviewToModel(data, paging)

	for i, v := range result.Data {
		likeCount, err := postReviewUsecase.likePostReviewSto.CountLikeByPostReview(ctx, int(v.ID))
		if err != nil {
			return nil, err
		}

		commentCount, err := postReviewUsecase.commentStore.CountCommentByPostReview(ctx, int(v.ID))
		if err != nil {
			return nil, err
		}

		result.Data[i].LikeCount = *likeCount
		result.Data[i].CommentCount = *commentCount

		likePostView, err := postReviewUsecase.likePostReviewSto.FindDataByCondition(ctx, map[string]interface{}{
			"account_id":     accountID,
			"post_review_id": v.ID,
		})
		if err != nil {
			return nil, err
		}

		if len(likePostView) == 0 || likePostView[0].Status == constant.UNLIKE_POST_REVIEW {
			result.Data[i].IsLiked = false
		} else {
			result.Data[i].IsLiked = true
		}

	}

	return &result, nil
}
