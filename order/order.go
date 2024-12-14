package order

import (
	"solid-go/customer"
	"solid-go/product"
)

type Order struct {
	ID uint
	Products []product.Product
	Customer customer.Customer
}