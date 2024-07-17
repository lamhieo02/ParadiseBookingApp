package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"paradise-booking/constant"
	"paradise-booking/entities"

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

	bguider, err := processor.bookingGuiderSto.GetByID(ctx, payload.BookingGuiderID)
	if err != nil {
		return fmt.Errorf("error when get booking guider by id: %w", err)

	}

	calendar, err := processor.calendarGuiderSto.GetByID(ctx, bguider.CalendarGuiderID)
	if err != nil {
		return fmt.Errorf("error when get calendar guider by id: %w", err)
	}

	guider, err := processor.accountSto.GetProfileByID(ctx, bguider.GuiderID)
	if err != nil {
		return fmt.Errorf("error when get guider by id: %w", err)
	}

	sendMailToVerifyBookingGuider(processor, infoCustomer, bguider, calendar, guider)
	log.Info().Msg("send verify booking success")

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", infoCustomer.Email).Msg("processed task")

	return nil
}

func sendMailToVerifyBookingGuider(processor *redisTaskProcessor, customer *InfoCustomer, bookingGuider *entities.BookingGuider, calendar *entities.CalendarGuider, guider *entities.Account) error {
	subject := "Welcome to Paradise Booking"
	verifyUrl := fmt.Sprintf("%s?booking_guider_id=%d&status=%d",
		UrlConfirmBookingGuider, bookingGuider.Id, constant.BookingGuiderStatusConfirmed)

	date := calendar.DateFrom.Format("2006-01-02") + " - " + calendar.DateTo.Format("2006-01-02")
	content := fmt.Sprintf(`Hello %s,<br/>
	Thank you for booking guider with us!<br/>
	Here is your booking information:<br/>
	- Full name: %s<br/>
	- Email: %s<br/>
	- Date: %s<br/>
	- Total price: %d<br/>
	- Payment method: %s<br/>
	Here is guider information:<br/>
	- Full name: %s<br/>
	- Email: %s<br/>
	- Phone: %s<br/>
	- Address: %s<br/>
	To join the trip, you need to contact with guider through email or phone number.<br/>
	Please <a href="%s">click here</a> to confirm your booking.<br/>
	`, customer.FullName, customer.FullName, customer.Email, date, int(bookingGuider.TotalPrice), constant.MapPaymentMethod[bookingGuider.PaymentMethod], guider.FullName, guider.Email, guider.Phone, guider.Address, verifyUrl)
	to := []string{customer.Email}

	err := processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send confirm booking guider: %w", err)
	}
	return nil
}
