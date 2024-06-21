package postguideusecase

import (
	"context"
	"fmt"
	"paradise-booking/entities"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
	googlemapprovider "paradise-booking/provider/googlemap"
	"strings"
)

func (uc *postGuideUsecase) CreatePostGuide(ctx context.Context, data *postguideiomodel.CreatePostGuideReq) error {

	lat := data.Lat
	lng := data.Lng

	address := &googlemapprovider.GoogleMapAddress{}
	var err error
	if lat != 0 && lng != 0 {
		// make lat and lng round to 2 decimal
		// lat = math.Round(lat*100) / 100
		// lng = math.Round(lng*100) / 100

		address, err = uc.googleMap.GetAddressFromLatLng(ctx, lat, lng)
		if err != nil {
			fmt.Println("Error get address from lat lng", err)
			address = &googlemapprovider.GoogleMapAddress{}
		}
	}

	entity := &entities.PostGuide{
		PostOwnerId: data.PostOwnerID,
		TopicID:     data.TopicID,
		Title:       data.Title,
		Description: data.Description,
		Cover:       strings.Join(data.Images, ","),
		Lat:         data.Lat,
		Lng:         data.Lng,
		Country:     address.Country,
		State:       address.State,
		District:    address.District,
		Address:     data.Address,
		Languages:   strings.Join(data.Languages, ","),
		Schedule:    data.Schedule,
	}

	if err := uc.postGuideSto.Create(ctx, entity); err != nil {
		return err
	}

	return nil
}
