package booking

type BookingRequest struct {
	HomeStayID uint   `json:"home_stay_id"`
	CheckIn    string `json:"check_in"`
	CheckOut   string `json:"check_out"`
}

type CallBack struct {
	Status string
}

type CallbackRequest struct {
	ExternalID    string `json:"external_id"`
	PaymentMethod string `json:"payment_method"`
	PaidAt        string `json:"paid_at"`
	Status        string `json:"status"`
}

type CheckOutRequest struct {
	InvoiceID string `json:"invoice_id"`
}

type ResecheduleRequest struct {
	InvoiceID string `json:"invoiceID"`
	CheckIn   string `json:"check_in"`
}
