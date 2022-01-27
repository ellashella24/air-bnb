package booking

import (
	"air-bnb/configs"
	"air-bnb/entities"
	"air-bnb/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestSetup(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.City{})
	db.Migrator().DropTable(&entities.Homestay{})
	db.Migrator().DropTable(&entities.Booking{})

	db.AutoMigrate(entities.User{})
	db.AutoMigrate(entities.City{})
	db.AutoMigrate(entities.Homestay{})
	db.AutoMigrate(entities.Booking{})

	var user entities.User
	user = entities.User{
		Model: gorm.Model{},
		Name:  "naufal",
		Email: "naufal@gmail.com",
		Role:  "user",
	}

	var city entities.City
	city = entities.City{
		Name:       "indramayu",
		HomestayID: nil,
	}

	var homestay entities.Homestay
	homestay = entities.Homestay{
		Name:           "koskosan",
		Price:          100000,
		Booking_Status: "available",
		HostID:         1,
		CityID:         1,
	}

	db.Create(&user)
	db.Create(&city)
	db.Create(&homestay)

}

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	bookingRepo := NewBookingRepository(db)

	t.Run(
		"succes case", func(t *testing.T) {

			checkIn, _ := time.Parse("2022-01-01", "2022-01-27")
			checkOut, _ := time.Parse("2022-01-01", "2022-01-28")

			mockRequest := entities.Booking{
				User_id: 1, HomestayID: 1, CheckIn: checkIn, CheckOut: checkOut, PaymentStatus: "PENDING",
			}
			res, err := bookingRepo.CreateBooking(mockRequest)
			assert.Nil(t, err)
			assert.Equal(t, uint(1), res.User_id)

		},
	)

	t.Run(
		"error case", func(t *testing.T) {
			checkIn, _ := time.Parse("2022-01-01", "2022-01-27")
			checkOut, _ := time.Parse("2022-01-01", "2022-01-28")

			mockRequest := entities.Booking{CheckIn: checkIn, CheckOut: checkOut}
			_, err := bookingRepo.CreateBooking(mockRequest)

			assert.NotNil(t, err)

		},
	)
}

func TestFindHomeStay(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	bookingRepo := NewBookingRepository(db)
	var homestay entities.Homestay
	db.Model(homestay).Where("id = ?", 1).Update("booking_status", "available")

	t.Run(
		"succes case", func(t *testing.T) {

			mockRequest := entities.Homestay{ID: 1, Price: 100000, Booking_Status: "available"}

			res, _, _ := bookingRepo.FindHomeStay(int(mockRequest.ID), 1)

			assert.Equal(t, 1, res[0].Quantity)

		},
	)

	t.Run(
		"error case", func(t *testing.T) {
			mockRequest := entities.Homestay{ID: 10}
			db.Model(mockRequest).Where("id = ?", mockRequest.ID).Update("booking_status", "available")
			_, _, err := bookingRepo.FindHomeStay(1, 100)

			assert.NotNil(t, err)

		},
	)
}

func TestFindCustomer(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	bookingRepo := NewBookingRepository(db)

	t.Run(
		"succes case", func(t *testing.T) {

			mockRequest := entities.User{ID: 1}

			res, _ := bookingRepo.FindCustomer(int(mockRequest.ID))

			assert.Equal(t, "naufal", res.GivenNames)

		},
	)

	t.Run(
		"error case", func(t *testing.T) {

			mockRequest := entities.User{ID: 10}

			_, err := bookingRepo.FindCustomer(int(mockRequest.ID))

			assert.NotNil(t, err)

		},
	)
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	bookingRepo := NewBookingRepository(db)

	t.Run(
		"succes case", func(t *testing.T) {
			var booking entities.Booking
			db.Where("id = ?", 1).First(&booking)
			mockRequest := entities.Booking{
				PaymentStatus: "PAID",
			}

			res, _ := bookingRepo.Update(booking.InvoiceID, mockRequest)

			assert.Equal(t, "PAID", res.PaymentStatus)

		},
	)

	t.Run(
		"error case", func(t *testing.T) {
			var booking entities.Booking
			db.Where("id = ?", 1).First(&booking)
			booking = entities.Booking{
				PaymentStatus: "PAID",
			}
			db.Model(&booking).Where("id = ?", 1).Update("payment_status", "PENDING")

			_, err := bookingRepo.Update("7ggi8u", booking)

			assert.NotNil(t, err)

		},
	)

	t.Run(
		"error case expired", func(t *testing.T) {
			var booking entities.Booking
			db.Where("id = ?", 1).First(&booking)

			db.Model(&booking).Where("id = ?", 1).Update("payment_status", "EXPIRED")

			_, err := bookingRepo.Update("7ggi8u", booking)

			assert.NotNil(t, err)

		},
	)
}

func TestFindBookingByUserID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	bookingRepo := NewBookingRepository(db)

	t.Run(
		"succes case", func(t *testing.T) {

			mockRequest := entities.User{ID: 1}

			res, _ := bookingRepo.FindBookingByUserID(int(mockRequest.ID))

			assert.Equal(t, "EXPIRED", res[0].PaymentStatus)

		},
	)

	t.Run(
		"error case", func(t *testing.T) {

			mockRequest := entities.User{ID: 10}

			_, err := bookingRepo.FindBookingByUserID(int(mockRequest.ID))

			assert.NotNil(t, err)

		},
	)
}

func TestFindBookingByHostID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	bookingRepo := NewBookingRepository(db)

	t.Run(
		"succes case", func(t *testing.T) {

			mockRequest := entities.User{ID: 1}

			res, _ := bookingRepo.FindBookingByHostID(int(mockRequest.ID))

			assert.Equal(t, "EXPIRED", res[0].PaymentStatus)

		},
	)

	t.Run(
		"error case", func(t *testing.T) {

			mockRequest := entities.User{ID: 10}

			_, err := bookingRepo.FindBookingByHostID(int(mockRequest.ID))

			assert.NotNil(t, err)

		},
	)
}

func TestCheckout(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	bookingRepo := NewBookingRepository(db)

	t.Run(
		"succes case", func(t *testing.T) {
			var booking entities.Booking

			db.Model(booking).Where("id = ?", 1).Update("invoice_id", "1")

			res, _ := bookingRepo.Checkout("1", 1)

			assert.Equal(t, "available", res.Booking_Status)

		},
	)

	t.Run(
		"error case booking data ", func(t *testing.T) {
			var booking entities.Booking

			db.Model(booking).Where("id = ?", 1).Update("invoice_id", "1")

			_, err := bookingRepo.Checkout("10", 10)

			assert.NotNil(t, err)

		},
	)

}

func TestReschedule(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	bookingRepo := NewBookingRepository(db)

	t.Run(
		"succes case", func(t *testing.T) {

			checkIn, _ := time.Parse("2006-01-02", "2022-02-01")

			res, _ := bookingRepo.Reschedule(1, "1", checkIn)

			assert.Equal(t, time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC), res.CheckOut)

		},
	)

	t.Run(
		"error case", func(t *testing.T) {

			checkIn, _ := time.Parse("2022-01-01", "2022-02-01")

			_, err := bookingRepo.Reschedule(1, "100", checkIn)

			assert.NotNil(t, err)

		},
	)

}
