package wishlistusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (uc *wishListUsecase) DeleteByID(ctx context.Context, id int) error {
	res, err := uc.wishListSto.GetByID(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(entities.WishList{}.TableName(), err)
	}

	// check exist by id
	if res == nil {
		return common.ErrEntityNotFound(entities.WishList{}.TableName(), nil)
	}

	// delete
	if err := uc.wishListSto.DeleteByID(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(entities.WishList{}.TableName(), err)
	}

	return nil
}
