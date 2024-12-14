package order

type OrderCalculator interface {
	CalculatorTotalPrice(* Order) float64
}