package discounts

import (
	"testing"

	"discounts-applier/internal/discounts/products"

	"github.com/stretchr/testify/assert"
)

func TestNewDiscountApplier(t *testing.T) {
	da := NewDiscountApplier()
	assert.IsType(t, ActualDiscountApplier{}, da)
}

func TestActualDiscountApplier_ApplyToList(t *testing.T) {
	type fields struct {
		availableDiscounts []Discount
	}
	type args struct {
		products []products.Product
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Product
	}{
		{
			"base",
			fields{
				availableDiscounts: availableDiscounts,
			},
			args{
				products: []products.Product{
					{
						SKU:      "000001",
						Name:     "Boot 1",
						Category: bootsCategory,
						Price:    100000,
					},
				},
			},
			[]Product{
				{
					SKU:      "000001",
					Name:     "Boot 1",
					Category: bootsCategory,
					Price: Price{
						Original:           100000,
						Final:              70000,
						DiscountPercentage: 30,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := ActualDiscountApplier{availableDiscounts: tt.fields.availableDiscounts}
			dp := da.ApplyToList(tt.args.products)
			assert.Equal(t, tt.want, dp)
		})
	}
}

func TestActualDiscountApplier_Apply(t *testing.T) {
	type product struct {
		sku           string
		name          string
		category      string
		originalPrice int
		finalPrice    int
		discountPerc  int
	}
	tests := []struct {
		name               string
		availableDiscounts []Discount
		product            product
	}{
		{
			"Product in boots category receives 30",
			availableDiscounts,
			product{
				sku:           "00001",
				name:          "Boots 1",
				category:      bootsCategory,
				originalPrice: 100000,
				finalPrice:    70000,
				discountPerc:  30,
			},
		},
		{
			"Product with SKU " + sku3 + " get 15",
			availableDiscounts,
			product{
				sku:           sku3,
				name:          "Product",
				category:      "sandals",
				originalPrice: 10000,
				finalPrice:    8500,
				discountPerc:  15,
			},
		},
		{
			"product without discount",
			availableDiscounts,
			product{
				sku:           "00002",
				name:          "Product",
				category:      "sandals",
				originalPrice: 19000,
				finalPrice:    19000,
				discountPerc:  0,
			},
		},
		{
			"product candidate to 2 discounts get bigger",
			availableDiscounts,
			product{
				sku:           sku3,
				name:          "Boots 3",
				category:      bootsCategory,
				originalPrice: 20000,
				finalPrice:    14000,
				discountPerc:  30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := ActualDiscountApplier{tt.availableDiscounts}
			pp := products.Product{
				SKU:      tt.product.sku,
				Name:     tt.product.name,
				Category: tt.product.category,
				Price:    tt.product.originalPrice,
			}
			dp := da.Apply(pp)
			assert.Equal(t, Product{
				SKU:      tt.product.sku,
				Name:     tt.product.name,
				Category: tt.product.category,
				Price: Price{
					Original:           tt.product.originalPrice,
					Final:              tt.product.finalPrice,
					DiscountPercentage: tt.product.discountPerc,
				},
			}, dp)
		})
	}
}
