package main

import (
	"solid-go/order"
	"solid-go/order/calculator"
	"solid-go/order/payment"
	"solid-go/product"
)

func main() {
	o := order.Order{
		ID: 1,
		Products: []product.Product{
			{
				ID: 1,
				Name: "Book",
				Price: 3000,
			},
			{
				ID: 2,
				Name: "Pen",
				Price: 5000,
			},
		},
	}

	inv := &order.IntegratedInvoice{
		Payment: &payment.Bpayment{},
	}
	
	cal := calculator.StandardCalculator{}
	inv.GenerateInvoice(&o, &cal)

	cal2 := calculator.DiscountCalculator{Percentage: 50}
	inv.GenerateInvoice(&o, &cal2)

	cal3 := calculator.VipCalculator{Amount: 1000}
	inv.GenerateInvoice(&o, &cal3)
}