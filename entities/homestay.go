package entities

import "gorm.io/gorm"

type Homestay struct {
	gorm.Model
	ID             uint
	Name           string
	Price          int
	Booking_Status string
	HostID         uint
	City_id        uint
	BookingID      []Booking
}
