package placeusecase

import (
	"context"
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
			return nil, err
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
