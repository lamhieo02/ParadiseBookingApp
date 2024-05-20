package constant

const (
	BookingStatusPending = iota + 1
	BookingStatusConfirmed
	BookingStatusCheckIn
	BookingStatusCheckOut
	BookingStatusCompleted
	BookingStatusCancel
)

type TypeBooking int

const (
	TypeBookingForMySelf TypeBooking = iota + 1
	TypeBookingForFriend
)

const (
	BookingGuiderStatusPending = iota + 1
	BookingGuiderStatusConfirmed
	BookingGuiderStatusCompleted
	BookingGuiderStatusCancel
)

var MapBookingGuiderStatus = map[int]string{
	BookingGuiderStatusPending:   "Pending",
	BookingGuiderStatusConfirmed: "Confirmed",
	BookingGuiderStatusCompleted: "Completed",
	BookingGuiderStatusCancel:    "Cancel",
}

const (
	BookingRatingObjectTypePlace = 1
	BookingRatingObjectTypeGuide = 2
)
