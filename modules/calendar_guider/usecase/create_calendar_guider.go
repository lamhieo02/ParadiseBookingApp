package calendarguiderusecase

import (
	"context"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
)

func (uc *calendarGuiderUC) CreateCalendarGuider(ctx context.Context, data *calendarguideriomodel.CreateCalendarGuiderReq) error {

	entity, err := data.ToEntity()
	if err != nil {
		return err
	}

	if err := uc.calendarGuiderSto.Create(ctx, entity); err != nil {
		return err
	}

	return nil
}
