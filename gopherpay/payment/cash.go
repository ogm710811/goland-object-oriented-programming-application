package payment

import (
	"fmt"
)

// Cash data structure
type Cash struct{}

// CreateCashAccount constructor
func CreateCashAccount() *Cash {
	return &Cash{}
}

// ProcessPayment for cash payment
func (c Cash) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a cash transaction ...")
	return true
}
