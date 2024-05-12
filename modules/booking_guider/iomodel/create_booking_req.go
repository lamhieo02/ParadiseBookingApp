package bookingguideriomodel

import "paradise-booking/entities"

type CreateBookingReq struct {
	CalendarGuiderID int    `json:"calendar_guider_id"`
	Email            string `json:"email"`
	NumberOfPeople   int    `json:"number_of_people"`
	Name             string `json:"name"`
	Note             string `json:"note"`
	// Status           int    `json:"status"`
	TotalPrice    float64 `json:"total_price"`
	Phone         string  `json:"phone"`
	PaymentMethod int     `json:"payment_method"`
}

type CreateBookingResp struct {
	PaymentUrl string `json:"payment_url"`
}

func (req *CreateBookingReq) ToEntity() *entities.BookingGuider {
	return &entities.BookingGuider{
		CalendarGuiderID: req.CalendarGuiderID,
		Email:            req.Email,
		NumberOfPeople:   req.NumberOfPeople,
		Name:             req.Name,
		Note:             req.Note,
		// Status:           req.Status,
		TotalPrice:    float64(req.TotalPrice),
		Phone:         req.Phone,
		PaymentMethod: req.PaymentMethod,
	}
}
