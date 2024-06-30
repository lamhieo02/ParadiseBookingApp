package placewishlistusecase

import (
	"context"
	"paradise-booking/common"
	placewishlistiomodel "paradise-booking/modules/place_wishlist/iomodel"
	"strings"
)

func (uc *placeWishListUsecase) GetPlaceByWishListID(ctx context.Context, wishListID int, paging *common.Paging, userID int) ([]placewishlistiomodel.DataPlace, error) {
	paging.Process()

	var result []placewishlistiomodel.DataPlace
	// get list placeIDS by wishListID
	placeIDs, err := uc.placeWishListSto.GetPlaceIDs(ctx, wishListID, paging, userID)
	if err != nil {
		return nil, err
	}

	places, err := uc.placeSto.ListPlaceInIDs(ctx, placeIDs)
	if err != nil {
		return nil, err
	}

	for _, place := range places {
		result = append(result, placewishlistiomodel.DataPlace{
			ID:               place.Id,
			VendorID:         place.VendorID,
			Name:             place.Name,
			Description:      place.Description,
			PricePerNight:    place.PricePerNight,
			Address:          place.Address,
			Images:           strings.Split(place.Cover, ","),
			Lat:              place.Lat,
			Lng:              place.Lng,
			Country:          place.Country,
			State:            place.State,
			District:         place.District,
			MaxGuest:         place.MaxGuest,
			NumBed:           place.NumBed,
			BedRoom:          place.BedRoom,
			NumPlaceOriginal: place.NumPlaceOriginal,
		})
	}

	return result, nil

}
