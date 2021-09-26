// Package discounts provides all the needed to obtain products with discounts applied.
package discounts

import (
	"discounts-applier/internal/discounts/products"
)

// A Manager return a slice of Product with discount applied if appropriate. It receives filters to select the products
// to obtain.
type Manager interface {
	GetProductsWithDiscount(filters ...products.Filter) ([]Product, error)
}

// ActualManager is the Default implementation of Manager, it is the used during the program execution.
type ActualManager struct {
	products        products.Repository
	discountApplier DiscountApplier
}

// GetProductsWithDiscount returns the filtered products with the correspondent discounts applied.
func (pd ActualManager) GetProductsWithDiscount(filters ...products.Filter) ([]Product, error) {
	pp, err := pd.products.Find(filters...)
	if err != nil {
		return nil, err
	}
	return pd.discountApplier.ApplyToList(pp), nil
}

// NewManager returns an instance of Manager, in concrete an ActualManager.
func NewManager(connectionURL string) Manager {
	return &ActualManager{}
}
