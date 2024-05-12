package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"paradise-booking/constant"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	TaskSendConfirmBookingGuider = "task:send_confirm_booking_guider"
	UrlConfirmBookingGuider      = constant.URL_HOST_PROD + "/confirm_booking_guider"
)

type PayloadSendConfirmBookingGuider struct {
	Email           string `json:"email"`
	BookingGuiderID int    `json:"booking_guider_id"`
	FullName        string `json:"full_name"`
}

func (distributor *redisTaskDistributor) DistributeTaskSendConfirmBookingGuider(
	ctx context.Context,
	payload *PayloadSendConfirmBookingGuider,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error when marshal payload: %v", err)
	}
	task := asynq.NewTask(TaskSendConfirmBookingGuider, jsonPayload, opts...)

	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("error when enqueue task: %v", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *redisTaskProcessor) ProcessTaskSendConfirmBookingGuider(ctx context.Context, task *asynq.Task) error {
	log.Info().Msg("process task send confirm booking guider")
	var payload PayloadSendConfirmBookingGuider
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("error when unmarshal payload: %w", asynq.SkipRetry)
	}

	// account, err := processor.accountSto.GetAccountByEmail(ctx, payload.Email)
	// if err == gorm.ErrRecordNotFound {
	// 	return fmt.Errorf("account with email %s not found: %w", payload.Email, asynq.SkipRetry)
	// }
	// if err != nil {
	// 	return fmt.Errorf("error when get account by email: %w", err)
	// }
	infoCustomer := &InfoCustomer{
		FullName: payload.FullName,
		Email:    payload.Email,
	}

	sendMailToVerifyBookingGuider(processor, infoCustomer, payload.BookingGuiderID)
	log.Info().Msg("send verify booking success")

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", infoCustomer.Email).Msg("processed task")

	return nil
}

func sendMailToVerifyBookingGuider(processor *redisTaskProcessor, customer *InfoCustomer, bookingGuiderID int) error {
	subject := "Welcome to Paradise Booking"
	verifyUrl := fmt.Sprintf("%s?booking_guider_id=%d&status=%d",
		UrlConfirmBookingGuider, bookingGuiderID, constant.BookingGuiderStatusConfirmed)
	content := fmt.Sprintf(`Hello %s,<br/>
	Thank you for booking guider with us!<br/>
	Please <a href="%s">click here</a> to confirm your booking.<br/>
	`, customer.FullName, verifyUrl)
	to := []string{customer.Email}

	err := processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send confirm booking guider: %w", err)
	}
	return nil
}
