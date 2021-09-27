package dependencies

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRealDependencies_GetDiscountsManager(t *testing.T) {
	tests := []struct {
		name          string
		connectionURL string
		err           error
	}{
		{
			"newDiscountsManager returns ok",
			"mongo_uri",
			nil,
		},
		{
			"newDiscountsManager returns error",
			"",
			errors.New("some error"),
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

			sdm := new(StubDiscountsManager)
			stop := sdm.StartStub(t, tt.connectionURL, tt.err)
			defer stop()

			d := RealDependencies{}
			man, err := d.GetDiscountsManager()

			assert.Equal(t, tt.err, err)
			if tt.err == nil {
				assert.Equal(t, sdm, man)
			}
		})
	}
}
