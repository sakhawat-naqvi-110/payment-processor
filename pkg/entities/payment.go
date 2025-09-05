package entities

import "github.com/shopspring/decimal"

type Payment struct {
	AuditTrail
	InvoiceID     uint            `gorm:"column:invoice_id" json:"invoice_id"`
	MerchantID    uint            `gorm:"column:merchant_id" json:"merchant_id"`
	CustomerID    uint            `gorm:"column:customer_id" json:"customer_id"`
	Amount        decimal.Decimal `gorm:"column:amount" json:"amount"`
	PaymentStatus string          `gorm:"column:payment_status" json:"payment_status"`
	PaymentMethod string          `gorm:"column:payment_method" json:"payment_method"`
	PaymentSource string          `gorm:"column:payment_source" json:"payment_source"`
}

func (Payment) TableName() string {
	return "payment"
}
