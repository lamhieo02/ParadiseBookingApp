package placeusecase

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"time"

	"github.com/samber/lo"
)

func (uc *placeUseCase) GetDatesBookedPlace(ctx context.Context, placeId int) ([][]string, error) {
	conditions := []common.Condition{}
	conditions = append(conditions, common.Condition{
		Field:    "place_id",
		Operator: common.OperatorEqual,
		Value:    placeId,
	})
	conditions = append(conditions, common.Condition{
		Field:    "status_id",
		Operator: common.OperatorNotEqual,
		Value:    constant.BookingStatusCancel,
	})
	conditions = append(conditions, common.Condition{
		Field:    "status_id",
		Operator: common.OperatorNotEqual,
		Value:    constant.BookingStatusCompleted,
	})

	bookings, err := uc.bookingSto.ListAllBookingWithCondition(ctx, conditions)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	layout := "02-01-2006"

	dates := lo.FilterMap(bookings, func(item entities.Booking, _ int) ([]string, bool) {
		t, err := time.Parse(layout, item.CheckInDate)
		if err != nil {
			log.Printf("Error parse time: %v", err)
		}
		if t.After(timeNow) {
			return []string{item.CheckInDate, item.ChekoutDate}, true
		}
		return []string{}, false
	})
	return dates, nil
}
