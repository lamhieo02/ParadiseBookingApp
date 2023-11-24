package verifyemailsusecase

import (
	"context"
	"paradise-booking/entities"
)

type verifyEmailsStorage interface {
	Create(ctx context.Context, data *entities.VerifyEmail) (*entities.VerifyEmail, error)
	Get(ctx context.Context, email string, verifyCode string) (*entities.VerifyEmail, error)
}

type VerifyEmailsUseCase interface {
	CreateVerifyEmails(ctx context.Context, email string) (*entities.VerifyEmail, error)
}

type AccountStorage interface {
	UpdateIsVerifyEmailByEmail(ctx context.Context, email string) error
}

type verifyEmailsUseCase struct {
	verifyEmailsStore verifyEmailsStorage
	accountStore      AccountStorage
}

func NewVerifyEmailsUseCase(verifyEmailsStore verifyEmailsStorage) *verifyEmailsUseCase {
	return &verifyEmailsUseCase{verifyEmailsStore: verifyEmailsStore}
}
