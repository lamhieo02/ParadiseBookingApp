package placeusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"

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

	dates := lo.Map(bookings, func(item entities.Booking, _ int) []string {
		return []string{item.CheckInDate, item.ChekoutDate}
	})
	return dates, nil
}
