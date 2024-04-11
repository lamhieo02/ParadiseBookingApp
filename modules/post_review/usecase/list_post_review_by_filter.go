package postreviewusecase

import (
	"context"
	"fmt"
	"paradise-booking/common"
	"paradise-booking/modules/post_review/convert"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
	googlemapprovider "paradise-booking/provider/googlemap"
)

func (postReviewUsecase *postReviewUsecase) ListPostReviewByFilter(ctx context.Context, paging *common.Paging, filter *postreviewiomodel.Filter) (*postreviewiomodel.ListPostReviewResp, error) {

	paging.Process()
	data, err := postReviewUsecase.postReviewStore.ListPostReviewByFilter(ctx, paging, filter)
	if err != nil {
		return nil, err
	}

	result := convert.ConvertListPostReviewToModel(data, paging)

	for i, v := range result.Data {
		likeCount, err := postReviewUsecase.likePostReviewSto.CountLikeByPostReview(ctx, int(v.ID))
		if err != nil {
			return nil, err
		}

		commentCount, err := postReviewUsecase.commentStore.CountCommentByPostReview(ctx, int(v.ID))
		if err != nil {
			return nil, err
		}

		result.Data[i].LikeCount = *likeCount
		result.Data[i].CommentCount = *commentCount

		// get location
		location, err := postReviewUsecase.googleMap.GetAddressFromLatLng(ctx, result.Data[i].Lat, result.Data[i].Lng)
		if err != nil {
			fmt.Printf("error get location from lat lng: %v", err)
			location = &googlemapprovider.GoogleMapAddress{}
		}
		result.Data[i].Country = location.Country
		result.Data[i].State = location.State
		result.Data[i].District = location.District
	}

	return &result, nil
}
