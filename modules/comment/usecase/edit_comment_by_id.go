package commentusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *commentUseCase) EditCommentByID(ctx context.Context, id int, data *entities.Comment) error {
	if err := uc.commentStore.UpdateByID(ctx, id, data); err != nil {
		return err
	}

	return nil
}
