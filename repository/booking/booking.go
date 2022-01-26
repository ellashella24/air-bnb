package booking

import (
	"air-bnb/entities"
	"github.com/labstack/gommon/log"
	"github.com/xendit/xendit-go"
	"gorm.io/gorm"
)

type Booking interface {
	CreateBooking(booking entities.Booking) (entities.Booking, error)
	FindHomeStay(days, id int) ([]xendit.InvoiceItem, int, error)
	FindCustomer(id int) (xendit.InvoiceCustomer, error)
	Update(bookingID string, booking entities.Booking) (entities.Booking, error)
}

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{db}
}

func (br *BookingRepository) CreateBooking(booking entities.Booking) (entities.Booking, error) {
	var homestay entities.Homestay

	if err := br.db.Where(
		"booking_status = ? and id =?", "available", booking.HomestayID,
	).First(homestay).Error; err != nil {
		return booking, err
	}
	if err := br.db.Create(&booking).Error; err != nil {
		return booking, err
	}

	return booking, nil
}

func (br *BookingRepository) FindHomeStay(days, id int) ([]xendit.InvoiceItem, int, error) {
	var item []xendit.InvoiceItem
	var homeStay entities.Homestay

	err := br.db.Where("id = ?", id).First(&homeStay).Error
	if err != nil {
		log.Error(err)
		return item, 0, err
	}

	item = append(
		item, xendit.InvoiceItem{
			Name:     homeStay.Name,
			Price:    float64(homeStay.Price),
			Quantity: days,
		},
	)

	return item, homeStay.Price, nil
}

func (br *BookingRepository) FindCustomer(id int) (xendit.InvoiceCustomer, error) {
	var customer xendit.InvoiceCustomer
	var user entities.User

	err := br.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return customer, err
	}

	customer.GivenNames = user.Name
	customer.Email = user.Email

	return customer, nil
}

func (br *BookingRepository) Update(bookingID string, booking entities.Booking) (entities.Booking, error) {

	if err := br.db.Where("invoice_id = ?", bookingID).Model(&booking).Updates(booking).Error; err != nil {
		return booking, err
	}

	return booking, nil
}
