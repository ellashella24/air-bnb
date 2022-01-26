package booking

import (
	"air-bnb/entities"
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
		br.db.Table("homestays").Joins("join bookings on bookings.homestay_id = homestays.id").Where(
			"booking_status", "not available",
		).Update("booking_status", "available")
	}

	if err := br.db.Where("invoice_id = ?", invoiceID).Model(&booking).Updates(booking).Error; err != nil {
		return booking, err
	}

	return booking, nil
}

func (br *BookingRepository) FindBookingByUserID(id int) ([]entities.Booking, error) {
	var booking []entities.Booking

	br.db.Where("user_id = ?", id).Find(&booking)

	return booking, nil

}

func (br *BookingRepository) FindBookingByHostID(id int) ([]entities.Booking, error) {
	var booking []entities.Booking

	br.db.Table("bookings").Joins("join homestays on booking.homestay_id = homestays.id").Where(
		"homestays.host_id = ?", id,
	).Find(&booking)

	return booking, nil
}

func (br *BookingRepository) Checkout(invoiceID string, hostID int) (entities.Homestay, error) {
	var homestay entities.Homestay
	var booking entities.Booking

	err := br.db.Where("invoice_id = ?", invoiceID).First(&booking).Error

	if err != nil {
		return homestay, err
	}

	err = br.db.Model(homestay).Where("id = ? and host_id = ?", booking.HomestayID, hostID).Update(
		"booking_status", "available",
	).Error
	if err != nil {
		return homestay, err
	}

	return homestay, nil

}
