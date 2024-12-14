package order

import "fmt"

type OrderShipping struct{}

func(s *OrderShipping) Shipping(o *Order) {
	fmt.Println("==== SHIPPING REQUESTED ! ORDER ID :", o.ID, "====")
}