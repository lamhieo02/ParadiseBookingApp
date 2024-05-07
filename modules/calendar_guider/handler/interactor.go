package calendarguiderhandler

import (
	"context"
	"paradise-booking/common"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
)

type calendarGuiderUseCase interface {
	CreateCalendarGuider(ctx context.Context, data *calendarguideriomodel.CreateCalendarGuiderReq) error
	GetCalendarGuiderByID(ctx context.Context, id int) (*calendarguideriomodel.GetCalendarGuiderResp, error)
	ListCalendarGuiderByFilter(ctx context.Context, paging *common.Paging, filter *calendarguideriomodel.Filter) ([]calendarguideriomodel.GetCalendarGuiderResp, error)
}

type calendarGuiderHandler struct {
	calendarGuiderUC calendarGuiderUseCase
}

func NewCalendarGuiderHandler(calendarGuiderUseCase calendarGuiderUseCase) *calendarGuiderHandler {
	return &calendarGuiderHandler{calendarGuiderUseCase}
}
