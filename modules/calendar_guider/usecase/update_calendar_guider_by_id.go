package calendarguiderusecase

import (
	"context"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
)

func (uc *calendarGuiderUC) UpdateCalendarGuiderByID(ctx context.Context, id int, calendarGuiderModel *calendarguideriomodel.UpdateCalendarGuiderReq) error {
	entity, err := calendarGuiderModel.ToEntity()
	if err != nil {
		return err
	}

	err = uc.calendarGuiderSto.UpdateByID(ctx, id, entity)
	if err != nil {
		return err
	}

	return nil
}
