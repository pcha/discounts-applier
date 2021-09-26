package dependencies

import (
	"os"

	"discounts-applier/internal/productsdiscounts"
)

const ConnectionUrlKey = "MONGO_URL"

type Dependencies interface {
	GetProductsDiscounts() productsdiscounts.Manager
}

type RealDependencies struct{}

func (d RealDependencies) GetProductsDiscounts() productsdiscounts.Manager {
	return productsdiscounts.NewManager(os.Getenv(ConnectionUrlKey))
}
