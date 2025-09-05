package services

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go/payment-processor/pkg/dto"
	"go/payment-processor/pkg/entities"
	"go/payment-processor/pkg/mapper"
	"go/payment-processor/pkg/repository"
	"go/payment-processor/pkg/utils"
	"strings"
	"time"
)

type PaymentService interface {
	ProcessPayment(paymentRequest *dto.ProcessPaymentRequest) (*entities.Payment, error)
	GetPaymentStatus(invoiceID uint) (string, error)
	GetPaymentStatusBasedOnPaymentSource(paymentSource string) string
}

type paymentService struct {
	log       *zap.Logger
	repo      repository.Repository
	validator *validator.Validate
}

func NewPaymentService(log *zap.Logger, repo repository.Repository, validator *validator.Validate) PaymentService {
	return &paymentService{log: log,
		repo: repo, validator: validator}
}

// ProcessPayment - Business logic for processing payments
func (s *paymentService) ProcessPayment(paymentRequest *dto.ProcessPaymentRequest) (*entities.Payment, error) {
	s.log.Info("Processing payment", zap.Any("payment", paymentRequest))

	if paymentRequest.InvoiceID == 0 {
		err := errors.New("invalid payment data")
		s.log.Error("Invalid payment data", zap.Error(err))
		return nil, err
	}

	invoice, err := s.repo.DoesInvoiceExist(paymentRequest.InvoiceID)
	if err != nil {
		s.log.Error("Error checking invoice existence", zap.Error(err))
		return nil, errors.New("internal error while validating invoice ID")
	}

	//TO DO: Implement Encryption/Decryption logic for Payment Source (BAN, Card Number)
	payment := mapper.ToPaymentEntity(paymentRequest)
	payment.PaymentStatus = s.GetPaymentStatusBasedOnPaymentSource(paymentRequest.PaymentSource)
	payment.Amount = invoice.Amount
	payment.InvoiceID = invoice.ID
	payment.CustomerID = invoice.CustomerID
	payment.MerchantID = invoice.MerchantID
	processedPayment, err := s.repo.ProcessPayment(payment)
	if err != nil {
		s.log.Error("Failed to process payment", zap.Error(err))
		return nil, err
	}

	s.log.Info("Payment processed successfully", zap.Uint("payment_id", processedPayment.ID))
	return processedPayment, nil
}

func (s *paymentService) GetPaymentStatus(invoiceID uint) (string, error) {
	s.log.Info("Fetching payment status", zap.Uint("invoice_id", invoiceID))

	status, err := s.repo.GetPaymentStatus(invoiceID)
	if err != nil {
		s.log.Error("Failed to fetch payment status", zap.Uint("invoice_id", invoiceID), zap.Error(err))
		return "", err
	}

	s.log.Info("Successfully fetched payment status", zap.Uint("invoice_id", invoiceID), zap.String("status", status))
	return status, nil
}

func (s *paymentService) GetPaymentStatusBasedOnPaymentSource(paymentSource string) string {
	if strings.HasSuffix(paymentSource, "1212") {
		return utils.PaymentStatusInsufficientFunds
	} else if strings.HasSuffix(paymentSource, "2323") {
		return utils.PaymentStatusDoNotHonor
	} else if strings.HasSuffix(paymentSource, "3434") {
		return utils.PaymentStatusDeclined
	} else if strings.HasSuffix(paymentSource, "4545") {
		time.Sleep(10000 * time.Hour) // Simulating an infinite delay
	}
	return utils.PaymentStatusSuccess
}
