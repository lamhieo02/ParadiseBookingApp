package wishlistusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (uc *wishListUsecase) GetWishListByUserID(ctx context.Context, userId int) ([]entities.WishList, error) {
	res, err := uc.wishListSto.GetByUserID(ctx, userId)
	if err != nil {
		return nil, common.ErrCannotGetEntity(entities.WishList{}.TableName(), err)
	}

	return res, nil
}
