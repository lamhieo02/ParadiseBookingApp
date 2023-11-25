package worker

import (
	"context"
	"paradise-booking/entities"
	"paradise-booking/provider/mail"

	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	QueueSendVerifyEmail = "send_verify_email"
	QueueDefault         = "default"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type AccountStorage interface {
	GetAccountByEmail(ctx context.Context, email string) (account *entities.Account, err error)
}

type VerifyEmailsUseCase interface {
	CreateVerifyEmails(ctx context.Context, email string) (*entities.VerifyEmail, error)
}

type redisTaskProcessor struct {
	server         *asynq.Server
	accountSto     AccountStorage
	verifyEmailsUC VerifyEmailsUseCase
	mailer         mail.EmailSender
}

func NewRedisTaskProcessor(redisOpt *asynq.RedisClientOpt, accountSto AccountStorage, verifyEmailsUC VerifyEmailsUseCase, mailer mail.EmailSender) TaskProcessor {

	logger := NewLogger()
	redis.SetLogger(logger)
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueSendVerifyEmail: 10,
				QueueDefault:         5,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Error().Err(err).Str("task type", task.Type()).
					Bytes("payload", task.Payload()).
					Msg("error when process task")
			}),
			Logger: logger,
		})
	return &redisTaskProcessor{server: server, accountSto: accountSto, verifyEmailsUC: verifyEmailsUC, mailer: mailer}
}

func (processor *redisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)

	return processor.server.Start(mux)

}