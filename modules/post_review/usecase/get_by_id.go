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

	likeCount, err := postReviewUsecase.likePostReviewSto.CountLikeByPostReview(ctx, postReviewID)
	if err != nil {
		return nil, err
	}

	commentCount, err := postReviewUsecase.commentStore.CountCommentByPostReview(ctx, postReviewID)
	if err != nil {
		return nil, err
	}

	result.LikeCount = *likeCount
	result.CommentCount = *commentCount

	// get reply comment
	for i, v := range result.Comments {
		replyComments, err := postReviewUsecase.replyComment.GetBySourceCommentID(ctx, int(v.ID))
		if err != nil {
			return nil, err
		}

		tmpReplyComments := []postreviewiomodel.ReplyCommentResp{}
		for _, replyComment := range replyComments {
			tmpReplyComments = append(tmpReplyComments, postreviewiomodel.ReplyCommentResp{
				ID:        int64(replyComment.Id),
				Content:   replyComment.Content,
				Image:     replyComment.Image,
				Videos:    replyComment.Videos,
				AccountID: int64(replyComment.AccountID),
			})
		}

		result.Comments[i].ReplyComments = tmpReplyComments
	}

	return result, nil
}
