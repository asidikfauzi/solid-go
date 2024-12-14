package calculator

import "solid-go/order"

type StandardCalculator struct{}

func (cal *StandardCalculator) CalculatorTotalPrice(o *order.Order) float64 {
	var total float64 = 0
	for _, el := range o.Products {
		total += el.Price
	}

	return total
}