package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	futureYear = time.Now().AddDate(4, 0, 0).Year()
	lastYear   = time.Now().AddDate(-1, 0, 0).Year()
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("45454545454545", "Mad Max", 11, futureYear, 456)
	assert.Error(t, err)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5296212435100121", "Mad Max", 11, futureYear, 456)
	assert.Nil(t, err)

}

func TestCreditCardMonth(t *testing.T) {
	_, err := NewCreditCard("5296212435100121", "Mad Max", 13, futureYear, 456)
	assert.Error(t, err)
	assert.Equal(t, "invalid credit card expiration month", err.Error())

	_, err = NewCreditCard("5296212435100121", "Mad Max", 0, futureYear, 456)
	assert.Error(t, err)
	assert.Equal(t, "invalid credit card expiration month", err.Error())

	_, err = NewCreditCard("5296212435100121", "Mad Max", 11, futureYear, 456)
	assert.Nil(t, err)
}

func TestCreditCardYear(t *testing.T) {

	_, err := NewCreditCard("5296212435100121", "Mad Max", 11, lastYear, 456)
	assert.Error(t, err)
	assert.Equal(t, "invalid credit card expiration year", err.Error())

	_, err = NewCreditCard("5296212435100121", "Mad Max", 11, 2026, 456)
	assert.Nil(t, err)

}

func TestCreditCardCVV(t *testing.T) {
	_, err := NewCreditCard("5296212435100121", "Mad Max", 11, futureYear, 4560)
	assert.Error(t, err)
	assert.Equal(t, "invalid credit card expiration CVV", err.Error())

	_, err = NewCreditCard("5296212435100121", "Mad Max", 11, futureYear, 60)
	assert.Error(t, err)
	assert.Equal(t, "invalid credit card expiration CVV", err.Error())

	_, err = NewCreditCard("5296212435100121", "Mad Max", 11, 2026, 456)
	assert.Nil(t, err)

}
