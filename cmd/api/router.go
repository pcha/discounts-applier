package main

import (
	"net/http"
	"strconv"

	"discounts-applier/cmd/api/dependencies"
	"discounts-applier/internal/productsdiscounts/discounts"
	"discounts-applier/internal/productsdiscounts/products"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

func setupRouter(dep dependencies.Dependencies) *gin.Engine {
	r := gin.Default()
	r.GET("/products", func(c *gin.Context) {
		man := dep.GetProductsDiscounts()
		filters := make([]products.Filter, 0)
		if catCrit := c.Query("category"); catCrit != "" {
			filters = append(filters, products.GetFilterByCategory(catCrit))
		}
		dp, err := man.GetProductsWithDiscount(filters...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		pp := make([]PresentableProduct, len(dp))
		for i, p := range dp {
			pp[i] = present(p)
		}
		c.JSON(http.StatusOK, pp)
	})
	return r
}

type PresentablePrice struct {
	Original           int         `json:"original"`
	Final              int         `json:"final"`
	DiscountPercentage null.String `json:"discount_percentage"`
	Currency           string      `json:"currency"`
}

type PresentableProduct struct {
	SKU      string           `json:"sku"`
	Name     string           `json:"name"`
	Category string           `json:"category"`
	Price    PresentablePrice `json:"price"`
}

func present(p discounts.Product) PresentableProduct {
	var disc null.String
	if p.Price.DiscountPercentage == 0 {
		disc = null.StringFromPtr(nil)
	} else {
		disc = null.StringFrom(strconv.Itoa(p.Price.DiscountPercentage) + "%")
	}
	return PresentableProduct{
		SKU:      p.SKU,
		Name:     p.Name,
		Category: p.Category,
		Price: PresentablePrice{
			Original:           p.Price.Original,
			Final:              p.Price.Final,
			DiscountPercentage: disc,
			Currency:           "EUR",
		},
	}
}
