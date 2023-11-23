package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

const (
	TaskSendVerifyEmail = "task:send_verify_email"
)

type PayloadSendVerifyEmail struct {
	Email string `json:"email"`
}

func (distributor *redisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error when marshal payload: %v", err)
	}
	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("error when enqueue task: %v", err)
	}

	fmt.Print(ctx, "Enqueued task: %v", info)
	return nil
}

func (processor *redisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("error when unmarshal payload: %w", asynq.SkipRetry)
	}

	account, err := processor.accountSto.GetAccountByEmail(ctx, payload.Email)
	if err == gorm.ErrRecordNotFound {
		return fmt.Errorf("account with email %s not found: %w", payload.Email, asynq.SkipRetry)
	}
	if err != nil {
		return fmt.Errorf("error when get account by email: %w", err)
	}

	//TODO: send verify email

	fmt.Println(ctx, "Process task: %v, sending email %s", task, account.Email)
	return nil
}
