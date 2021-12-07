package entity

import "errors"

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

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("you can't do transcation with value greater than 1000")
	}
	return nil
}
