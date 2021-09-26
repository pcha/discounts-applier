package productsdiscounts

import (
	"discounts-applier/internal/productsdiscounts/discounts"
	"discounts-applier/internal/productsdiscounts/products"
)

type Manager interface {
	GetProductsWithDiscount(filters ...products.Filter) ([]discounts.Product, error)
}

type ActualManager struct {
	products        products.Repository
	discountApplier DiscountApplier
}

func (pd ActualManager) GetProductsWithDiscount(filters ...products.Filter) ([]discounts.Product, error) {
	pp, err := pd.products.Find(filters...)
	if err != nil {
		return nil, err
	}
	return pd.discountApplier.ApplyToList(pp), nil
}

func NewManager(connectionURL string) Manager {
	return &ActualManager{}
}
