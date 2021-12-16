package entity

import "errors"

const (
	REJECT   = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	ID           string
	Amount       float64
	AccountID    string
	Status       string
	CreditCard   CreditCard
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) SetCreditCard(c CreditCard) {
	t.CreditCard = c
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000.0 {
		return errors.New("you can't do transcation with value greater than 1000.0")
	}

	if t.Amount < 0.0 {
		return errors.New("you can't do transcation with value less than 0.0")
	}
	return nil
}
