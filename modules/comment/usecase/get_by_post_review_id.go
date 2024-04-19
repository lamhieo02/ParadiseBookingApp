package commentusecase

import (
	"context"
	commentiomodel "paradise-booking/modules/comment/iomodel"
)

func (uc *commentUseCase) GetCommentByPostReviewID(ctx context.Context, id int) ([]*commentiomodel.CommentResp, error) {
	data, err := uc.commentStore.GetByPostReviewID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]*commentiomodel.CommentResp, len(data))

	// aggregate data for response
	for i, v := range data {
		replyComments, err := uc.replyComment.GetBySourceCommentID(ctx, v.Id)
		if err != nil {
			return nil, err
		}

		ownerComment, err := uc.accountStore.GetProfileByID(ctx, int(v.AccountID))
		if err != nil {
			return nil, err
		}

		tmpReplyComments := []*commentiomodel.ReplyCommentResp{}
		for _, replyComment := range replyComments {
			// get owner of reply comment
			owner, err := uc.accountStore.GetProfileByID(ctx, int(replyComment.AccountID))
			if err != nil {
				return nil, err
			}
			tmpReplyComments = append(tmpReplyComments, &commentiomodel.ReplyCommentResp{
				ID:          int64(replyComment.Id),
				Content:     replyComment.Content,
				Image:       replyComment.Image,
				Videos:      replyComment.Videos,
				AccountID:   int64(replyComment.AccountID),
				DateComment: replyComment.CreatedAt,
				Owner: commentiomodel.OwnerResp{
					UserName: owner.Username,
					Avatar:   owner.Avatar,
					FullName: owner.FullName,
					Email:    owner.Email,
				},
			})
		}

		result[i] = &commentiomodel.CommentResp{
			ID:          int64(v.Id),
			Content:     v.Content,
			Image:       v.Image,
			Videos:      v.Videos,
			AccountID:   int64(v.AccountID),
			DateComment: v.CreatedAt,
			Owner: commentiomodel.OwnerResp{
				UserName: ownerComment.Username,
				Avatar:   ownerComment.Avatar,
				FullName: ownerComment.FullName,
				Email:    ownerComment.Email,
			},
			ReplyComments: tmpReplyComments,
		}
	}

	return result, nil
}
