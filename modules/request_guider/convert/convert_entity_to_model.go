package requestguiderconvert

import (
	"paradise-booking/entities"
	requestguideriomodel "paradise-booking/modules/request_guider/iomodel"
	"strings"
)

func ConvertRequestGuiderEntityToModel(entity *entities.RequestGuider) *requestguideriomodel.GetRequestGuiderResp {
	return &requestguideriomodel.GetRequestGuiderResp{
		ID:            entity.Id,
		UserID:        entity.UserId,
		FullName:      entity.FullName,
		Username:      entity.Username,
		Email:         entity.Email,
		Phone:         entity.Phone,
		DOB:           entity.DOB,
		Address:       entity.Address,
		Description:   entity.Description,
		Experience:    entity.Experience,
		Reason:        entity.Reason,
		GoalsOfTravel: strings.Split(entity.GoalOfTravel, ","),
		Languages:     strings.Split(entity.Languages, ","),
		Status:        entity.Status,
	}
}
