package dependencies

import (
	"os"

	"discounts-applier/internal/discounts"
)

const ConnectionUrlKey = "MONGO_URL"

// Dependencies return the dependencies that could need the application
type Dependencies interface {
	GetDiscountsManager() discounts.Manager
}

// RealDependencies is an implementation of Dependencies which returns the "real" dependencies and not mocks. It's the
//implementation of Dependenies used by the  main method.
type RealDependencies struct{}

// GetDiscountsManager returns an instance of internal.discounts.Manager
func (d RealDependencies) GetDiscountsManager() discounts.Manager {
	return discounts.NewManager(os.Getenv(ConnectionUrlKey))
}
