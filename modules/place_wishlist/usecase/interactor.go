package placewishlistusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

type PlaceWishListSto interface {
	Create(ctx context.Context, data *entities.PlaceWishList) error
	Delete(ctx context.Context, place_id, wish_list_id int) error
	GetPlaceIDs(ctx context.Context, wish_list_id int, paging *common.Paging, userID int) ([]int, error)
}

type PlaceSto interface {
	ListPlaceInIDs(ctx context.Context, placeIds []int) ([]entities.Place, error)
}

type placeWishListUsecase struct {
	placeWishListSto PlaceWishListSto
	placeSto         PlaceSto
}

func NewPlaceWishListUseCase(PlaceWishListSto PlaceWishListSto, placeSto PlaceSto) *placeWishListUsecase {
	return &placeWishListUsecase{placeWishListSto: PlaceWishListSto, placeSto: placeSto}
}
