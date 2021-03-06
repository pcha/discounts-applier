// Code generated by mockery 2.9.4. DO NOT EDIT.

package discounts

import (
	products "discounts-applier/internal/discounts/products"

	mock "github.com/stretchr/testify/mock"
)

// MockDiscountApplier is an autogenerated mock type for the DiscountApplier type
type MockDiscountApplier struct {
	mock.Mock
}

// ApplyToList provides a mock function with given fields: _a0
func (_m *MockDiscountApplier) ApplyToList(_a0 []products.Product) []Product {
	ret := _m.Called(_a0)

	var r0 []Product
	if rf, ok := ret.Get(0).(func([]products.Product) []Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Product)
		}
	}

	return r0
}
