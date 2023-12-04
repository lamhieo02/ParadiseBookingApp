package placewishlistusecase

import (
	"context"
	"paradise-booking/common"
)

func (uc *placeWishListUsecase) DeletePlaceWishList(ctx context.Context, placeId, wishListID int) error {

	if err := uc.placeWishListSto.Delete(ctx, placeId, wishListID); err != nil {
		return common.ErrCannotDeleteEntity("place_wish_list", err)
	}
	return nil
}
