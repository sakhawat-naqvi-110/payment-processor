package services

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"go/payment-processor/pkg/dto"

	"go/payment-processor/pkg/mapper"
	"go/payment-processor/pkg/repository"
	"go/payment-processor/pkg/utils"
)

type InvoiceService interface {
	CreateInvoice(invoice *dto.CreateInvoiceRequest) (*dto.InvoiceResponse, error)
	GetInvoiceByID(id uint) (*dto.InvoiceResponse, error)
	ValidateInvoiceRequest(invoiceRequest *dto.CreateInvoiceRequest) error
}
type invoiceService struct {
	log       *zap.Logger
	repo      repository.Repository
	validator *validator.Validate
}

func NewInvoiceService(log *zap.Logger, repo repository.Repository, validator *validator.Validate) InvoiceService {
	return &invoiceService{log: log,
		repo: repo, validator: validator}
}

func (is *invoiceService) CreateInvoice(invoiceRequest *dto.CreateInvoiceRequest) (*dto.InvoiceResponse, error) {
	is.log.Info("Attempting to create a new invoice", zap.Any("invoice", invoiceRequest))

	// Validate inputs
	if invoiceRequest.MerchantID == 0 || invoiceRequest.CustomerID == 0 ||
		utils.ConvertFloat64ToDecimal(invoiceRequest.Amount).LessThanOrEqual(decimal.NewFromInt(0)) {
		err := errors.New("invalid invoice data")
		is.log.Error("Invalid invoice data", zap.Error(err))
		return nil, err
	}

	err := is.ValidateInvoiceRequest(invoiceRequest)
	if err != nil {
		is.log.Error("Validation failed for create invoice", zap.Error(err))
		return nil, err
	}
	invoice := mapper.ToInvoiceEntity(invoiceRequest)

	// Call repository to create the invoice
	createdInvoice, err := is.repo.CreateInvoice(invoice)
	if err != nil {
		is.log.Error("Failed to create invoice", zap.Error(err))
		return nil, err
	}

	is.log.Info("Invoice created successfully", zap.Uint("invoice_id", createdInvoice.ID))
	return mapper.ToInvoiceResponse(createdInvoice), nil
}

func (is *invoiceService) GetInvoiceByID(id uint) (*dto.InvoiceResponse, error) {
	is.log.Info("Fetching invoice by ID", zap.Uint("invoice_id", id))

	invoice, err := is.repo.GetInvoiceByID(id)
	if err != nil {
		is.log.Error("Invoice not found", zap.Uint("invoice_id", id), zap.Error(err))
		return nil, err
	}

	is.log.Info("Successfully fetched invoice", zap.Uint("invoice_id", id))
	return mapper.ToInvoiceResponse(invoice), nil
}

func (is *invoiceService) ValidateInvoiceRequest(invoiceRequest *dto.CreateInvoiceRequest) error {
	// Validate required fields
	if invoiceRequest.MerchantID == 0 || invoiceRequest.CustomerID == 0 {
		return errors.New("merchant ID and customer ID must be provided")
	}
	if utils.ConvertFloat64ToDecimal(invoiceRequest.Amount).LessThanOrEqual(decimal.NewFromInt(0)) {
		return errors.New("invoice amount must be greater than zero")
	}

	// Check if merchant exists
	merchantExists, err := is.repo.DoesMerchantExist(invoiceRequest.MerchantID)
	if err != nil {
		is.log.Error("Error checking merchant existence", zap.Error(err))
		return errors.New("internal error while validating merchant ID")
	}
	is.log.Info("Merchant Found against Merchant ID: ", zap.Any("Customer ID", merchantExists.ID))

	// Check if customer exists
	customerExists, err := is.repo.DoesCustomerExist(invoiceRequest.CustomerID)
	if err != nil {
		is.log.Error("Error checking customer existence", zap.Error(err))
		return errors.New("internal error while validating customer ID")
	}

	is.log.Info("Customer Found against Customer ID: ", zap.Any("Customer ID", customerExists.ID))

	allowedCurrencies, err := is.repo.GetAllowedCurrenciesForMerchant(invoiceRequest.MerchantID)
	if err != nil {
		is.log.Error("Error fetching allowed currencies for merchant", zap.Error(err))
		return errors.New("internal error while validating currency")
	}

	if invoiceRequest.Currency != allowedCurrencies {
		is.log.Warn("Currency is not allowed for the merchant",
			zap.Uint("merchant_id", invoiceRequest.MerchantID),
			zap.String("currency", invoiceRequest.Currency),
		)
		return errors.New("currency is not allowed for this merchant")
	}

	return nil
}
