package calendarguiderusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
)

type CalendarGuiderStorage interface {
	Create(ctx context.Context, data *entities.CalendarGuider) error
	DeleteByID(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*entities.CalendarGuider, error)
	ListByFilter(ctx context.Context, paging *common.Paging, filter *calendarguideriomodel.Filter) ([]*entities.CalendarGuider, error)
}

type calendarGuiderUC struct {
	calendarGuiderSto CalendarGuiderStorage
}

func NewCalendarGuiderUseCase(calendarGuiderSto CalendarGuiderStorage) *calendarGuiderUC {
	return &calendarGuiderUC{calendarGuiderSto}
}
