package bookingguideriomodel

import "time"

type GetBookingGuiderResp struct {
	ID               int            `json:"id"`
	CalendarGuiderID int            `json:"calendar_guider_id"`
	CalendarGuider   CalendarGuider `json:"calendar_guider"`
	Email            string         `json:"email"`
	NumberOfPeople   int            `json:"number_of_people"`
	Name             string         `json:"name"`
	Note             string         `json:"note"`
	StatusID         int            `json:"status_id"`
	Status           string         `json:"status"`
	TotalPrice       int            `json:"total_price"`
	Phone            string         `json:"phone"`
	CreatedAt        time.Time      `json:"created_at"`
	PaymentMethod    string         `json:"payment_method"`
}

type CalendarGuider struct {
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
}
