package replycommentusecase

import (
	"context"
	"paradise-booking/entities"
	replycommentiomodel "paradise-booking/modules/reply_comment/iomodel"
)

func (uc *replyCommentUsecase) ReplySourceComment(ctx context.Context, data *replycommentiomodel.ReplyCommentReq) error {
	// check source_comment is exist
	_, err := uc.commentSto.GetByID(ctx, data.SourceCommentID)
	if err != nil {
		return err
	}

	replyCommentEntity := entities.ReplyComment{
		SourcecommentID: data.SourceCommentID,
		Content:         data.Content,
		Image:           data.Image,
		Videos:          data.Videos,
		AccountID:       data.AccountID,
	}

	// create reply comment
	if err := uc.replyCommentSto.Create(ctx, &replyCommentEntity); err != nil {
		return err
	}

	return nil
}
