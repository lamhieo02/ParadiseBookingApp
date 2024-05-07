package postguideusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
)

func (uc *postGuideUsecase) ListPostGuideByFilter(ctx context.Context, paging *common.Paging, filter *postguideiomodel.Filter) (*postguideiomodel.ListPostGuideResp, error) {
	paging.Process()

	data, err := uc.postGuideSto.ListByFilter(ctx, paging, filter)
	if err != nil {
		return nil, err
	}

	result := &postguideiomodel.ListPostGuideResp{}
	result.Data = make([]postguideiomodel.GetPostGuideResp, len(data))
	for i, v := range data {
		result.Data[i] = postguideiomodel.GetPostGuideResp{
			ID:          v.Id,
			PostOwnerId: v.PostOwnerId,
			TopicID:     v.TopicID,
			TopicName:   constant.MapPostGuideTopic[v.TopicID],
			Title:       v.Title,
			Description: v.Description,
			Cover:       v.Cover,
			Lat:         v.Lat,
			Lng:         v.Lng,
			Address:     v.Address,
			Location: postguideiomodel.Location{
				Country:  v.Country,
				State:    v.State,
				District: v.District,
			},
		}

		owner, err := uc.accountCache.GetProfileByID(ctx, v.PostOwnerId)
		if err != nil {
			return nil, err
		}

		result.Data[i].PostOwner = postguideiomodel.OwnerResp{
			UserName: owner.Username,
			Avatar:   owner.Avatar,
			FullName: owner.FullName,
			Email:    owner.Email,
		}
	}

	result.Paging = paging

	return result, nil

}
