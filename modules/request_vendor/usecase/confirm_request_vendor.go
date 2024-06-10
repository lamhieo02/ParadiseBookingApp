package requestvendorusecase

import (
	"context"
	"errors"
	"paradise-booking/constant"
	"paradise-booking/entities"
)

func (uc *requestVendorUC) ConfirmRequestVendor(ctx context.Context, requestVendorID int, typeConfirm int) error {

	requestVendor, err := uc.requestVendorSto.GetByID(ctx, requestVendorID)
	if err != nil {
		return err
	}

	if typeConfirm == constant.RequestVendorTypeConfirmAccept {
		// update role user to vendor
		accountUpdate := &entities.Account{
			Role: int(constant.VendorRole),
		}
		if err := uc.accountSto.UpdateAccountById(ctx, requestVendor.UserId, accountUpdate); err != nil {
			return err
		}

		if err := uc.requestVendorSto.UpdateWithMap(ctx, requestVendor, map[string]interface{}{
			"status": constant.RequestVendorStatusSuccess,
		}); err != nil {
			return err
		}

	} else if typeConfirm == constant.RequestVendorTypeConfirmReject {
		if err := uc.requestVendorSto.UpdateWithMap(ctx, requestVendor, map[string]interface{}{
			"status": constant.RequestVendorStatusReject,
		}); err != nil {
			return err
		}

	} else {
		return errors.New("invalid type confirm")
	}

	return nil
}
