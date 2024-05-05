package calendarguiderusecase

import (
	"context"
	"paradise-booking/entities"
)

type CalendarGuiderStorage interface {
	Create(ctx context.Context, data *entities.CalendarGuider) error
	DeleteByID(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*entities.CalendarGuider, error)
}

type calendarGuiderUC struct {
	calendarGuiderSto CalendarGuiderStorage
}

func NewCalendarGuiderUseCase(calendarGuiderSto CalendarGuiderStorage) *calendarGuiderUC {
	return &calendarGuiderUC{calendarGuiderSto}
}
