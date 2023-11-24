package verifyemailsstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *verifyEmailsStorage) Get(ctx context.Context, email string, verifyCode string) (*entities.VerifyEmail, error) {
	db := s.db
	data := &entities.VerifyEmail{}
	err := db.Table(entities.VerifyEmail{}.TableName()).Where("email = ? AND scret_code = ?", email, verifyCode).First(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
