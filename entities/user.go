package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Email    string
	Password string
	Role     string `gorm:"default:user"`
	Booking  []Booking
	Homestay []Homestay `gorm:"foreignKey:HostID"`
}
