package discounts

import (
	"errors"
	"testing"

	"discounts-applier/internal/discounts/products"

	"github.com/stretchr/testify/assert"
)

func TestNewManager(t *testing.T) {
	tests := []struct {
		name          string
		connectionURI string
		err           error
	}{
		{
			"products repository returns ok",
			"some uri",
			nil,
		},
		{
			"products repository returns error",
			"some uri",
			errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spr := new(StubProductsRepository)
			stop := spr.StartStub(tt.err)
			defer stop()
			man, err := NewManager(tt.connectionURI)
			assert.Equal(t, tt.err, err)
			if tt.err == nil {
				assert.Equal(t, &ActualManager{
					products:        spr,
					discountApplier: NewDiscountApplier(),
				}, man)
			}
		})
	}
}

func TestActualManager_GetProductsWithDiscount(t *testing.T) {
	type mocks struct {
		products    []products.Product
		productsErr error
	}
	type args struct {
		filters []products.Filter
	}
	product1 := products.Product{
		SKU:      "0001",
		Name:     "Product 1",
		Category: "cat1",
		Price:    100000,
	}
	product2 := products.Product{
		SKU:      "0002",
		Name:     "Product 2",
		Category: "cat2",
		Price:    200000,
	}
	product3 := products.Product{
		SKU:      "0003",
		Name:     "Prduct 3",
		Category: "cat1",
		Price:    150000,
	}

	discountedProduct1 := Product{
		SKU:      "0001",
		Name:     "Product 1",
		Category: "cat1",
		Price: Price{
			Original:           100000,
			Final:              70000,
			DiscountPercentage: 30,
		},
	}
	discountedProduct2 := Product{
		SKU:      "0002",
		Name:     "Product 2",
		Category: "cat2",
		Price: Price{
			Original:           200000,
			Final:              180000,
			DiscountPercentage: 10,
		},
	}
	discountedProduct3 := Product{
		SKU:      "0003",
		Name:     "Product 3",
		Category: "cat1",
		Price: Price{
			Original:           150000,
			Final:              150000,
			DiscountPercentage: 0,
		},
	}

	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    []Product
		wantErr bool
	}{
		{
			name: "without filters",
			mocks: mocks{
				products: []products.Product{
					product1,
					product2,
					product3,
				},
			},
			args: args{
				filters: nil,
			},
			want: []Product{
				discountedProduct1,
				discountedProduct2,
				discountedProduct3,
			},
			wantErr: false,
		},
		{
			name: "with filter",
			mocks: mocks{
				products: []products.Product{
					product1,
					product3,
				},
			},
			args: args{
				filters: []products.Filter{
					products.GetFilterByCategory("cat1"),
				},
			},
			want: []Product{
				discountedProduct1,
				discountedProduct3,
			},
			wantErr: false,
		},
		{
			name: "returning error",
			mocks: mocks{
				products:    nil,
				productsErr: errors.New("connection error"),
			},
			args: args{
				filters: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := new(products.MockRepository)
			filters := make([]interface{}, len(tt.args.filters))
			for i, f := range tt.args.filters {
				filters[i] = f
			}
			pr.On("Find", filters...).Return(tt.mocks.products, tt.mocks.productsErr)

			da := new(MockDiscountApplier)
			da.On("ApplyToList", tt.mocks.products).Return(tt.want)

			man := &ActualManager{
				products:        pr,
				discountApplier: da,
			}

			pd, err := man.GetProductsWithDiscount(tt.args.filters...)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, pd)
		})
	}
}
