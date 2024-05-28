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
	UpdateByID(ctx context.Context, id int, postGuideData *entities.CalendarGuider) error
}

type BookingGuiderSto interface {
	ListByCondition(ctx context.Context, conditions []common.Condition) ([]*entities.BookingGuider, error)
}

type calendarGuiderUC struct {
	calendarGuiderSto CalendarGuiderStorage
	bookingGuiderSto  BookingGuiderSto
}

func NewCalendarGuiderUseCase(calendarGuiderSto CalendarGuiderStorage, bookingGuiderSto BookingGuiderSto) *calendarGuiderUC {
	return &calendarGuiderUC{calendarGuiderSto, bookingGuiderSto}
}
