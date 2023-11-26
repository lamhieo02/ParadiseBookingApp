package accountusecase

import (
	"context"
	"paradise-booking/worker"
	"time"

	"github.com/hibiken/asynq"
)

func (uc *accountUseCase) ForgotPassword(ctx context.Context, email string) error {
	// send email

	taskPayload := worker.PayloadSendVerifyResetCodePassword{
		Email: email,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueSendResetCodePassword),
	}
	_ = uc.taskDistributor.DistributeTaskSendVerifyResetCodePassword(ctx, &taskPayload, opts...)
	return nil
}
