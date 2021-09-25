package productsdiscounts

import (
	"discounts-applier/internal/productsdiscounts/discounts"
	"discounts-applier/internal/productsdiscounts/products"
)

type Manager interface {
	GetProductsWithDiscount(filter ...products.Filter) ([]discounts.Product, error)
}

type DefaultProductsDiscounts struct{}

func (pd DefaultProductsDiscounts) GetProductsWithDiscount(filter ...products.Filter) ([]discounts.Product, error) {
	panic("implement me")
}
