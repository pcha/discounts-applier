package dependencies

import (
	"os"
	"testing"

	"discounts-applier/internal/productsdiscounts"

	"github.com/stretchr/testify/assert"
)

func TestRealDependencies_GetProductsDiscounts(t *testing.T) {
	tests := []struct {
		name          string
		connectionURL string
		want          productsdiscounts.Manager
	}{
		{
			"with envvar set",
			"mongo_url",
			productsdiscounts.NewManager("mongo_url"),
		},
		{
			"without envvar set",
			"",
			productsdiscounts.NewManager(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalVal := os.Getenv(ConnectionUrlKey)
			defer func(t *testing.T) {
				err := os.Setenv(ConnectionUrlKey, originalVal)
				if err != nil {
					t.Fatal(err)
				}
			}(t)
			err := os.Setenv(ConnectionUrlKey, tt.connectionURL)
			if err != nil {
				t.Fatal(err)
			}
			d := RealDependencies{}
			man := d.GetProductsDiscounts()
			assert.Equal(t, tt.want, man)
		})
	}
}
