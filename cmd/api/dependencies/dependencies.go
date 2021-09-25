package dependencies

import (
	"discounts-applier/internal/productsdiscounts"
)

type Dependencies interface {
	GetProductsDiscounts() productsdiscounts.Manager
}

type RealDependencies struct{}

func (d RealDependencies) GetProductsDiscounts() productsdiscounts.Manager {
	return nil
}
