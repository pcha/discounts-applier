package productsdiscounts

import (
	"errors"
	"testing"

	"discounts-applier/internal/productsdiscounts/discounts"
	discountsmocks "discounts-applier/internal/productsdiscounts/mocks"
	"discounts-applier/internal/productsdiscounts/products"
	productsmocks "discounts-applier/internal/productsdiscounts/products/mocks"

	"github.com/stretchr/testify/assert"
)

func TestNewManager(t *testing.T) {
	url := ""
	man := NewManager(url)
	assert.Equal(t, &ActualManager{
		products:        products.NewRepository(url),
		discountApplier: NewDiscountApplier(),
	}, man)
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

	discountedProduct1 := discounts.Product{
		SKU:      "0001",
		Name:     "Product 1",
		Category: "cat1",
		Price: discounts.Price{
			Original:           100000,
			Final:              70000,
			DiscountPercentage: 30,
		},
	}
	discountedProduct2 := discounts.Product{
		SKU:      "0002",
		Name:     "Product 2",
		Category: "cat2",
		Price: discounts.Price{
			Original:           200000,
			Final:              180000,
			DiscountPercentage: 10,
		},
	}
	discountedProduct3 := discounts.Product{
		SKU:      "0003",
		Name:     "Product 3",
		Category: "cat1",
		Price: discounts.Price{
			Original:           150000,
			Final:              150000,
			DiscountPercentage: 0,
		},
	}

	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    []discounts.Product
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
			want: []discounts.Product{
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
			want: []discounts.Product{
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
			pr := new(productsmocks.Repository)
			filters := make([]interface{}, len(tt.args.filters))
			for i, f := range tt.args.filters {
				filters[i] = f
			}
			pr.On("Find", filters...).Return(tt.mocks.products, tt.mocks.productsErr)

			da := new(discountsmocks.DiscountApplier)
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
