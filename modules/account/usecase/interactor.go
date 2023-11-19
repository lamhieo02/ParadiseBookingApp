package accountusecase

import (
	"context"
	"paradise-booking/config"
	"paradise-booking/entities"
)

type AccountStorage interface {
	Create(ctx context.Context, account *entities.Account) (err error)
	GetAccountByEmail(ctx context.Context, email string) (account *entities.Account, err error)
	UpdateAccountById(ctx context.Context, id int, accountUpdate *entities.Account) error
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type accountUseCase struct {
	accountStorage AccountStorage
	cfg            *config.Config
}

func NewUserUseCase(cfg *config.Config, accountSto AccountStorage) *accountUseCase {
	return &accountUseCase{accountSto, cfg}
}
