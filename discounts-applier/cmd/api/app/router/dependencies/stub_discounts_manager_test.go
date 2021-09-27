package dependencies

import (
	"errors"
	"testing"

	"discounts-applier/internal/discounts/products"

	"github.com/stretchr/testify/assert"
)

func TestStubDiscountsManager_StartStub(t *testing.T) {
	tests := []struct {
		name         string
		receivingURI string
		expectedURI  string
		err          error
		testingErr   bool
	}{
		{
			"no error",
			"valid connection URI",
			"valid connection URI",
			nil,
			false,
		},
		{
			"mock error",
			"invalid connection URI",
			"invalid connection URI",
			errors.New("connection error"),
			false,
		},
		{
			"the expecting and tehe received URI don't match",
			"received uri",
			"expecting uri",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := new(StubDiscountsManager)

			st := new(testing.T)
			stop := s.StartStub(st, tt.expectedURI, tt.err)

			c, err := newDiscountsManager(tt.receivingURI)
			if tt.testingErr {
				assert.Equal(t, tt.testingErr, st.Failed())
			}
			assert.Equal(t, tt.err, err)
			if tt.err == nil {
				assert.Same(t, s, c)
			}

			// assert that the original new repository function is restored
			stop()
			dm, err := newDiscountsManager(tt.expectedURI)
			// if the functionality of newDiscountsManager was restored to discounts.NewManager they must return equivalent values
			wdm, werr := products.NewRepository(tt.expectedURI)
			assert.Equal(t, wdm, dm)
			assert.Equal(t, werr, err)
		})
	}
}
