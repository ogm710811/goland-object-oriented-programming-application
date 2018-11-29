package main

import (
	"github.com/ogarciamartinez/gopherpay/payment"
)

func main() {
	// declaring a variable of type option interface
	var option payment.Option

	// for credit card payment
	option = payment.CreateCreditCardAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123,
	)

	option.ProcessPayment(500)

	// for cash payment
	option = payment.CreateCashAccount()
	option.ProcessPayment(500)

	// fmt.Printf("Owner name: %v\n", credit.OwnerName())
	// fmt.Printf("Card number: %v\n", credit.CardNumber())
	// err := credit.SetCardNumber("invalid card")
	// if err != nil {
	// 	fmt.Printf("That didn't work: %v\n", err)
	// }
	// fmt.Printf("Available credit: %v\n", credit.AvailableCredit())

	// for credit account payment using channels example
	chargeCh := make(chan float32)
	payment.CreateCreditAccount("Debra Ingram", chargeCh)
	chargeCh <- 500
	// var a string
	// fmt.Scanln(a)

	// type embedding implementation
	ac := &payment.AccountChild{}
	ac.AvailableFuncs()
	ac.ProcessPayment(500)
	ac.RequestCreditIncrease()

}
