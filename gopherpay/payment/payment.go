package payment

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// Option interface refers to payment options
type Option interface {
	ProcessPayment(float32) bool
}

// CreditCard data structure
type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

// CreateCreditCardAccount constructor
func CreateCreditCardAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

// ProcessPayment for credit card payment
func (c CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Procesing a credit card payment ...")
	return true
}

// OwnerName getter
func (c CreditCard) OwnerName() string {
	return c.ownerName
}

// SetOwnerName setter
func (c CreditCard) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid name provided")
	}
	c.ownerName = value
	return nil
}

// CardNumber getter
func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

// SetCardNumber setter
func (c CreditCard) SetCardNumber(value string) error {
	var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")
	if !cardNumberPattern.Match([]byte(value)) {
		return errors.New("Invalid credit card number format")
	}
	c.cardNumber = value
	return nil
}

// ExpirationDate getter
func (c CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

// SetExpirationDate setter
func (c CreditCard) SetExpirationDate(month, year int) error {
	now := time.Now()
	if year < now.Year() || (year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Expiration day must lie in the future")
	}
	c.expirationMonth, c.expirationYear = month, year
	return nil
}

// SecurityCode getter
func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

// SetSecurityCode setter
func (c CreditCard) SetSecurityCode(value int) error {
	if value < 100 || value > 999 {
		return errors.New("Invalid security code must have 3 digits")
	}
	c.securityCode = value
	return nil
}

// AvailableCredit retrives the available amount
func (c CreditCard) AvailableCredit() float32 {
	return 5000 // this can come from a web service, client doesn't know or care about
}
