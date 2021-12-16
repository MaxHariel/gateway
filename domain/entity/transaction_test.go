package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransactionIsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = uuid.NewString()
	transaction.AccountID = uuid.NewString()
	transaction.Amount = 950

	assert.Nil(t, transaction.IsValid())
}

func TestTransactionAmountIsGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = uuid.NewString()
	transaction.AccountID = uuid.NewString()
	transaction.Amount = 1010

	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "you can't do transcation with value greater than 1000.0", err.Error())
}

func TestTransactionAmountIsLessThan0(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = uuid.NewString()
	transaction.AccountID = uuid.NewString()
	transaction.Amount = -1

	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "you can't do transcation with value less than 0.0", err.Error())
}
