package placeusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/modules/place/convert"
	"paradise-booking/modules/place/iomodel"

	"gorm.io/gorm"
)

func (uc *placeUseCase) GetPlaceByID(ctx context.Context, placeID int, userEmail string) (result *iomodel.GetPlaceResp, err error) {
	place, err := uc.placeStorage.GetPlaceByID(ctx, placeID)
	if err != nil {
		return nil, err
	}

	if place == nil {
		return nil, common.ErrEntityNotFound("place", err)
	}

	isFree := true

	userID := 0
	if userEmail != "" {
		user, err := uc.accountSto.GetAccountByEmail(ctx, userEmail)
		if err != nil {
			return nil, err
		}
		userID = user.Id
		// user, err := uc.accountSto.GetAccountByEmail(ctx, userEmail)
		// if err != nil {
		// 	return nil, err
		// }

		placeWishList, err := uc.placeWishSto.GetByCondition(ctx, map[string]interface{}{"user_id": userID, "place_id": place.Id})
		if err != nil {
			return nil, err
		}

		if len(placeWishList) > 0 {
			isFree = false
		}
	}

	// get rating average
	ratingAverage, err := uc.placeStorage.GetRatingAverageByPlaceId(ctx, int64(placeID))
	if err != nil {
		return nil, err
	}

	if ratingAverage == nil {
		defaulRating := 0.0
		ratingAverage = &defaulRating
	}

	result = convert.ConvertPlaceEntityToGetModel(place, isFree, ratingAverage)

	// get post guide related to place
	condition := map[string]interface{}{
		"state": place.State,
	}
	postGuideIds, err := uc.postGuideSto.ListPostGuideIdsByCondition(ctx, 10, condition)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	for _, postGuideID := range postGuideIds {
		postGuide, err := uc.postGuideCache.GetByID(ctx, postGuideID)
		if err != nil {
			return nil, err
		}

		result.PostGuideRelates = append(result.PostGuideRelates, iomodel.PostGuideRelate{
			ID:          postGuide.Id,
			TopicID:     postGuide.TopicID,
			TopicName:   constant.MapPostGuideTopic[postGuide.TopicID],
			Title:       postGuide.Title,
			Description: postGuide.Description,
			Cover:       postGuide.Cover,
			Country:     postGuide.Country,
			State:       postGuide.State,
			District:    postGuide.District,
			Address:     postGuide.Address,
		})
	}

	return result, nil
}
