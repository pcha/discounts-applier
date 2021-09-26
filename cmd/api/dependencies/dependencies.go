package dependencies

import (
	"os"

	"discounts-applier/internal/discounts"
)

const ConnectionUrlKey = "MONGO_URL"

// Dependencies return the dependencies that could need the application
type Dependencies interface {
	GetDiscountsManager() (discounts.Manager, error)
}

// RealDependencies is an implementation of Dependencies which returns the "real" dependencies and not mocks. It's the
// implementation of Dependencies used by the  main method.
type RealDependencies struct{}

type NewManFunc func(uri string) (discounts.Manager, error)

var newDiscountsManager NewManFunc = discounts.NewManager

// GetDiscountsManager returns an instance of internal.discounts.Manager
func (d RealDependencies) GetDiscountsManager() (discounts.Manager, error) {
	return newDiscountsManager(os.Getenv(ConnectionUrlKey))
}
