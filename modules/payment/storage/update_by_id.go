package paymentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *paymentStorage) UpdateByID(ctx context.Context, id int, data *entities.Payment) error {
	db := s.db.Table(data.TableName())

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
