package placewishlistusecase

import (
	"context"
	"paradise-booking/entities"
	"paradise-booking/modules/place_wishlist/iomodel"
)

func (placeWishListUsecase *placeWishListUsecase) CreatePlaceWishList(ctx context.Context, data *iomodel.CreatePlaceWishListReq) (*entities.PlaceWishList, error) {
	model := entities.PlaceWishList{
		PlaceId:    data.PlaceID,
		WishListId: data.WishListID,
	}
	if err := placeWishListUsecase.placeWishListSto.Create(ctx, &model); err != nil {
		return nil, err
	}

	return &model, nil
}
