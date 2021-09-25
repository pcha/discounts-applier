package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"discounts-applier/cmd/api/dependencies/mocks"
	"discounts-applier/internal/productsdiscounts/discounts"
	pdmocks "discounts-applier/internal/productsdiscounts/mocks"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func Test_setupRouter(t *testing.T) {
	withDiscountsProductsMock := []discounts.Product{
		{
			SKU:      "000001",
			Name:     "BV Lean leather ankle boots",
			Category: "boots",
			Price: discounts.Price{
				Original:           89000,
				Final:              62300,
				DiscountPercentage: 30,
			},
		},
		{
			SKU:      "000002",
			Name:     "BV Lean leather ankle boots",
			Category: "boots",
			Price: discounts.Price{
				Original:           99000,
				Final:              69300,
				DiscountPercentage: 30,
			},
		},
		{
			SKU:      "000004",
			Name:     "Naima embellished suede sandals",
			Category: "sandals",
			Price: discounts.Price{
				Original:           79500,
				Final:              79500,
				DiscountPercentage: 0,
			},
		},
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

	tests := []struct {
		name           string
		url            string
		mockedProducts []discounts.Product
		mockedErr      error
		wantedCode     int
		wantedBody     string
	}{
		{
			"GET /products when there are products with discounts",
			"/products",
			withDiscountsProductsMock,
			nil,
			http.StatusOK,
			withDiscountsWantedBody,
		},
		{
			"GET /products receives an error",
			"/products",
			nil,
			errors.New("some error"),
			http.StatusInternalServerError,
			`{"error": "some error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			pd := new(pdmocks.Manager)
			pd.On("GetProductsWithDiscount").Return(tt.mockedProducts, tt.mockedErr)
			dep := new(mocks.Dependencies)
			dep.On("GetProductsDiscounts").Return(pd)

			// when
			r := setupRouter(dep)

			// then
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
