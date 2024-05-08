package calendarguideriomodel

import (
	"paradise-booking/entities"
	"paradise-booking/utils"
)

type CreateCalendarGuiderReq struct {
	PostGuideID    int     `json:"post_guide_id"`
	GuiderID       int     `json:"guider_id"`
	Note           string  `json:"note"`
	DateFrom       string  `json:"date_from"`
	DateTo         string  `json:"date_to"`
	PricePerPerson float64 `json:"price_per_person"`
	MaxGuest       int     `json:"max_guest"`
}

func (req *CreateCalendarGuiderReq) ToEntity() (*entities.CalendarGuider, error) {
	dateFrom, err := utils.ParseStringToTimeWithHour(req.DateFrom)
	if err != nil {
		return nil, err
	}

	dateTo, err := utils.ParseStringToTimeWithHour(req.DateTo)
	if err != nil {
		return nil, err
	}

	return &entities.CalendarGuider{
		PostGuideId:    req.PostGuideID,
		GuiderId:       req.GuiderID,
		Note:           req.Note,
		DateFrom:       dateFrom,
		DateTo:         dateTo,
		PricePerPerson: int(req.PricePerPerson),
		Status:         true, // default when create for the first time
	}, nil
}
