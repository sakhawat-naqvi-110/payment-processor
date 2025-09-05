package entities

type Merchant struct {
	AuditTrail
	MerchantName    string `gorm:"column:merchant_name" json:"merchant_name"`
	MerchantCode    string `gorm:"column:merchant_code" json:"merchant_code"`
	AllowedCurrency string `gorm:"column:allowed_currency" json:"allowed_currency"`
}

func (Merchant) TableName() string {
	return "merchant"
}
