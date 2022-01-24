package entities

import "gorm.io/gorm"

type City struct {
	gorm.Model
	ID         uint
	Name       string
	HomestayID []Homestay
}
