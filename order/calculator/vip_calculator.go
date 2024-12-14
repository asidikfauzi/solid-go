package calculator

import "solid-go/order"

type VipCalculator struct{
	Amount uint
}

func (cal *VipCalculator) CalculatorTotalPrice(o *order.Order) float64 {
	var total float64 = 0
	for _, el := range o.Products {
		total += el.Price
	}

	return total - float64(cal.Amount)
}