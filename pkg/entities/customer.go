package entities

type Customer struct {
	AuditTrail
	CustomerName    string `gorm:"column:customer_name" json:"customer_name"`
	CustomerEmail   string `gorm:"column:customer_email" json:"customer_email"`
	CustomerAddress string `gorm:"column:customer_address" json:"customer_address"`
}

func (Customer) TableName() string {
	return "customer"
}
