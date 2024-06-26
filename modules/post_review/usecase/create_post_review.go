package postreviewusecase

import (
	"context"
	"log"
	"paradise-booking/entities"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
	googlemapprovider "paradise-booking/provider/googlemap"
	"strings"
)

func (postReviewUsecase *postReviewUsecase) CreatePostReview(ctx context.Context, data *postreviewiomodel.CreatePostReviewReq) error {

	// get location from lat lng
	ggAddress, err := postReviewUsecase.googleMap.GetAddressFromLatLng(ctx, data.Lat, data.Lng)
	if err != nil {
		log.Printf("Error when get address from lat lng: %v", err)
		ggAddress = &googlemapprovider.GoogleMapAddress{}
		ggAddress.Country = ""
		ggAddress.State = ""
		ggAddress.District = ""
	}

	models := entities.PostReview{
		PostOwnerId: int(data.AccountID),
		Title:       data.Title,
		Topic:       data.Topic,
		Content:     data.Content,
		Lat:         data.Lat,
		Lng:         data.Lng,
		Image:       strings.Join(data.Images, ","),
		Videos:      strings.Join(data.Videos, ","),
		Country:     ggAddress.Country,
		State:       ggAddress.State,
		District:    ggAddress.District,
	}

	if err := postReviewUsecase.postReviewStore.Create(ctx, &models); err != nil {
		return err
	}

	return nil
}
