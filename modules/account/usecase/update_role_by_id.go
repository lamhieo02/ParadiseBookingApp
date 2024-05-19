package accountusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
	"sync"
	"time"
)

func (uc *accountUseCase) UpdateAccountRoleByID(ctx context.Context, accountModel *iomodel.AccountChangeRole, id int) (err error) {

	model := entities.Account{
		Role: accountModel.Role,
	}
	err = uc.accountStorage.UpdateAccountById(ctx, id, &model)
	if err != nil {
		return common.ErrInternal(err)
	}

	account, err := uc.accountStorage.GetProfileByID(ctx, id)
	if err != nil {
		return err
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)

	var errCache error
	go func() {
		defer wg.Done()
		if err := uc.redisCache.Set(ctx, account.CacheKeyID(), &account, 24*5*time.Hour); err != nil {
			errCache = err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := uc.redisCache.Set(ctx, account.CacheKeyEmail(), &account, 24*5*time.Hour); err != nil {
			errCache = err
			return
		}
	}()

	wg.Wait()
	if errCache != nil {
		return errCache
	}

	return nil
}
