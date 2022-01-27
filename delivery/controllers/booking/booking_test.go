package booking

import (
	"air-bnb/constants"
	"air-bnb/delivery/common"
	"air-bnb/entities"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xendit/xendit-go"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBooking(t *testing.T) {
	t.Run(
		"1. Success Booking Test", func(t *testing.T) {
			e := echo.New()

			requestBody, _ := json.Marshal(
				BookingRequest{
					HomeStayID: 1,
					CheckIn:    "2022-01-28",
					CheckOut:   "2022-01-29",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/")

			bookingController := NewBookingController(mockBookingRepository{})
			bookingController.Create()(context)

			var response common.ResponseSuccess

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Successful Operation", response.Message)
		},
	)
	t.Run(
		"2. Error Bad Request Checkin Booking Test", func(t *testing.T) {
			e := echo.New()

			requestBody, _ := json.Marshal(
				BookingRequest{
					HomeStayID: 1,
					CheckIn:    "2022-1-28",
					CheckOut:   "2022-01-29",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/")

			bookingController := NewBookingController(mockFalseBookingRepository{})
			bookingController.Create()(context)

			var response common.ResponseError

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Bad Request", response.Message)
		},
	)
	t.Run(
		"3. Error Bad Request Checkout Booking Test", func(t *testing.T) {
			e := echo.New()

			requestBody, _ := json.Marshal(
				BookingRequest{
					HomeStayID: 1,
					CheckIn:    "2022-01-28",
					CheckOut:   "2022-1-29",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/")

			bookingController := NewBookingController(mockFalseBookingRepository{})
			bookingController.Create()(context)

			var response common.ResponseError

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Bad Request", response.Message)
		},
	)
	t.Run(
		"4. Error Not Found Homestays or not available Checkout Booking Test", func(t *testing.T) {
			e := echo.New()

			requestBody, _ := json.Marshal(
				BookingRequest{

					CheckIn:  "2022-01-28",
					CheckOut: "2022-01-29",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/")

			bookingController := NewBookingController(mockFalseBookingRepository{})
			bookingController.Create()(context)

			var response common.ResponseError

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, 404, response.Code)
		},
	)
}

func TestCallback(t *testing.T) {
	t.Run(
		"1. Success Callback Booking Test", func(t *testing.T) {
			e := echo.New()

			requestBody, _ := json.Marshal(
				CallbackRequest{
					ExternalID:    "1",
					PaymentMethod: "TRANSFER_BANK",
					PaidAt:        "2020-01-27",
					Status:        "PAID",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			req.Header.Set("X-Callback-Token", constants.CallbackToken)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/callback")

			bookingController := NewBookingController(mockBookingRepository{})
			bookingController.Callback()(context)

			var response common.ResponseSuccess

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Successful Operation", response.Message)
		},
	)
	t.Run(
		"2. Fail Callback Booking Test", func(t *testing.T) {
			e := echo.New()

			requestBody, _ := json.Marshal(
				CallbackRequest{
					ExternalID:    "1",
					PaymentMethod: "TRANSFER_BANK",
					PaidAt:        "2020-01-27",
					Status:        "PAID",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			req.Header.Set("X-Callback-Token", constants.CallbackToken+"salah")
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/callback")

			bookingController := NewBookingController(mockBookingRepository{})
			bookingController.Callback()(context)

			var response common.ResponseError

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Unauthorized", response.Message)
		},
	)
}

func TestBookingByUserID(t *testing.T) {
	t.Run(
		"1. Success Find Booking by userid Test", func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(
				"Bearer",
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQG1haWwuY29tIiwiZXhwIjoxNjQzMzEzOTU3LCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjF9.8j7OV5D2CyfTWS-axnpwRCEXO89eTpu6NQ9NY_V6E24",
			)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/history")

			bookingController := NewBookingController(mockBookingRepository{})
			bookingController.BookingByUserID()(context)

			var response common.ResponseSuccess

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Successful Operation", response.Message)
		},
	)
	t.Run(
		"2. Fail Find Booking By User Id Test", func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(
				"Bearer",
				"leyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIyQG1haWwuY29tIiwiZXhwIjoxNjQzMzE0MDY5LCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjJ9.M3y9yVXPTEK-KDtcn2fcHXCExdszxvj7uWXz27pq3ms",
			)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/history")

			bookingController := NewBookingController(mockFalseBookingRepository{})
			bookingController.BookingByUserID()(context)

			var response common.ResponseError

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, 404, response.Code)
		},
	)
}

func TestBookingByHostID(t *testing.T) {
	t.Run(
		"1. Success Find Booking by hostid Test", func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(
				"Bearer",
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQG1haWwuY29tIiwiZXhwIjoxNjQzMzE4MTIzLCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjF9.NeV66QY1zESc_9zmi2iMTcC0ubWT0bxfLaHSAnOjfr8",
			)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/recap")

			bookingController := NewBookingController(mockBookingRepository{})
			bookingController.FindBookingByHostID()(context)

			var response common.ResponseSuccess

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Successful Operation", response.Message)
		},
	)
	t.Run(
		"2. Fail Find Booking By Host Id Test", func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(
				"Bearer",
				"leyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIyQG1haWwuY29tIiwiZXhwIjoxNjQzMzE0MDY5LCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjJ9.M3y9yVXPTEK-KDtcn2fcHXCExdszxvj7uWXz27pq3ms",
			)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/recap")

			bookingController := NewBookingController(mockFalseBookingRepository{})
			bookingController.FindBookingByHostID()(context)

			var response common.ResponseError

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, 404, response.Code)
		},
	)
	t.Run(
		"2. Fail Find Booking By Host Id Test", func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(
				"Bearer",
				"leyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIyQG1haWwuY29tIiwiZXhwIjoxNjQzMzE0MDY5LCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjJ9.M3y9yVXPTEK-KDtcn2fcHXCExdszxvj7uWXz27pq3ms",
			)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/recap")

			bookingController := NewBookingController(mockFalseBookingRepository{})
			bookingController.FindBookingByHostID()(context)

			var response common.ResponseError

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, 404, response.Code)
		},
	)

}

func TestCheckOut(t *testing.T) {
	t.Run(
		"1. Success Checkout Booking Test", func(t *testing.T) {
			e := echo.New()

			requestBody, _ := json.Marshal(CheckOutRequest{InvoiceID: "1"})

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			req.Header.Set(
				"Bearer",
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQG1haWwuY29tIiwiZXhwIjoxNjQzMzE4MTIzLCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjF9.NeV66QY1zESc_9zmi2iMTcC0ubWT0bxfLaHSAnOjfr8",
			)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/checkout")

			bookingController := NewBookingController(mockBookingRepository{})
			bookingController.CheckOut()(context)

			var response common.ResponseSuccess

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Successful Operation", response.Message)
		},
	)
	t.Run(
		"2. Fail Checkout Booking Test", func(t *testing.T) {
			e := echo.New()

			requestBody, _ := json.Marshal(
				CheckOutRequest{
					InvoiceID: "2",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			req.Header.Set(
				"Bearer",
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQG1haWwuY29tIiwiZXhwIjoxNjQzMzE4MTIzLCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjF9.NeV66QY1zESc_9zmi2iMTcC0ubWT0bxfLaHSAnOjfr8",
			)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/checkout")

			bookingController := NewBookingController(mockFalseBookingRepository{})
			bookingController.CheckOut()(context)

			var response common.ResponseError

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Not Found", response.Message)
		},
	)
}

func TestReschedule(t *testing.T) {
	t.Run(
		"1. Success Reschedule Booking Test", func(t *testing.T) {
			e := echo.New()

			requestBody, _ := json.Marshal(
				ResecheduleRequest{
					InvoiceID: "1",
					CheckIn:   "2021-01-12",
				},
			)

			req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
			req.Header.Set(
				"Bearer",
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQG1haWwuY29tIiwiZXhwIjoxNjQzMzE4MTIzLCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjF9.NeV66QY1zESc_9zmi2iMTcC0ubWT0bxfLaHSAnOjfr8",
			)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/reschedule")

			bookingController := NewBookingController(mockBookingRepository{})
			bookingController.Reschedule()(context)

			var response common.ResponseSuccess

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Successful Operation", response.Message)
		},
	)
	t.Run(
		"2. Fail Reschedule Booking Test", func(t *testing.T) {
			e := echo.New()

			requestBody, _ := json.Marshal(
				ResecheduleRequest{
					InvoiceID: "100",
					CheckIn:   "2021-01-12",
				},
			)

			req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
			req.Header.Set(
				"Bearer",
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQG1haWwuY29tIiwiZXhwIjoxNjQzMzE4MTIzLCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjF9.NeV66QY1zESc_9zmi2iMTcC0ubWT0bxfLaHSAnOjfr8",
			)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/reschedule")

			bookingController := NewBookingController(mockFalseBookingRepository{})
			bookingController.Reschedule()(context)

			var response common.ResponseError

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Not Found", response.Message)
		},
	)
	t.Run(
		"3. Fail Checkin bad request Booking Test", func(t *testing.T) {
			e := echo.New()

			requestBody, _ := json.Marshal(
				ResecheduleRequest{
					InvoiceID: "100",
					CheckIn:   "2021-1-12",
				},
			)

			req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
			req.Header.Set(
				"Bearer",
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQG1haWwuY29tIiwiZXhwIjoxNjQzMzE4MTIzLCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjF9.NeV66QY1zESc_9zmi2iMTcC0ubWT0bxfLaHSAnOjfr8",
			)
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/booking/reschedule")

			bookingController := NewBookingController(mockFalseBookingRepository{})
			bookingController.Reschedule()(context)

			var response common.ResponseError

			json.Unmarshal([]byte(res.Body.Bytes()), &response)

			assert.Equal(t, "Bad Request", response.Message)
		},
	)

}

type mockBookingRepository struct{}

func (mb mockBookingRepository) CreateBooking(booking entities.Booking) (entities.Booking, error) {
	CheckIn, _ := time.Parse("2022-01-28", "2022-01-28")
	CheckOut, _ := time.Parse("2022-01-28", "2022-01-29")
	return entities.Booking{
		ID:            1,
		User_id:       1,
		CheckIn:       CheckIn,
		CheckOut:      CheckOut,
		PaymentStatus: "PENDING",
		InvoiceID:     "1",
		PaymentMethod: "",
		PaymentURL:    "localhost",
		PaidAt:        time.Time{},
		PriceBooking:  100000,
		HomestayID:    1,
	}, nil
}
func (mb mockBookingRepository) FindHomeStay(days, id int) ([]xendit.InvoiceItem, int, error) {

	return []xendit.InvoiceItem{}, 1, nil
}
func (mb mockBookingRepository) FindCustomer(id int) (xendit.InvoiceCustomer, error) {
	return xendit.InvoiceCustomer{}, nil
}
func (mb mockBookingRepository) Update(bookingID string, booking entities.Booking) (entities.Booking, error) {
	return entities.Booking{
		User_id:       1,
		CheckIn:       booking.CheckIn,
		CheckOut:      booking.CheckOut,
		PaymentStatus: "PAID",
		InvoiceID:     "1",
		PaymentMethod: "TRANSFER_BANK",
		PaymentURL:    "localhost",
		PaidAt:        time.Time{},
		PriceBooking:  100000,
		HomestayID:    1,
	}, nil
}
func (mb mockBookingRepository) FindBookingByUserID(id int) ([]entities.Booking, error) {
	return []entities.Booking{
		{
			User_id:       1,
			CheckIn:       time.Time{},
			CheckOut:      time.Time{},
			PaymentStatus: "PAID",
			InvoiceID:     "1",
			PaymentMethod: "TRANSFER_BANK",
			PaymentURL:    "localhost",
			PaidAt:        time.Time{},
			PriceBooking:  100000,
		},
	}, nil
}
func (mb mockBookingRepository) FindBookingByHostID(id int) ([]entities.Booking, error) {
	return []entities.Booking{
		{
			User_id:       1,
			CheckIn:       time.Time{},
			CheckOut:      time.Time{},
			PaymentStatus: "PAID",
			InvoiceID:     "1",
			PaymentMethod: "TRANSFER_BANK",
			PaymentURL:    "localhost",
			PaidAt:        time.Time{},
			PriceBooking:  100000,
		},
	}, nil
}
func (mb mockBookingRepository) Checkout(invoiceID string, userId int) (entities.Homestay, error) {
	return entities.Homestay{

		ID:             1,
		Name:           "Kos Kosan",
		Price:          100000,
		Booking_Status: "Available",
		HostID:         1,
		City_id:         1,
		BookingID:      nil,
	}, nil
}
func (mb mockBookingRepository) Reschedule(userID int, invoiceID string, checkIN time.Time) (
	entities.Booking, error,
) {
	return entities.Booking{
		User_id:       1,
		CheckIn:       time.Time{},
		CheckOut:      time.Time{},
		PaymentStatus: "PAID",
		InvoiceID:     "1",
		PaymentMethod: "TRANSFER_BANK",
		PaymentURL:    "localhost",
		PaidAt:        time.Time{},
		PriceBooking:  100000,
		HomestayID:    1,
	}, nil
}

type mockFalseBookingRepository struct{}

func (mf mockFalseBookingRepository) CreateBooking(booking entities.Booking) (entities.Booking, error) {
	CheckIn, _ := time.Parse("2022-01-28", "2022-01-28")
	CheckOut, _ := time.Parse("2022-01-28", "2022-01-29")
	return entities.Booking{
		ID:            1,
		User_id:       1,
		CheckIn:       CheckIn,
		CheckOut:      CheckOut,
		PaymentStatus: "PENDING",
		InvoiceID:     "1",
		PaymentMethod: "",
		PaymentURL:    "localhost",
		PaidAt:        time.Time{},
		PriceBooking:  100000,
		HomestayID:    1,
	}, errors.New("error")
}
func (mf mockFalseBookingRepository) FindHomeStay(days, id int) ([]xendit.InvoiceItem, int, error) {

	return []xendit.InvoiceItem{}, 1, errors.New("error")
}
func (mf mockFalseBookingRepository) FindCustomer(id int) (xendit.InvoiceCustomer, error) {
	return xendit.InvoiceCustomer{}, nil
}
func (mf mockFalseBookingRepository) Update(bookingID string, booking entities.Booking) (entities.Booking, error) {
	return entities.Booking{
		User_id:       1,
		CheckIn:       booking.CheckIn,
		CheckOut:      booking.CheckOut,
		PaymentStatus: "PAID",
		InvoiceID:     "1",
		PaymentMethod: "TRANSFER_BANK",
		PaymentURL:    "localhost",
		PaidAt:        time.Time{},
		PriceBooking:  100000,
		HomestayID:    1,
	}, errors.New("error")
}
func (mf mockFalseBookingRepository) FindBookingByUserID(id int) ([]entities.Booking, error) {
	return []entities.Booking{
		{
			User_id:       10,
			CheckIn:       time.Time{},
			CheckOut:      time.Time{},
			PaymentStatus: "PAID",
			InvoiceID:     "1",
			PaymentMethod: "TRANSFER_BANK",
			PaymentURL:    "localhost",
			PaidAt:        time.Time{},
			PriceBooking:  100000,
		},
	}, errors.New("error")
}
func (mf mockFalseBookingRepository) FindBookingByHostID(id int) ([]entities.Booking, error) {
	return []entities.Booking{
		{
			User_id:       10,
			CheckIn:       time.Time{},
			CheckOut:      time.Time{},
			PaymentStatus: "PAID",
			InvoiceID:     "1",
			PaymentMethod: "TRANSFER_BANK",
			PaymentURL:    "localhost",
			PaidAt:        time.Time{},
			PriceBooking:  100000,
		},
	}, errors.New("error")
}
func (mf mockFalseBookingRepository) Checkout(invoiceID string, userID int) (entities.Homestay, error) {
	return entities.Homestay{

		ID:             1,
		Name:           "Kos Kosan",
		Price:          100000,
		Booking_Status: "Available",
		HostID:         1,
		City_id:         1,
		BookingID:      nil,
	}, errors.New("error")
}
func (mf mockFalseBookingRepository) Reschedule(userID int, invoiceID string, checkIN time.Time) (
	entities.Booking, error,
) {
	return entities.Booking{
		User_id:       1,
		CheckIn:       time.Time{},
		CheckOut:      time.Time{},
		PaymentStatus: "PAID",
		InvoiceID:     "1",
		PaymentMethod: "TRANSFER_BANK",
		PaymentURL:    "localhost",
		PaidAt:        time.Time{},
		PriceBooking:  100000,
		HomestayID:    1,
	}, errors.New("error")
}
