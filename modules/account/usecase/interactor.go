package accountusecase

import (
	"context"
	"paradise-booking/config"
	"paradise-booking/entities"
	"paradise-booking/worker"
)

type AccountStorage interface {
	Create(ctx context.Context, account *entities.Account) (err error)
	GetAccountByEmail(ctx context.Context, email string) (account *entities.Account, err error)
	UpdateAccountById(ctx context.Context, id int, accountUpdate *entities.Account) error
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type accountUseCase struct {
	accountStorage  AccountStorage
	cfg             *config.Config
	taskDistributor worker.TaskDistributor
}

func NewUserUseCase(cfg *config.Config, accountSto AccountStorage, taskDistributor worker.TaskDistributor) *accountUseCase {
	return &accountUseCase{accountSto, cfg, taskDistributor}
}
