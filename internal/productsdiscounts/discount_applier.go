package productsdiscounts

import (
	"discounts-applier/internal/productsdiscounts/discounts"
	"discounts-applier/internal/productsdiscounts/products"
)

type DiscountApplier interface {
	ApplyToList([]products.Product) []discounts.Product
}

func NewDiscountApplier() DiscountApplier {
	return nil
}
