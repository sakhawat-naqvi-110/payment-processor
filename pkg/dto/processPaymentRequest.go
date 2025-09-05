package dto

type ProcessPaymentRequest struct {
	InvoiceID     uint   `json:"invoice_id" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
	PaymentSource string `json:"payment_source" binding:"required"`
}
