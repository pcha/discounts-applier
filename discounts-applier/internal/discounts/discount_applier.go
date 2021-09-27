package discounts

import (
	"discounts-applier/internal/discounts/products"
)

// A DiscountApplier has the ability of apply the appropriate discounts to the given products.
type DiscountApplier interface {
	ApplyToList([]products.Product) []Product
}

// ActualDiscountApplier is the default implementation of DiscountApplier. It contains the available discounts.
type ActualDiscountApplier struct {
	availableDiscounts []Discount
}

// ApplyToList receives a list of products and returns them with ths corresponding discounts applied.
func (da ActualDiscountApplier) ApplyToList(products []products.Product) []Product {
	dp := make([]Product, len(products))
	for i, p := range products {
		dp[i] = da.Apply(p)
	}
	return dp
}

func (da ActualDiscountApplier) Apply(p products.Product) Product {
	perc := da.findDiscountPercentageFor(p)
	return Product{
		SKU:      p.SKU,
		Name:     p.Name,
		Category: p.Category,
		Price: Price{
			Original:           p.Price,
			Final:              p.Price * (100 - perc) / 100,
			DiscountPercentage: perc,
		},
	}
}

func (da ActualDiscountApplier) findDiscountPercentageFor(p products.Product) int {
	perc := 0
	for _, d := range da.availableDiscounts {
		if d.condition(p) && d.percentage > perc {
			perc = d.percentage
		}
	}
	return perc
}

// NewDiscountApplier returns an instance of DiscountApplier, namely an ActualDiscountApplier.
func NewDiscountApplier() DiscountApplier {
	return ActualDiscountApplier{
		availableDiscounts: availableDiscounts,
	}
}
