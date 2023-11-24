package cmdworker

import (
	"log"
	"paradise-booking/config"
	accountusecase "paradise-booking/modules/account/usecase"
	verifyemailsusecase "paradise-booking/modules/verify_emails/usecase"
	"paradise-booking/provider/mail"
	"paradise-booking/worker"

	"github.com/hibiken/asynq"
)

func RunTaskProcessor(redisOpt *asynq.RedisClientOpt, accountSto accountusecase.AccountStorage, cfg *config.Config, verifyEmailsUC verifyemailsusecase.VerifyEmailsUseCase) {
	mailer := mail.NewGmailSender(cfg.Email.EmailSenderName, cfg.Email.EmailSenderAddress, cfg.Email.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, accountSto, verifyEmailsUC, mailer)
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal(err)
	}
}
