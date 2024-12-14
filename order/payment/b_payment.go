package payment

import "fmt"

type Bpayment struct{}

func (a *Bpayment) Pay(amount float64) {
	fmt.Println("AMOUNT THAT SHOULD BE PAID WITH B VENDOR IS :", amount)
}