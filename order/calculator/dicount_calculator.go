package calculator

import "solid-go/order"

type DiscountCalculator struct{
	Percentage uint
}

func (cal *DiscountCalculator) CalculatorTotalPrice(o *order.Order) float64 {
	var total float64 = 0
	for _, el := range o.Products {
		total += el.Price
	}

	return total - (total * float64(cal.Percentage)/100)
}