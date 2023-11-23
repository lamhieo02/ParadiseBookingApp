package worker

import (
	"context"
	"log"
	"paradise-booking/entities"

	"github.com/hibiken/asynq"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type AccountStorage interface {
	GetAccountByEmail(ctx context.Context, email string) (account *entities.Account, err error)
}

type redisTaskProcessor struct {
	server     *asynq.Server
	accountSto AccountStorage
}

func NewRedisTaskProcessor(redisOpt *asynq.RedisClientOpt, accountSto AccountStorage) TaskProcessor {
	server := asynq.NewServer(redisOpt, asynq.Config{})
	return &redisTaskProcessor{server: server, accountSto: accountSto}
}

func (processor *redisTaskProcessor) Start() error {
	log.Println("Starting task processor...")
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)

	return processor.server.Start(mux)

}
