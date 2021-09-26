package dependencies

import (
	"os"
	"testing"

	"discounts-applier/internal/discounts"

	"github.com/stretchr/testify/assert"
)

func TestRealDependencies_GetDiscountsManager(t *testing.T) {
	tests := []struct {
		name          string
		connectionURL string
		want          discounts.Manager
	}{
		{
			"with envvar set",
			"mongo_url",
			discounts.NewManager("mongo_url"),
		},
		{
			"without envvar set",
			"",
			discounts.NewManager(""),
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
			man := d.GetDiscountsManager()
			assert.Equal(t, tt.want, man)
		})
	}
}
