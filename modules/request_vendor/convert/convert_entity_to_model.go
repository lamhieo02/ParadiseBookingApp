package requestvendorconvert

import (
	"paradise-booking/entities"
	requestvendoriomodel "paradise-booking/modules/request_vendor/iomodel"
)

func ConvertRequestVendorEntityToModel(entity *entities.RequestVendor) *requestvendoriomodel.GetRequestVendorResp {
	return &requestvendoriomodel.GetRequestVendorResp{
		ID:          entity.Id,
		UserID:      entity.UserId,
		FullName:    entity.FullName,
		Username:    entity.Username,
		Email:       entity.Email,
		Phone:       entity.Phone,
		DOB:         entity.DOB,
		Address:     entity.Address,
		Description: entity.Description,
		Experience:  entity.Experience,
		Status:      entity.Status,
	}
}
