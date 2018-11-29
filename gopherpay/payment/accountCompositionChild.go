package payment

import (
	"fmt"
)

// AccountChild refers to a embedded type composition
type AccountChild struct {
	AccountParent
}

// RequestCreditIncrease method belogs to the AccountChild
func (ac *AccountChild) RequestCreditIncrease() float32 {
	fmt.Println("Requesting credit account increase ...")
	return 0
}
