package process_transaction

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/maxhariel/gateway/domain/entity"
	mock_repository "github.com/maxhariel/gateway/domain/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTranscation_ExecuteInvalidCreditCard(t *testing.T) {

	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "45454545454545",
		CreditCardExpirationMonth: 10,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             234,
		Amount:                    300.0,
	}

	expectOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECT,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectOutput.Status, expectOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Excute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectOutput, output)
}

func TestProcessTranscation_ExecuteRejectTransaction(t *testing.T) {

	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "5296212435100121",
		CreditCardExpirationMonth: 10,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             234,
		Amount:                    1010,
	}

	expectOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECT,
		ErrorMessage: "you can't do transcation with value greater than 1000.0",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectOutput.Status, expectOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Excute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectOutput, output)
}

func TestProcessTranscation_ExecuteApprovedTransaction(t *testing.T) {

	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "5296212435100121",
		CreditCardExpirationMonth: 10,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             234,
		Amount:                    1000,
	}

	expectOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectOutput.Status, expectOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Excute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectOutput, output)
}
