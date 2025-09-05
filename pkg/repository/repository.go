package repository

import (
	"errors"
	"github.com/labstack/gommon/log"
	"go/payment-processor/pkg/entities"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	CreateInvoice(invoice *entities.Invoice) (*entities.Invoice, error)
	GetInvoiceByID(id uint) (*entities.Invoice, error)
	ProcessPayment(payment *entities.Payment) (*entities.Payment, error)
	GetPaymentStatus(invoiceID uint) (string, error)
	DoesMerchantExist(merchantID uint) (*entities.Merchant, error)
	DoesCustomerExist(customerID uint) (*entities.Customer, error)
	GetAllowedCurrenciesForMerchant(merchantID uint) (string, error)
	DoesInvoiceExist(invoiceID uint) (*entities.Invoice, error)
}

type repository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewRepository(db *gorm.DB, logger *zap.Logger) Repository {
	return &repository{
		db:  db,
		log: logger}
}

func (r *repository) CreateInvoice(invoice *entities.Invoice) (*entities.Invoice, error) {
	if err := r.db.Create(&invoice).Error; err != nil {
		return nil, err
	}
	return invoice, nil
}

func (r *repository) GetInvoiceByID(id uint) (*entities.Invoice, error) {
	var invoice entities.Invoice
	if err := r.db.First(&invoice, id).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *repository) ProcessPayment(payment *entities.Payment) (*entities.Payment, error) {
	// Example logic for processing a payment
	if err := r.db.Create(&payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *repository) GetPaymentStatus(invoiceID uint) (string, error) {
	var payment entities.Payment
	if err := r.db.Where("invoice_id = ?", invoiceID).First(&payment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("payment not found for this invoice")
		}
		return "", err
	}
	return payment.PaymentStatus, nil
}

func (r *repository) DoesMerchantExist(merchantID uint) (*entities.Merchant, error) {
	var merchant entities.Merchant
	result := r.db.Where("id = ?", merchantID).First(&merchant)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Error("No record found in merchant", zap.Error(result.Error))
		}
		return nil, result.Error
	}
	return &merchant, nil
}

func (r *repository) DoesCustomerExist(customerID uint) (*entities.Customer, error) {
	var customer entities.Customer
	result := r.db.Where("id = ?", customerID).First(&customer)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Error("No record found in customer", zap.Error(result.Error))
		}
		return nil, result.Error
	}
	return &customer, nil
}

func (r *repository) GetAllowedCurrenciesForMerchant(merchantID uint) (string, error) {
	var currency string
	err := r.db.Model(&entities.Merchant{}).
		Select("allowed_currency").
		Where("id = ?", merchantID).
		Scan(&currency).Error

	return currency, err
}

func (r *repository) DoesInvoiceExist(invoiceID uint) (*entities.Invoice, error) {
	var invoice *entities.Invoice
	result := r.db.
		Where("id = ?", invoiceID).Find(&invoice)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Error("No record found in invoice", zap.Error(result.Error))
		}
		return nil, result.Error
	}
	return invoice, nil
}
