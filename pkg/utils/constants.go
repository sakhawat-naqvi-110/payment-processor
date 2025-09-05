package utils

const (
	DB_CONNECTION_URL = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
)

// Payment status constants
const (
	PaymentStatusSuccess           = "SUCCESS"
	PaymentStatusInsufficientFunds = "INSUFFICIENT_FUNDS"
	PaymentStatusDoNotHonor        = "DO_NOT_HONOR"
	PaymentStatusDeclined          = "DECLINED"
)
