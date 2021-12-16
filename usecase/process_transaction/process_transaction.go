package process_transaction

import (
	"github.com/maxhariel/gateway/domain/entity"
	"github.com/maxhariel/gateway/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Excute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	cc, invalidCC := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)

	if invalidCC != nil {
		return p.insertTransaction(transaction, entity.REJECT, invalidCC.Error())
	}

	transaction.SetCreditCard(*cc)
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		return p.insertTransaction(transaction, entity.REJECT, invalidTransaction.Error())
	}
	return p.insertTransaction(transaction, entity.APPROVED, "")
}

func (p *ProcessTransaction) insertTransaction(transaction *entity.Transaction, status string, errorMesage string) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, status, errorMesage)
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       status,
		ErrorMessage: errorMesage,
	}

	return output, nil
}
