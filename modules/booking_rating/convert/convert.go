package bookingratingconvert

import (
	"paradise-booking/entities"
	bookingratingiomodel "paradise-booking/modules/booking_rating/iomodel"
	"strings"
)

func ConvertPlaceEntityToModel(place *entities.Place) *bookingratingiomodel.DataPlace {
	return &bookingratingiomodel.DataPlace{
		ID:               place.Id,
		VendorID:         place.VendorID,
		Name:             place.Name,
		Description:      place.Description,
		PricePerNight:    place.PricePerNight,
		Address:          place.Address,
		Images:           strings.Split(place.Cover, ","),
		Lat:              place.Lat,
		Lng:              place.Lng,
		Country:          place.Country,
		State:            place.State,
		District:         place.District,
		MaxGuest:         place.MaxGuest,
		NumBed:           place.NumBed,
		BedRoom:          place.BedRoom,
		NumPlaceOriginal: place.NumPlaceOriginal,
	}
}

func ConvertPostGuideEntityToModel(postGuide *entities.PostGuide) *bookingratingiomodel.DataPostGuide {
	return &bookingratingiomodel.DataPostGuide{
		ID:          postGuide.Id,
		PostOwnerId: postGuide.PostOwnerId,
		TopicID:     postGuide.TopicID,
		Title:       postGuide.Title,
		Description: postGuide.Description,
		Images:      strings.Split(postGuide.Cover, ","), // Convert string to array of string
		Lat:         postGuide.Lat,
		Lng:         postGuide.Lng,
		Country:     postGuide.Country,
		State:       postGuide.State,
		District:    postGuide.District,
		Address:     postGuide.Address,
		Languages:   postGuide.Languages,
		Schedule:    postGuide.Schedule,
	}
}

func ConvertDataBookingRatingEntityToModel(dataBookingRating *entities.BookingRating) *bookingratingiomodel.DataBookingRating {
	return &bookingratingiomodel.DataBookingRating{
		Id:         dataBookingRating.Id,
		UserId:     dataBookingRating.UserId,
		BookingId:  dataBookingRating.BookingId,
		ObjectId:   dataBookingRating.ObjectId,
		ObjectType: dataBookingRating.ObjectType,
		Title:      dataBookingRating.Title,
		Content:    dataBookingRating.Content,
		Rating:     dataBookingRating.Rating,
	}
}
