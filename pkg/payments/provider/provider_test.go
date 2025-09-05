package provider

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPayRequestSuccess(t *testing.T) {
	provider := PaymentProvider{
		byIDs:          make(map[uuid.UUID]PaymentStatus),
		byReferenceIDs: make(map[uuid.UUID]PaymentStatus),
	}

	payment, err := provider.Pay(context.Background(), PaymentDetails{
		CardNumber:        "4242424242424242",
		CardHolder:        "John Smith",
		Expiry:            "01/25",
		CVC:               "123",
		Amount:            100.00,
		CurrencyCode:      "USD",
		BankAccountNumber: "123-123-123-123",
	})

	assert.Nil(t, err)
	assert.Equal(t, PaymentStatusSuccess, payment.Status)
}

func TestPayRequestInsufficientFunds(t *testing.T) {
	provider := PaymentProvider{
		byIDs:          make(map[uuid.UUID]PaymentStatus),
		byReferenceIDs: make(map[uuid.UUID]PaymentStatus),
	}

	payment, err := provider.Pay(context.Background(), PaymentDetails{
		CardNumber:        "4242424242421212",
		CardHolder:        "John Smith",
		Expiry:            "01/25",
		CVC:               "123",
		Amount:            100.00,
		CurrencyCode:      "USD",
		BankAccountNumber: "123-123-123-123",
	})

	assert.Nil(t, err)
	assert.Equal(t, PaymentStatusInsufficientFunds, payment.Status)
}

func TestPayRequestDoNotHonor(t *testing.T) {
	provider := PaymentProvider{
		byIDs:          make(map[uuid.UUID]PaymentStatus),
		byReferenceIDs: make(map[uuid.UUID]PaymentStatus),
	}

	payment, err := provider.Pay(context.Background(), PaymentDetails{
		CardNumber:        "4242424242422323",
		CardHolder:        "John Smith",
		Expiry:            "01/25",
		CVC:               "123",
		Amount:            100.00,
		CurrencyCode:      "USD",
		BankAccountNumber: "123-123-123-123",
	})

	assert.Nil(t, err)
	assert.Equal(t, PaymentStatusDoNotHonor, payment.Status)
}

func TestPayRequestDeclined(t *testing.T) {
	provider := PaymentProvider{
		byIDs:          make(map[uuid.UUID]PaymentStatus),
		byReferenceIDs: make(map[uuid.UUID]PaymentStatus),
	}

	payment, err := provider.Pay(context.Background(), PaymentDetails{
		CardNumber:        "4242424242423434",
		CardHolder:        "John Smith",
		Expiry:            "01/25",
		CVC:               "123",
		Amount:            100.00,
		CurrencyCode:      "USD",
		BankAccountNumber: "123-123-123-123",
	})

	assert.Nil(t, err)
	assert.Equal(t, PaymentStatusDeclined, payment.Status)
}

func TestByID(t *testing.T) {
	uuid1, _ := uuid.NewV7()
	uuid2, _ := uuid.NewV7()
	uuid3, _ := uuid.NewV7()
	uuid4, _ := uuid.NewV7()

	provider := PaymentProvider{
		byIDs: map[uuid.UUID]PaymentStatus{
			uuid1: PaymentStatusSuccess,
			uuid2: PaymentStatusInsufficientFunds,
			uuid3: PaymentStatusDoNotHonor,
			uuid4: PaymentStatusDeclined,
		},
	}

	status, ok := provider.ByID(uuid1)
	assert.True(t, ok)
	assert.Equal(t, PaymentStatusSuccess, status)

	status, ok = provider.ByID(uuid2)
	assert.True(t, ok)
	assert.Equal(t, PaymentStatusInsufficientFunds, status)

	status, ok = provider.ByID(uuid3)
	assert.True(t, ok)
	assert.Equal(t, PaymentStatusDoNotHonor, status)

	status, ok = provider.ByID(uuid4)
	assert.True(t, ok)
	assert.Equal(t, PaymentStatusDeclined, status)
}

func TestByReferenceID(t *testing.T) {
	uuid1, _ := uuid.NewV7()
	uuid2, _ := uuid.NewV7()
	uuid3, _ := uuid.NewV7()
	uuid4, _ := uuid.NewV7()

	provider := PaymentProvider{
		byReferenceIDs: map[uuid.UUID]PaymentStatus{
			uuid1: PaymentStatusSuccess,
			uuid2: PaymentStatusInsufficientFunds,
			uuid3: PaymentStatusDoNotHonor,
			uuid4: PaymentStatusDeclined,
		},
	}

	status, ok := provider.ByReferenceID(uuid1)
	assert.True(t, ok)
	assert.Equal(t, PaymentStatusSuccess, status)

	status, ok = provider.ByReferenceID(uuid2)
	assert.True(t, ok)
	assert.Equal(t, PaymentStatusInsufficientFunds, status)

	status, ok = provider.ByReferenceID(uuid3)
	assert.True(t, ok)
	assert.Equal(t, PaymentStatusDoNotHonor, status)

	status, ok = provider.ByReferenceID(uuid4)
	assert.True(t, ok)
	assert.Equal(t, PaymentStatusDeclined, status)
}
