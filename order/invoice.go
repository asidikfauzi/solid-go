package order

import (
	"fmt"
)

type OrderInvoice interface {
	GenerateInvoice(*Order, OrderCalculator)
}

type StandardInvoice struct{}

func (i *StandardInvoice) GenerateInvoice(o *Order, cal OrderCalculator) {
	fmt.Println("==== INVOICE ====")
	var total float64 = 0
	for _, el := range o.Products {
		fmt.Println(el.Name, " ", el.Price)
	}
	
	total = cal.CalculatorTotalPrice(o)

	fmt.Println("----------------")
	fmt.Println("TOTAL: ", total)
}

type IntegratedInvoice struct {
	Payment Payment
}

func (i *IntegratedInvoice) GenerateInvoice(o *Order, cal OrderCalculator) {
	total := cal.CalculatorTotalPrice(o)
	i.Pay(total)
}

func (i *IntegratedInvoice) Pay(amount float64) {
	i.Payment.Pay(amount)
}
