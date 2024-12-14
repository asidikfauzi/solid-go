package payment

import "fmt"

type APayment struct {}

func (a *APayment) Pay(amount float64) {
	fmt.Println("AMOUNT THAT SHOULD BE PAID WITH A VENDOR IS :", amount)
}