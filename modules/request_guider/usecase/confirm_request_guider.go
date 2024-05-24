package requestguiderusecase

import (
	"context"
	"errors"
	"paradise-booking/constant"
	"paradise-booking/entities"
)

func (uc *requestGuiderUC) ConfirmRequestGuider(ctx context.Context, requestGuiderID int, typeConfirm int) error {

	requestGuider, err := uc.requestGuiderSto.GetByID(ctx, requestGuiderID)
	if err != nil {
		return err
	}

	if typeConfirm == constant.RequestGuiderTypeConfirmAccept {
		// update role user to guider
		accountUpdate := &entities.Account{
			Role: int(constant.GuiderRole),
		}
		if err := uc.accountSto.UpdateAccountById(ctx, requestGuider.UserId, accountUpdate); err != nil {
			return err
		}

		if err := uc.requestGuiderSto.UpdateWithMap(ctx, requestGuider, map[string]interface{}{
			"status": constant.RequestGuiderStatusSuccess,
		}); err != nil {
			return err
		}

	} else if typeConfirm == constant.RequestGuiderTypeConfirmReject {
		if err := uc.requestGuiderSto.UpdateWithMap(ctx, requestGuider, map[string]interface{}{
			"status": constant.RequestGuiderStatusReject,
		}); err != nil {
			return err
		}

	} else {
		return errors.New("invalid type confirm")
	}

	return nil
}
