package cmdworker

import (
	"log"
	accountusecase "paradise-booking/modules/account/usecase"
	"paradise-booking/worker"

	"github.com/hibiken/asynq"
)

func RunTaskProcessor(redisOpt *asynq.RedisClientOpt, accountSto accountusecase.AccountStorage) {
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, accountSto)
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal(err)
	}
}
