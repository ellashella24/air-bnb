package booking

import (
	"air-bnb/entities"
	"fmt"
	"github.com/pkg/errors"
	"github.com/xendit/xendit-go"
	"gorm.io/gorm"
)

type Booking interface {
	CreateBooking(booking entities.Booking) (entities.Booking, error)
	FindHomeStay(days, id int) ([]xendit.InvoiceItem, int, error)
	FindCustomer(id int) (xendit.InvoiceCustomer, error)
	Update(bookingID string, booking entities.Booking) (entities.Booking, error)
	FindBookingByUserID(id int) ([]entities.Booking, error)
	FindBookingByHostID(id int) ([]entities.Booking, error)
	Checkout(invoiceID string, hostID int) (entities.Homestay, error)
}

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{db}
}

func (br *BookingRepository) CreateBooking(booking entities.Booking) (entities.Booking, error) {
	var homestay entities.Homestay
	if err := br.db.Create(&booking).Error; err != nil {
		return booking, err
	}

	br.db.Model(&homestay).Where("id = ?", booking.HomestayID).Update("booking_status", "not available")

	return booking, nil
}

func (br *BookingRepository) FindHomeStay(days, id int) ([]xendit.InvoiceItem, int, error) {
	var item []xendit.InvoiceItem
	var homeStay entities.Homestay

	err := br.db.Where("id = ? and booking_status =?", id, "available").First(&homeStay).Error
	if err != nil {

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

func (br *BookingRepository) Update(invoiceID string, booking entities.Booking) (entities.Booking, error) {

	if booking.PaymentStatus == "EXPIRED" {
		var bookingData entities.Booking
		var homestay entities.Homestay
		br.db.Where("invoice_id = ?", invoiceID).First(&bookingData)
		err := br.db.Model(&homestay).Where("id = ?", bookingData.HomestayID).Update(
			"booking_status", "available",
		).Error

		if err != nil || bookingData.ID == 0 {
			return bookingData, errors.New("Not Found")
		}
	}

	err := br.db.Where("invoice_id = ?", invoiceID).Model(&booking).Updates(booking).Error

	if err != nil || booking.PaymentStatus == "PENDING" {
		return booking, errors.New("Not found")
	}

	return booking, nil
}

func (br *BookingRepository) FindBookingByUserID(id int) ([]entities.Booking, error) {
	var booking []entities.Booking

	err := br.db.Where("user_id = ?", id).Find(&booking).Error

	if err != nil || len(booking) == 0 {
		return booking, errors.New("Not Found")
	}

	return booking, nil

}

func (br *BookingRepository) FindBookingByHostID(id int) ([]entities.Booking, error) {
	var booking []entities.Booking

	err := br.db.Table("bookings").Joins("join homestays on bookings.homestay_id = homestays.id").Where(
		"homestays.host_id = ?", id,
	).Find(&booking).Error

	if err != nil || len(booking) == 0 {
		return booking, errors.New("not found")
	}
	fmt.Println("ini booking", booking)

	return booking, nil
}

func (br *BookingRepository) Checkout(invoiceID string, hostID int) (entities.Homestay, error) {
	var homestay entities.Homestay
	var booking entities.Booking

	err := br.db.Where("invoice_id = ?", invoiceID).First(&booking).Error
	if err != nil {
		return homestay, errors.New("not found")
	}

	br.db.Model(&homestay).Where("id = ? and host_id = ?", booking.HomestayID, hostID).Update(
		"booking_status", "available",
	)

	return homestay, nil

}
