package calendarguideriomodel

import (
	"paradise-booking/entities"
	"paradise-booking/utils"
)

type UpdateCalendarGuiderReq struct {
	PostGuideID    int     `json:"post_guide_id"`
	GuiderID       int     `json:"guider_id"`
	Note           string  `json:"note"`
	DateFrom       string  `json:"date_from"`
	DateTo         string  `json:"date_to"`
	PricePerPerson float64 `json:"price_per_person"`
	Status         bool    `json:"status"`
	MaxGuest       int     `json:"max_guest"`
}

func (req *UpdateCalendarGuiderReq) ToEntity() (*entities.CalendarGuider, error) {
	res := &entities.CalendarGuider{
		PostGuideId:    req.PostGuideID,
		GuiderId:       req.GuiderID,
		Note:           req.Note,
		PricePerPerson: int(req.PricePerPerson),
		Status:         req.Status,
		MaxGuest:       req.MaxGuest,
	}

	if req.DateFrom != "" {
		dateFrom, err := utils.ParseStringToTimeWithHour(req.DateFrom)
		if err != nil {
			return nil, err
		}
		res.DateFrom = dateFrom
	}

	if req.DateTo != "" {
		dateTo, err := utils.ParseStringToTimeWithHour(req.DateTo)
		if err != nil {
			return nil, err
		}
		res.DateTo = dateTo
	}

	return res, nil
}
