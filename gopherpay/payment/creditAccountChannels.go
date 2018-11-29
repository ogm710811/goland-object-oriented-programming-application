package payment

import (
	"fmt"
)

// CreditAccount data structure
type CreditAccount struct {
	ownerName       string
	availableCredit float32
}

// ProcessPayment for credit account payment
func (c *CreditAccount) processPayment(amount float32) {
	fmt.Println("Processing credit account payment ...")
}

// CreateCreditAccount constructor
func CreateCreditAccount(ownerName string, chargeCh chan float32) *CreditAccount {
	creditAccount := &CreditAccount{
		ownerName: ownerName,
	}
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeCh)

	return creditAccount
}
