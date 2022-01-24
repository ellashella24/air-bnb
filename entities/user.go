package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         uint
	Name       string
	Email      string
	Password   string
	Role       string `gorm:"default:user"`
	BookingID  []Booking
	HomestayID []Homestay
}
