package wishlisthandler

import (
	"context"
	"paradise-booking/entities"
	wishlistiomodel "paradise-booking/modules/wishlist/iomodel"
)

type wishListUseCase interface {
	CreateWishList(ctx context.Context, data *wishlistiomodel.CreateWishListReq) (*entities.WishList, error)
	GetWishListByID(ctx context.Context, id int) (*entities.WishList, error)
	GetWishListByUserID(ctx context.Context, userId int) ([]entities.WishList, error)
}

type wishListHandler struct {
	wishListUC wishListUseCase
}

func NewWishListHandler(wishListUC wishListUseCase) *wishListHandler {
	return &wishListHandler{wishListUC: wishListUC}
}
