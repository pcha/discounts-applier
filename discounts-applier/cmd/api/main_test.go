package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"discounts-applier/cmd/api/dependencies"
	"discounts-applier/internal/discounts"
	"discounts-applier/internal/discounts/products"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func Test_setupRouter(t *testing.T) {
	product1 := discounts.Product{
		SKU:      "000001",
		Name:     "BV Lean leather ankle boots",
		Category: "boots",
		Price: discounts.Price{
			Original:           89000,
			Final:              62300,
			DiscountPercentage: 30,
		},
	}
	product2 := discounts.Product{
		SKU:      "000002",
		Name:     "BV Lean leather ankle boots",
		Category: "boots",
		Price: discounts.Price{
			Original:           99000,
			Final:              69300,
			DiscountPercentage: 30,
		},
	}
	product3 := discounts.Product{
		SKU:      "000004",
		Name:     "Naima embellished suede sandals",
		Category: "sandals",
		Price: discounts.Price{
			Original:           79500,
			Final:              79500,
			DiscountPercentage: 0,
		},
	}

	withDiscountsProductsMock := []discounts.Product{
		product1,
		product2,
		product3,
	}
	withDiscountsWantedBody := `[
		{
		  "sku": "000001",
		  "name": "BV Lean leather ankle boots",
		  "category": "boots",
		  "price": {
			"original": 89000,
			"final": 62300,
			"discount_percentage": "30%",
			"currency": "EUR"
		  }
		},
		{
			"sku":      "000002",
			"name":     "BV Lean leather ankle boots",
			"category": "boots",
			"price": {
				"original":           99000,
				"final":              69300,
				"discount_percentage": "30%",
				"currency": "EUR"
			}
		},
		{
		  "sku": "000004",
		  "name": "Naima embellished suede sandals",
		  "category": "sandals",
		  "price": {
			"original": 79500,
			"final": 79500,
			"discount_percentage": null,
			"currency": "EUR"
		  }
		}
	]`

	filterByCategoryProductsMock := []discounts.Product{
		product1,
		product2,
	}
	filterByCategoryWantedBody := `[
		{
		  "sku": "000001",
		  "name": "BV Lean leather ankle boots",
		  "category": "boots",
		  "price": {
			"original": 89000,
			"final": 62300,
			"discount_percentage": "30%",
			"currency": "EUR"
		  }
		},
		{
			"sku":      "000002",
			"name":     "BV Lean leather ankle boots",
			"category": "boots",
			"price": {
				"original":           99000,
				"final":              69300,
				"discount_percentage": "30%",
				"currency": "EUR"
			}
		}
	]`
	filterByPriceLessThanProductsMock := []discounts.Product{
		product1,
		product3,
	}
	filterByPriceLessThanWantedBody := `[
		{
		  "sku": "000001",
		  "name": "BV Lean leather ankle boots",
		  "category": "boots",
		  "price": {
			"original": 89000,
			"final": 62300,
			"discount_percentage": "30%",
			"currency": "EUR"
		  }
		},
		{
		  "sku": "000004",
		  "name": "Naima embellished suede sandals",
		  "category": "sandals",
		  "price": {
			"original": 79500,
			"final": 79500,
			"discount_percentage": null,
			"currency": "EUR"
		  }
		}
	]`

	tests := []struct {
		name                string
		url                 string
		mockedProducts      []discounts.Product
		mockedHandlerErr    error
		mockedDependencyErr error
		filters             []products.Filter
		wantedCode          int
		wantedBody          string
	}{
		{
			name:                "can't setup routes due depenedency error",
			url:                 "/products",
			mockedDependencyErr: errors.New("dependency err"),
		},
		{
			name:           "GET /products when there are products",
			url:            "/products",
			mockedProducts: withDiscountsProductsMock,
			wantedCode:     http.StatusOK,
			wantedBody:     withDiscountsWantedBody,
		},
		{
			name:             "GET /products receives an error",
			url:              "/products",
			mockedHandlerErr: errors.New("some error"),
			wantedCode:       http.StatusInternalServerError,
			wantedBody:       `{"error": "some error"}`,
		},
		{
			name:           "GET /products?category={category} filter category",
			url:            "/products?category=boots",
			mockedProducts: filterByCategoryProductsMock,
			filters: []products.Filter{
				products.GetFilterByCategory("boots"),
			},
			wantedCode: http.StatusOK,
			wantedBody: filterByCategoryWantedBody,
		},
		{
			name:           "GET /products?priceLessThan={price} filter by price less than or equal the given",
			url:            "/products?priceLessThan=89000",
			mockedProducts: filterByPriceLessThanProductsMock,
			filters: []products.Filter{
				products.GetFilterByPriceLessThan(89000),
			},
			wantedCode: http.StatusOK,
			wantedBody: filterByPriceLessThanWantedBody,
		},
		{
			name:           "GET /products?priceLessThan={string} fails for invalid format filter",
			url:            "/products?priceLessThan=3312.50",
			mockedProducts: filterByPriceLessThanProductsMock,
			wantedCode:     http.StatusBadRequest,
			wantedBody:     `{"error": "invalid value \"3312.50\" for priceLessThan, the value must be a integer"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			pd := new(discounts.MockManager)
			filters := make([]interface{}, len(tt.filters))
			for i, f := range tt.filters {
				filters[i] = f
			}
			pd.On("GetProductsWithDiscount", filters...).Return(tt.mockedProducts, tt.mockedHandlerErr)
			dep := new(dependencies.MockDependencies)
			if tt.mockedDependencyErr != nil {
				dep.On("GetDiscountsManager").Return(nil, tt.mockedDependencyErr)
			}
			dep.On("GetDiscountsManager").Return(pd, nil)

			// when
			r, err := setupRouter(dep)

			// then
			assert.Equal(t, tt.mockedDependencyErr, err)
			if tt.mockedDependencyErr != nil {
				return
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.wantedCode, w.Code)
			assert.JSONEq(t, tt.wantedBody, w.Body.String())
		})
	}
}

func Test_present(t *testing.T) {
	type args struct {
		sku           string
		name          string
		category      string
		originalPrice int
		finalPrice    int
		intDiscount   int
		strDiscount   null.String
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"WithDiscount",
			args{
				sku:           "000001",
				name:          "Product with discount",
				category:      "discount",
				originalPrice: 100000,
				finalPrice:    700000,
				intDiscount:   30,
				strDiscount:   null.StringFrom("30%"),
			},
		},
		{
			"WithoutDiscount",
			args{
				sku:           "000002",
				name:          "Product without discount",
				category:      "non-discount",
				originalPrice: 80000,
				finalPrice:    80000,
				intDiscount:   0,
				strDiscount:   null.StringFromPtr(nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dp := discounts.Product{
				SKU:      tt.args.sku,
				Name:     tt.args.name,
				Category: tt.args.category,
				Price: discounts.Price{
					Original:           tt.args.originalPrice,
					Final:              tt.args.finalPrice,
					DiscountPercentage: tt.args.intDiscount,
				},
			}

			pp := present(dp)

			assert.Equal(t, PresentableProduct{
				SKU:      tt.args.sku,
				Name:     tt.args.name,
				Category: tt.args.category,
				Price: PresentablePrice{
					Original:           tt.args.originalPrice,
					Final:              tt.args.finalPrice,
					DiscountPercentage: tt.args.strDiscount,
					Currency:           "EUR",
				},
			}, pp)
		})
	}
}
