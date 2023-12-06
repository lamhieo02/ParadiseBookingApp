package wishlistusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

type WishListSto interface {
	Create(ctx context.Context, data *entities.WishList) error
	GetByID(ctx context.Context, id int) (*entities.WishList, error)
	GetByUserID(ctx context.Context, userId int, paging *common.Paging) ([]entities.WishList, error)
}

type wishListUsecase struct {
	wishListSto WishListSto
}

func NewWishListUseCase(wishListSto WishListSto) *wishListUsecase {
	return &wishListUsecase{wishListSto: wishListSto}
}
