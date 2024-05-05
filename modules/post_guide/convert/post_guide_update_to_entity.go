package postguideconvert

import (
	"paradise-booking/entities"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
)

func ConvertPostGuideUpdateToEntity(model *postguideiomodel.UpdatePostGuideReq) *entities.PostGuide {
	return &entities.PostGuide{
		TopicID:     model.TopicID,
		Title:       model.Title,
		Description: model.Description,
		Cover:       model.Cover,
		Lat:         model.Lat,
		Lng:         model.Lng,
	}
}