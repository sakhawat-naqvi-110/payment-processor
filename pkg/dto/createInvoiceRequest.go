package dto

type CreateInvoiceRequest struct {
	MerchantID          uint    `json:"merchant_id" binding:"required"`
	CustomerID          uint    `json:"customer_id" binding:"required"`
	Amount              float64 `json:"amount" binding:"required"`
	Currency            string  `json:"currency" binding:"required"`
	OptionalDescription string  `json:"optional_description,omitempty"`
}
