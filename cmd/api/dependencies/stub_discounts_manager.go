package dependencies

import (
	"testing"

	"discounts-applier/internal/discounts"
)

type StubDiscountsManager struct {
	discounts.MockManager
	originalNewFunc NewManFunc
	t               *testing.T
}

type StopStub func()

func (s *StubDiscountsManager) StartStub(t *testing.T, expectedUri string, returningErr error) StopStub {
	s.originalNewFunc = newDiscountsManager
	newDiscountsManager = func(connectionURI string) (discounts.Manager, error) {
		if expectedUri != connectionURI {
			t.Errorf("StubDiscountManager was started expecting the uri %q but was instantiated receiving %q", expectedUri, connectionURI)
		}
		if returningErr != nil {
			return nil, returningErr
		}
		return s, nil
	}
	return func() {
		newDiscountsManager = s.originalNewFunc
	}
}
