package utils

import (
	"github.com/shopspring/decimal"
	"math"
	"strings"
)

func ConvertFloat64ToDecimal(value float64) decimal.Decimal {
	if math.IsNaN(value) {
		return decimal.NewFromInt(0)
	}
	if math.IsInf(value, 0) {
		return decimal.NewFromInt(0)
	}
	decimalValue := decimal.NewFromFloat(value)

	return decimalValue
}

func IsCurrencyAllowed(currency string, allowedCurrencies []string) bool {
	for _, allowedCurrency := range allowedCurrencies {
		if strings.EqualFold(currency, allowedCurrency) {
			return true
		}
	}
	return false
}
