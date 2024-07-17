package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
	"paradise-booking/utils"
	"strings"
)

func ConvertPlaceEntityToModel(place *entities.Place) *iomodel.DataPlace {
	if place == nil {
		return nil
	}
	images := []string{}
	if place.Cover != "" {
		images = strings.Split(place.Cover, ",")
	}
	placeModel := iomodel.DataPlace{
		Id:            place.Id,
		VendorID:      place.VendorID,
		Name:          place.Name,
		Description:   place.Description,
		PricePerNight: place.PricePerNight,
		Address:       place.Address,
		Images:        images,
		Lat:           place.Lat,
		Lng:           place.Lng,
		Country:       place.Country,
		State:         place.State,
		District:      place.District,
		MaxGuest:      place.MaxGuest,
		NumBed:        place.NumBed,
		BedRoom:       place.BedRoom,
		CreatedAt:     place.CreatedAt,
		UpdatedAt:     place.UpdatedAt,
	}

	return &placeModel
}

func ConvertBookingModelToGetByPlaceResp(user *entities.Account, dataBooking *entities.Booking, place *entities.Place, bookingDetail *entities.BookingDetail) *iomodel.GetBookingByPlaceResp {
	// parse checkin and checkout date from string to time.Time
	checkInTime := utils.ParseTimeToString(dataBooking.CheckInDate)
	checkOutTime := utils.ParseTimeToString(dataBooking.ChekoutDate)

	placeModel := iomodel.DataPlace{
		Id:            place.Id,
		VendorID:      place.VendorID,
		Name:          place.Name,
		Description:   place.Description,
		PricePerNight: place.PricePerNight,
		Address:       place.Address,
		Images:        strings.Split(place.Cover, ","),
		Lat:           place.Lat,
		Lng:           place.Lng,
		Country:       place.Country,
		State:         place.State,
		District:      place.District,
		MaxGuest:      place.MaxGuest,
		NumBed:        place.NumBed,
		BedRoom:       place.BedRoom,
		CreatedAt:     place.CreatedAt,
		UpdatedAt:     place.UpdatedAt,
	}
	return &iomodel.GetBookingByPlaceResp{
		Id:              dataBooking.Id,
		CreatedAt:       dataBooking.CreatedAt,
		UpdatedAt:       dataBooking.UpdatedAt,
		UserId:          user.Id,
		User:            *user,
		PlaceId:         dataBooking.PlaceId,
		Place:           &placeModel,
		StatusId:        dataBooking.StatusId,
		CheckInDate:     checkInTime,
		ChekoutDate:     checkOutTime,
		GuestName:       bookingDetail.GuestName,
		TotalPrice:      bookingDetail.TotalPrice,
		NumberOfGuest:   bookingDetail.NumberOfGuest,
		ContentToVendor: bookingDetail.ContentToVendor,
		PaymentMethod:   bookingDetail.PaymentMethod,
	}
}
