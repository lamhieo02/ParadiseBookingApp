package utils

import (
	"paradise-booking/constant"
	"time"
)

func ParseStringToTime(date string) (*time.Time, error) {
	// location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	// if err != nil {
	// 	fmt.Println("Error loading location:", err)
	// 	return nil, nil
	// }

	// dateRes, err := time.ParseInLocation(constant.LayoutDateTime, date, location)
	// if err != nil {
	// 	fmt.Println("Error parsing date:", err)
	// 	return nil, nil
	// }
	dateTime, err := time.Parse(constant.LayoutDateTime, date)
	if err != nil {
		return nil, err
	}

	// res := dateTime.Add(7 * time.Hour) // Add 7 hours to get the correct time in Ho Chi Minh timezone

	return &dateTime, nil
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
	// location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	// if err != nil {
	// 	fmt.Println("Error loading location:", err)
	// 	return nil, nil
	// }

	// dateRes, err := time.ParseInLocation(constant.LayoutDateTimeWithHour, date, location)
	// if err != nil {
	// 	fmt.Println("Error parsing date:", err)
	// 	return nil, nil
	// }
	dateTime, err := time.Parse(constant.LayoutDateTimeWithHour, date)
	if err != nil {
		return nil, err
	}

	// res := dateTime.Add(7 * time.Hour) // Add 7 hours to get the correct time in Ho Chi Minh timezone

	return &dateTime, nil
}
