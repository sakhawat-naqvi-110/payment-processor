package http

import (
	"go/payment-processor/pkg/dto"
	"go/payment-processor/pkg/repository"
	services "go/payment-processor/pkg/service"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Group, db *gorm.DB, logger *zap.Logger, validator *validator.Validate) {
	handler := NewHandler(db, logger, validator)
	e.POST("/invoices", handler.CreateInvoice)
	e.GET("/invoices/:id", handler.getInvoice)
	e.POST("/invoices/:id/payments", handler.processPayment)
	e.GET("/invoices/:id/payment-status", handler.getPaymentStatus)
}

type Handler struct {
	log            *zap.Logger
	validator      *validator.Validate
	invoiceService services.InvoiceService
	paymentService services.PaymentService
}

func NewHandler(db *gorm.DB, logger *zap.Logger, validate *validator.Validate) *Handler {
	repo := repository.NewRepository(db, logger)
	invoiceService := services.NewInvoiceService(logger, repo, validate)
	paymentService := services.NewPaymentService(logger, repo, validate)

	return &Handler{log: logger, validator: validate, invoiceService: invoiceService, paymentService: paymentService}
}

func (h *Handler) CreateInvoice(c echo.Context) error {
	var req dto.CreateInvoiceRequest
	if err := c.Bind(&req); err != nil {
		h.log.Error("Invalid request payload", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.validator.Struct(req); err != nil {
		h.log.Error("Validation failed", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Create invoice
	invoice, err := h.invoiceService.CreateInvoice(&req)
	if err != nil {
		h.log.Error("Failed to create invoice", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create invoice"})
	}

	return c.JSON(http.StatusCreated, invoice)
}

func (h *Handler) getInvoice(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Invalid invoice ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid invoice ID"})
	}

	invoice, err := h.invoiceService.GetInvoiceByID(uint(id))
	if err != nil {
		h.log.Error("Invoice not found", zap.Error(err))
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Invoice not found"})
	}

	return c.JSON(http.StatusOK, invoice)
}

func (h *Handler) processPayment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Invalid invoice ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid invoice ID"})
	}

	var req dto.ProcessPaymentRequest
	if err := c.Bind(&req); err != nil {
		h.log.Error("Invalid request payload", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Validate request
	req.InvoiceID = uint(id)
	if err := h.validator.Struct(req); err != nil {
		h.log.Error("Validation failed", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Process payment
	payment, err := h.paymentService.ProcessPayment(&req)
	if err != nil {
		h.log.Error("Failed to process payment", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to process payment"})
	}

	return c.JSON(http.StatusOK, payment)
}

func (h *Handler) getPaymentStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("Invalid invoice ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid invoice ID"})
	}

	status, err := h.paymentService.GetPaymentStatus(uint(id))
	if err != nil {
		h.log.Error("Payment status not found", zap.Error(err))
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Payment status not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"payment_status": status})
}
