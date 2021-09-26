package discounts

import (
	"errors"
	"testing"

	"discounts-applier/internal/discounts/products"

	"github.com/stretchr/testify/assert"
)

func TestStubProductsRepository_StartStub(t *testing.T) {
	tests := []struct {
		name          string
		connectionURI string
		err           error
	}{
		{
			"no error",
			"valid connection URI",
			nil,
		},
		{
			"mock error",
			"invalid connection URI",
			errors.New("connection error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := new(StubProductsRepository)

			stop := s.StartStub(tt.err)

			c, err := newProductsRepository(tt.connectionURI)
			assert.Equal(t, tt.err, err)
			if tt.err == nil {
				assert.Same(t, s, c)
			}

			// assert that the original new repository function is restored
			stop()
			pr, err := newProductsRepository(tt.connectionURI)
			// if the functionality of was restored to products.NewRepository they must return equivalent values
			wpr, werr := products.NewRepository(tt.connectionURI)
			assert.Equal(t, wpr, pr)
			assert.Equal(t, werr, err)
		})
	}
}
