package placewishlistusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *placeWishListUsecase) GetPlaceByWishListID(ctx context.Context, wishListID int) ([]entities.Place, error) {

	// get list placeIDS by wishListID
	placeIDs, err := uc.placeWishListSto.GetPlaceIDs(ctx, wishListID)
	if err != nil {
		return nil, err
	}

	places, err := uc.placeSto.ListPlaceInIDs(ctx, placeIDs)
	if err != nil {
		return nil, err
	}
	return places, nil

}
