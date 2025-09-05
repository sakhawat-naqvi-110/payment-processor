package dto

import "github.com/shopspring/decimal"

type InvoiceResponse struct {
	ID         uint            `json:"id"`
	MerchantID uint            `json:"merchant_id"`
	CustomerID uint            `json:"customer_id"`
	Amount     decimal.Decimal `json:"amount"`
	Currency   string          `json:"currency"`
}
