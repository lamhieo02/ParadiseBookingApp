package storage

import (
	"context"
	"paradise-booking/entities"

	"gorm.io/gorm"
)

type userStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *userStorage {
	return &userStorage{db: db}
}

func (s *userStorage) Create(ctx context.Context, user *entities.User) (err error) {

	return nil
}
