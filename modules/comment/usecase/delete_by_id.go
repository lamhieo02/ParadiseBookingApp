package commentusecase

import (
	"context"
)

func (uc *commentUseCase) DeleteCommentByID(ctx context.Context, id int) error {
	if err := uc.commentStore.DeleteByID(ctx, id); err != nil {
		return err
	}

	return nil
}
