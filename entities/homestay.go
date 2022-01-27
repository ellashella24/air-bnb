package entities

import "gorm.io/gorm"

type Homestay struct {
	gorm.Model
	ID             uint
	Name           string `gorm:"not null"`
	Price          int
	Booking_Status string `gorm:"default:available"`
	HostID         uint
	City_id         uint
	BookingID      []Booking
}
