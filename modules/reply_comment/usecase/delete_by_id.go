package replycommentusecase

import (
	"context"
)

func (uc *replyCommentUsecase) DeleteByID(ctx context.Context, id int) error {
	if err := uc.replyCommentSto.DeleteByID(ctx, id); err != nil {
		return err
	}

	return nil
}
