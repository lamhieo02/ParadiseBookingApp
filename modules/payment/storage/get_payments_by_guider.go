package paymentstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *paymentStorage) GetPaymentByGuider(ctx context.Context, guiderId int, paging *common.Paging) ([]entities.Payment, error) {
	var payments []entities.Payment

	db := s.db.Table(entities.Payment{}.TableName())

	if err := db.Raw("call GetPaymentsForGuider(?,?,?)", guiderId, paging.Page, paging.Limit).Scan(&payments).Error; err != nil {
		return nil, err
	}

	count := int64(0)
	if err := db.Raw("call GetPaymentsSizeOfGuider(?)", guiderId).Scan(&count).Error; err != nil {
		return nil, err
	}

	paging.Total = count

	return payments, nil
}
