package placewishliststorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *placeWishListStorage) GetPlaceIDs(ctx context.Context, wish_list_id int) ([]int, error) {
	db := s.db
	var res []int

	if err := db.Model(entities.PlaceWishList{}).Where("wishlist_id = ?", wish_list_id).Pluck("place_id", &res).Error; err != nil {
		return nil, common.ErrorDB(err)
	}
	return res, nil
}
