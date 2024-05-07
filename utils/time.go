package utils

import (
	"fmt"
	"paradise-booking/constant"
	"time"
)

func ParseStringToTime(date string) (*time.Time, error) {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return nil, nil
	}

	dateRes, err := time.ParseInLocation(constant.LayoutDateTime, date, location)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, nil
	}
	return &dateRes, nil
}

func ParseTimeToString(date *time.Time) string {
	formattedTime := date.Format(constant.LayoutDateTime)
	return formattedTime
}

func ParseTimeWithHourToString(date *time.Time) string {
	formattedTime := date.Format(constant.LayoutDateTimeWithHour)
	return formattedTime
}

func ParseStringToTimeWithHour(date string) (*time.Time, error) {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return nil, nil
	}

	dateRes, err := time.ParseInLocation(constant.LayoutDateTimeWithHour, date, location)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, nil
	}

	return &dateRes, nil
}
