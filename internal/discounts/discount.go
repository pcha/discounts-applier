package discounts

import "discounts-applier/internal/discounts/products"

// A Discount contains the information about the discount itself, namely the percentage and the condition.
type Discount struct {
	condition  Condition
	percentage int
}

// A Condition is a function which given the product returns if the Discount applies or not.
type Condition func(p products.Product) bool
