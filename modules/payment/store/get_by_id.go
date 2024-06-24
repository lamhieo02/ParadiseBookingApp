package paymentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *paymentStorage) GetByID(ctx context.Context, id int) (*entities.Payment, error) {
	var payment entities.Payment

	if err := s.db.Where("id = ?", id).First(&payment).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}
