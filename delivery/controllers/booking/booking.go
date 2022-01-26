package booking

import (
	"air-bnb/constants"
	"air-bnb/delivery/common"
	"air-bnb/delivery/middlewares"
	"air-bnb/entities"
	"air-bnb/preference"
	"air-bnb/repository/booking"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"math/rand"
	"net/http"
	"time"
)

type BookingController struct {
	bookRepo booking.Booking
}

func NewBookingController(bookingRepo booking.Booking) *BookingController {
	return &BookingController{bookingRepo}
}

func (bc BookingController) Create(c echo.Context) error {
	var bookingRequest BookingRequest
	xendit.Opt.SecretKey = constants.XendToken

	// bind request data
	if err := c.Bind(&bookingRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	//get data from token
	userID := middlewares.NewAuth().ExtractTokenUserID(c)
	userEmail := middlewares.NewAuth().ExtractTokenEmail(c)

	//calculate how many days user stay
	checkIn, err := time.Parse("2006-01-02", bookingRequest.CheckIn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	checkOut, err := time.Parse("2006-01-02", bookingRequest.CheckOut)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	calcDays := checkOut.Sub(checkIn).Hours() / 24

	// data to create invoice,
	ExternalID := fmt.Sprint("Invoice-", userID, "-", bookingRequest.HomeStayID, bookingRequest.CheckIn, rand.Intn(100))
	var item, price, errBooking = bc.bookRepo.FindHomeStay(int(calcDays), int(bookingRequest.HomeStayID))

	if errBooking != nil {
		return c.JSON(http.StatusNotFound, "Homestay tidak ditemukan, atau booking status not available")
	}

	customer, _ := bc.bookRepo.FindCustomer(userID)

	data := invoice.CreateParams{
		ExternalID:                     ExternalID,
		Amount:                         calcDays * float64(price),
		PayerEmail:                     userEmail,
		Description:                    fmt.Sprint("Invoice sewa homestay Selama ", calcDays, " hari"),
		CustomerNotificationPreference: preference.SendNotifWith,
		ShouldSendEmail:                &preference.SendEmail,
		Items:                          item,
		Customer:                       customer,
	}

	resp, err := invoice.Create(&data)

	if err != nil {
		fmt.Println(err)
	}

	//save to db
	BookingData := entities.Booking{
		User_id:       uint(userID),
		InvoiceID:     ExternalID,
		HomestayID:    bookingRequest.HomeStayID,
		CheckIn:       checkIn,
		CheckOut:      checkOut,
		PaymentStatus: resp.Status,
		PaymentMethod: resp.PaymentMethod,
		PaymentURL:    resp.InvoiceURL,
		PriceBooking:  int(calcDays * float64(price)),
	}

	bc.bookRepo.CreateBooking(BookingData)

	return c.JSON(http.StatusOK, common.SuccessResponse(BookingData))
}

func (bc BookingController) Callback(c echo.Context) error {

	req := c.Request()
	headers := req.Header

	callBackToken := headers.Get("X-Callback-Token")

	if callBackToken != constants.CallbackToken {
		return c.JSON(http.StatusUnauthorized, common.NewUnauthorized())
	}

	var callBackRequest CallbackRequest
	err := c.Bind(&callBackRequest)
	if err != nil {
		return err
	}

	parsePaid, _ := time.Parse(time.RFC3339, callBackRequest.PaidAt)
	callBackData := entities.Booking{
		PaymentStatus: callBackRequest.Status,
		PaymentMethod: callBackRequest.PaymentMethod,
		PaidAt:        parsePaid,
	}

	callBack, _ := bc.bookRepo.Update(callBackRequest.ExternalID, callBackData)

	return c.JSON(http.StatusOK, callBack)
}

func (bc BookingController) BookingByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.NewAuth().ExtractTokenUserID(c)

		res, _ := bc.bookRepo.FindBookingByUserID(userID)

		if len(res) == 0 {
			c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (bc BookingController) FindBookingByHostID() echo.HandlerFunc {
	return func(c echo.Context) error {
		hostID := middlewares.NewAuth().ExtractTokenUserID(c)

		res, _ := bc.bookRepo.FindBookingByUserID(hostID)

		if len(res) == 0 {
			c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (bc BookingController) CheckOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		var invoiceID CheckOutRequest
		c.Bind(&invoiceID)

		hostID := middlewares.NewAuth().ExtractTokenUserID(c)

		_, err := bc.bookRepo.Checkout(invoiceID.InvoiceID, hostID)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}
