package placeusecase

import (
	"context"
	"fmt"
	"log"
	"paradise-booking/common"
	"paradise-booking/modules/place/convert"
	"paradise-booking/modules/place/iomodel"
	googlemapprovider "paradise-booking/provider/googlemap"
)

func (uc *placeUseCase) ListAllPlace(ctx context.Context, paging *common.Paging, filter *iomodel.Filter, userEmail string) (result []iomodel.GetPlaceResp, err error) {

	address := googlemapprovider.GoogleMapAddress{}

	// get geocode to fill country, state, district
	if filter.Lat != nil && filter.Lng != nil {
		lat := *filter.Lat
		lng := *filter.Lng
		address1, err := uc.googleMap.GetAddressFromLatLng(ctx, lat, lng)
		if err != nil {
			log.Printf("Error when get address from lat lng: %v", err)
			addr, err := uc.placeStorage.GetPlaceByCondition(ctx, map[string]interface{}{"lat": lat, "lng": lng})
			if err != nil {
				return nil, err
			}

			if len(addr) == 0 {
				return nil, fmt.Errorf("Cannot get address from lat %v lng %v", lat, lng)
			}

			if len(addr) > 0 {
				address.Country = addr[0].Country
				address.State = addr[0].State
				address.District = addr[0].District
			}
		} else {
			address = *address1
		}
	}

	paging.Process()
	places, err := uc.placeStorage.ListPlaces(ctx, paging, filter, &address)
	if err != nil {
		return nil, common.ErrCannotListEntity("place", err)
	}

	userID := 0
	if userEmail != "" {
		user, err := uc.accountSto.GetAccountByEmail(ctx, userEmail)
		if err != nil {
			return nil, err
		}
		userID = user.Id
	}

	// convert data to iomodel
	for _, place := range places {
		isFree := true

		if userID != 0 {
			placeWishList, err := uc.placeWishSto.GetByCondition(ctx, map[string]interface{}{"user_id": userID, "place_id": place.Id})
			if err != nil {
				return nil, err
			}

			if len(placeWishList) > 0 {
				isFree = false
			}
		}

		result = append(result, *convert.ConvertPlaceEntityToGetModel(&place, isFree))
	}
	return result, nil
}
