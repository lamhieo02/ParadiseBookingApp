package calendarguiderusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
	"paradise-booking/utils"
	"sync"

	"github.com/samber/lo"
)

func (uc *calendarGuiderUC) ListCalendarGuiderByFilter(ctx context.Context, paging *common.Paging, filter *calendarguideriomodel.Filter) ([]calendarguideriomodel.GetCalendarGuiderResp, error) {
	paging.Process()

	var (
		err            error
		wgErr          error
		result         []*entities.CalendarGuider
		bookingGuiders []*entities.BookingGuider
	)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		result, err = uc.calendarGuiderSto.ListByFilter(ctx, paging, filter)
		if err != nil {
			wgErr = err
			return
		}
	}()

	calendarGuiderIDs := lo.Map(result, func(x *entities.CalendarGuider, _ int) int {
		return x.Id
	})

	conditions := []common.Condition{
		{
			Field:    "calendar_guider_id",
			Operator: common.OperatorIn,
			Value:    calendarGuiderIDs,
		},
		{
			Field:    "status_id",
			Operator: common.OperatorIn,
			Value:    []int{constant.BookingGuiderStatusPending, constant.BookingGuiderStatusConfirmed},
		},
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		bookingGuiders, err = uc.bookingGuiderSto.ListByCondition(ctx, conditions)
		if err != nil {
			wgErr = err
			return
		}
	}()

	wg.Wait()

	if wgErr != nil {
		return nil, wgErr
	}

	mapCalendarIDWithNumGuest := make(map[int]int)
	for _, bookingGuider := range bookingGuiders {
		mapCalendarIDWithNumGuest[bookingGuider.CalendarGuiderID] += bookingGuider.NumberOfPeople
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
			MaxGuest:    calendarGuider.MaxGuest - mapCalendarIDWithNumGuest[calendarGuider.Id],
		})
	}

	return res, nil
}
