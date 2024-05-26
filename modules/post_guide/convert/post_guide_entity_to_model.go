package postguideconvert

import (
	"paradise-booking/constant"
	"paradise-booking/entities"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
	"strings"
)

func ConvertPostGuideEntityToModel(postGuideEntity *entities.PostGuide, owner *entities.Account) postguideiomodel.GetPostGuideResp {
	result := postguideiomodel.GetPostGuideResp{}
	result.ID = postGuideEntity.Id
	result.PostOwnerId = postGuideEntity.PostOwnerId
	result.Title = postGuideEntity.Title
	result.Description = postGuideEntity.Description
	result.Cover = postGuideEntity.Cover
	result.TopicID = postGuideEntity.TopicID
	result.TopicName = constant.MapPostGuideTopic[postGuideEntity.TopicID]
	result.Lat = postGuideEntity.Lat
	result.Lng = postGuideEntity.Lng
	result.Address = postGuideEntity.Address
	result.PostOwner = postguideiomodel.OwnerResp{
		UserName: owner.Username,
		Avatar:   owner.Avatar,
		FullName: owner.FullName,
		Email:    owner.Email,
	}
	result.Location = postguideiomodel.Location{
		Country:  postGuideEntity.Country,
		State:    postGuideEntity.State,
		District: postGuideEntity.District,
	}
	result.CreatedAt = *postGuideEntity.CreatedAt
	result.Languages = strings.Split(postGuideEntity.Languages, ",")
	result.Schedule = postGuideEntity.Schedule
	return result
}
