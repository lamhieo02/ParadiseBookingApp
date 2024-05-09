package calendarguiderusecase

import "context"

func (uc *calendarGuiderUC) DeleteCalendarGuiderByID(ctx context.Context, id int) error {
	err := uc.calendarGuiderSto.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
