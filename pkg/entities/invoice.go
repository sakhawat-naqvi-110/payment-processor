package entities

import (
	"github.com/shopspring/decimal"
	"time"
)

type Invoice struct {
	AuditTrail
	MerchantID          uint            `gorm:"column:merchant_id" json:"merchant_id"`
	CustomerID          uint            `gorm:"column:customer_id" json:"customer_id"`
	Amount              decimal.Decimal `gorm:"column:amount" json:"amount"`
	Currency            string          `gorm:"column:currency" json:"currency"`
	OptionalDescription string          `gorm:"column:optional_description" json:"optional_description,omitempty"`
}

func (Invoice) TableName() string {
	return "invoice"
}

type AuditTrail struct {
	ID            uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt     *time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	CreatedBy     string     `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	LastUpdatedAt *time.Time `gorm:"column:last_updated_at;autoUpdateTime" json:"last_updated_at"`
	LastUpdatedBy string     `gorm:"column:last_updated_by;type:varchar(255)" json:"last_updated_by"`
	IsActive      bool       `gorm:"column:is_active" json:"is_active"` // Indicates whether the configuration is active
}
