package productsdiscounts

import (
	"discounts-applier/internal/productsdiscounts/discounts"
	"discounts-applier/internal/productsdiscounts/products"
)

type Manager interface {
	GetProductsWithDiscount(filter ...products.Filter) ([]discounts.Product, error)
}

type ActualManager struct {
}

func (pd ActualManager) GetProductsWithDiscount(filter ...products.Filter) ([]discounts.Product, error) {
	panic("implement me")
}

func NewManager(connectionURL string) Manager {
	return &ActualManager{}
}
