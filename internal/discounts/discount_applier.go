package discounts

import (
	"discounts-applier/internal/discounts/products"
)

// A DiscountApplier has the ability of apply the appropriate discounts to the given products.
type DiscountApplier interface {
	ApplyToList([]products.Product) []Product
}

// ActualDiscountApplier is the default implementation of DiscountApplier.
type ActualDiscountApplier struct {
	//availableDiscounts []Di
}

// NewDiscountApplier returns an instance of DiscountApplier, namely an ActualDiscountApplier.
func NewDiscountApplier() DiscountApplier {
	return nil
}
