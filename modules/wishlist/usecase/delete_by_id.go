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

	// get all place in this wishlist and delete in cache
	placeWishList, err := uc.placeWishListSto.GetByCondition(ctx, map[string]interface{}{"wishlist_id": id})
	if err != nil {
		return common.ErrCannotGetEntity(entities.PlaceWishList{}.TableName(), err)
	}

	// delete in cache
	for _, v := range placeWishList {
		key := v.CacheKey()
		uc.cacheStore.Delete(ctx, key)
	}

	return nil
}
