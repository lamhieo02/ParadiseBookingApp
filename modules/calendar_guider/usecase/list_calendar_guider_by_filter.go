package calendarguiderusecase

import (
	"context"
	"paradise-booking/common"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
	"paradise-booking/utils"
)

func (uc *calendarGuiderUC) ListCalendarGuiderByFilter(ctx context.Context, paging *common.Paging, filter *calendarguideriomodel.Filter) ([]calendarguideriomodel.GetCalendarGuiderResp, error) {
	paging.Process()

	result, err := uc.calendarGuiderSto.ListByFilter(ctx, paging, filter)
	if err != nil {
		return nil, common.ErrCannotListEntity("calendar guider", err)
	}

	var res []calendarguideriomodel.GetCalendarGuiderResp
	for _, calendarGuider := range result {
		res = append(res, calendarguideriomodel.GetCalendarGuiderResp{
			ID:          calendarGuider.Id,
			PostGuideID: calendarGuider.PostGuideId,
			GuiderID:    calendarGuider.GuiderId,
			Note:        calendarGuider.Note,
			DateFrom:    utils.ParseTimeWithHourToString(calendarGuider.DateFrom),
			DateTo:      utils.ParseTimeWithHourToString(calendarGuider.DateTo),
			Price:       calendarGuider.PricePerPerson,
			Status:      calendarGuider.Status,
		})
	}

	return res, nil
}
