package entities

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID            uint
	User_id       uint
	CheckIn       time.Time
	CheckOut      time.Time
	PaymentStatus string
	InvoiceID     string
	PaymentMethod string
	BankID        string
	PaymentURL    string
	PaidAt        string
	PriceBooking  int
	HomestayID    uint
}
