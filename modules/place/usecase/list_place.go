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

func (uc *placeUseCase) ListAllPlace(ctx context.Context, paging *common.Paging, filter *iomodel.Filter) (result []iomodel.GetPlaceResp, err error) {

	address := &googlemapprovider.GoogleMapAddress{}

	// get geocode to fill country, state, district
	if filter.Lat != nil && filter.Lng != nil {
		lat := *filter.Lat
		lng := *filter.Lng
		address, err = uc.googleMap.GetAddressFromLatLng(ctx, lat, lng)
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
		}
	}

	paging.Process()
	places, err := uc.placeStorage.ListPlaces(ctx, paging, filter, address)
	if err != nil {
		return nil, common.ErrCannotListEntity("place", err)
	}

	// convert data to iomodel
	for _, place := range places {
		result = append(result, *convert.ConvertPlaceEntityToGetModel(&place))
	}
	return result, nil
}
