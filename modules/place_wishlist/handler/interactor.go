package placewishlisthandler

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	placewishlistiomodel "paradise-booking/modules/place_wishlist/iomodel"
)

type PlaceWishListUseCase interface {
	CreatePlaceWishList(ctx context.Context, data *placewishlistiomodel.CreatePlaceWishListReq, userID int) (*entities.PlaceWishList, error)
	DeletePlaceWishList(ctx context.Context, placeId, wishListID int) error
	GetPlaceByWishListID(ctx context.Context, wishListID int, paging *common.Paging, userID int) ([]placewishlistiomodel.DataPlace, error)
}

type placeWishListHandler struct {
	placeWishListUC PlaceWishListUseCase
}

func NewPlaceWishListHandler(PlaceWishListUC PlaceWishListUseCase) *placeWishListHandler {
	return &placeWishListHandler{placeWishListUC: PlaceWishListUC}
}
