package payment

import (
	"fmt"
)

// AccountParent refers to a parent struct
type AccountParent struct{}

// AvailableFuncs method to be embedded
func (a *AccountParent) AvailableFuncs() float32 {
	fmt.Println("Displaying available funds ...")
	return 0
}

// ProcessPayment method to be embedded
func (a *AccountParent) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment ...")
	return true
}
