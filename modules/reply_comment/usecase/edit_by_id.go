package replycommentusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *replyCommentUsecase) EditReplyCommentByID(ctx context.Context, id int, data *entities.ReplyComment) error {
	if err := uc.replyCommentSto.UpdateByID(ctx, id, data); err != nil {
		return err
	}

	return nil
}
