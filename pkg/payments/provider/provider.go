package provider

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
)

type PaymentProvider struct {
	byIDs          map[uuid.UUID]PaymentStatus
	byReferenceIDs map[uuid.UUID]PaymentStatus
}

type Payment struct {
	ID     uuid.UUID
	Status PaymentStatus
}

type PaymentDetails struct {
	ReferenceID       uuid.UUID
	CardNumber        string
	CardHolder        string
	Expiry            string
	CVC               string
	Amount            float64
	CurrencyCode      string
	BankAccountNumber string
}

type PaymentStatus int

const (
	PaymentStatusSuccess PaymentStatus = iota + 1
	PaymentStatusInsufficientFunds
	PaymentStatusDoNotHonor
	PaymentStatusDeclined
)

func New() PaymentProvider {
	return PaymentProvider{
		byIDs:          make(map[uuid.UUID]PaymentStatus),
		byReferenceIDs: make(map[uuid.UUID]PaymentStatus),
	}
}

func (p PaymentProvider) Pay(ctx context.Context, details PaymentDetails) (Payment, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return Payment{}, err
	}

	status := PaymentStatusSuccess
	if strings.HasSuffix(details.CardNumber, "1212") {
		status = PaymentStatusInsufficientFunds
	} else if strings.HasSuffix(details.CardNumber, "2323") {
		status = PaymentStatusDoNotHonor
	} else if strings.HasSuffix(details.CardNumber, "3434") {
		status = PaymentStatusDeclined
	} else if strings.HasSuffix(details.CardNumber, "4545") {
		time.Sleep(10000 * time.Hour)
	}

	p.byIDs[id] = status
	p.byReferenceIDs[id] = status

	return Payment{ID: id, Status: status}, nil
}

func (p PaymentProvider) ByID(id uuid.UUID) (PaymentStatus, bool) {
	if status, ok := p.byIDs[id]; ok {
		return status, true
	}
	return 0, false
}

func (p PaymentProvider) ByReferenceID(id uuid.UUID) (PaymentStatus, bool) {
	if status, ok := p.byReferenceIDs[id]; ok {
		return status, true
	}
	return 0, false
}
