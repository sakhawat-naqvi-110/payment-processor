package mapper

import (
	"go/payment-processor/pkg/dto"
	"go/payment-processor/pkg/entities"
	"go/payment-processor/pkg/utils"
)

func ToPaymentResponse(payment *entities.Payment) *dto.ProcessPaymentResponse {
	return &dto.ProcessPaymentResponse{
		ID:            payment.ID,
		InvoiceID:     payment.InvoiceID,
		Amount:        payment.Amount,
		PaymentStatus: payment.PaymentStatus,
		PaymentMethod: payment.PaymentMethod,
		PaymentSource: payment.PaymentSource,
	}
}

// ToPaymentEntity maps a CreatePaymentRequest DTO to a Payment entity
func ToPaymentEntity(request *dto.ProcessPaymentRequest) *entities.Payment {
	return &entities.Payment{
		InvoiceID:     request.InvoiceID,
		PaymentMethod: request.PaymentMethod,
		PaymentSource: request.PaymentSource,
	}
}

func ToInvoiceResponse(invoice *entities.Invoice) *dto.InvoiceResponse {
	return &dto.InvoiceResponse{
		ID:         invoice.ID,
		MerchantID: invoice.MerchantID,
		CustomerID: invoice.CustomerID,
		Amount:     invoice.Amount,
		Currency:   invoice.Currency,
	}
}

// ToInvoiceEntity maps a CreateInvoiceRequest DTO to an Invoice entity
func ToInvoiceEntity(request *dto.CreateInvoiceRequest) *entities.Invoice {
	return &entities.Invoice{
		MerchantID:          request.MerchantID,
		CustomerID:          request.CustomerID,
		Amount:              utils.ConvertFloat64ToDecimal(request.Amount),
		Currency:            request.Currency,
		OptionalDescription: request.OptionalDescription,
	}
}
