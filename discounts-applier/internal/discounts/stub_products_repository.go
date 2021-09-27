package discounts

import (
	"discounts-applier/internal/discounts/products"
)

type StubProductsRepository struct {
	products.MockRepository
	originalNewFunc NewProductsRepoFunc
}

type StopStub func()

func (s *StubProductsRepository) StartStub(err error) StopStub {
	s.originalNewFunc = newProductsRepository
	newProductsRepository = func(connectionURI string) (products.Repository, error) {
		if err != nil {
			return nil, err
		}
		return s, nil
	}
	return func() {
		newProductsRepository = s.originalNewFunc
	}
}
