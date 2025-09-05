package dto

import "github.com/shopspring/decimal"

type ProcessPaymentResponse struct {
	ID            uint            `json:"id"`
	InvoiceID     uint            `json:"invoice_id"`
	Amount        decimal.Decimal `json:"amount"`
	PaymentStatus string          `json:"payment_status"`
	PaymentMethod string          `json:"payment_method"`
	PaymentSource string          `json:"payment_source"`
}
