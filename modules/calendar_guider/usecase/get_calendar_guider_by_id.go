package calendarguiderusecase

import (
	"context"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
	"paradise-booking/utils"
)

func (uc *calendarGuiderUC) GetCalendarGuiderByID(ctx context.Context, id int) (*calendarguideriomodel.GetCalendarGuiderResp, error) {
	data, err := uc.calendarGuiderSto.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := calendarguideriomodel.GetCalendarGuiderResp{
		ID:          data.Id,
		PostGuideID: data.GuiderId,
		GuiderID:    data.GuiderId,
		Note:        data.Note,
		DateFrom:    utils.ParseTimeWithHourToString(data.DateFrom),
		DateTo:      utils.ParseTimeWithHourToString(data.DateTo),
		Price:       data.PricePerPerson,
		Status:      data.Status,
	}

	return &result, nil
}
