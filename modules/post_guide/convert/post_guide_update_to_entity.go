package postguideconvert

import (
	"paradise-booking/entities"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
	"strings"
)

func ConvertPostGuideUpdateToEntity(model *postguideiomodel.UpdatePostGuideReq) *entities.PostGuide {
	return &entities.PostGuide{
		TopicID:     model.TopicID,
		Title:       model.Title,
		Description: model.Description,
		Cover:       strings.Join(model.Images, ","),
		Lat:         model.Lat,
		Lng:         model.Lng,
		Address:     model.Address,
		Schedule:    model.Schedule,
		Languages:   strings.Join(model.Languages, ","), // Convert array to string
		Country:     model.Country,
		State:       model.State,
		District:    model.District,
	}
}
