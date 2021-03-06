package products

import (
	"errors"
	"testing"

	"discounts-applier/internal/discounts/products/client"

	"github.com/stretchr/testify/assert"
)

func TestStubMongoClient_StartStub(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			"no error",
			nil,
		},
		{
			"mock error",
			errors.New("connection error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := new(StubMongoClient)

			stop := s.StartStub(tt.err)

			c, err := newMongoClient()
			assert.Equal(t, tt.err, err)
			if tt.err == nil {
				assert.Same(t, s, c)
			}

			// assert that the original new client function is saved
			stop()
			oc, err := newMongoClient()
			assert.Nil(t, err)
			assert.IsType(t, &client.WrappedClient{}, oc)
		})
	}
}
