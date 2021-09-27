package router

import (
	"fmt"
	"net/http"
	"strconv"

	"discounts-applier/cmd/api/app/router/dependencies"
	"discounts-applier/internal/discounts"
	"discounts-applier/internal/discounts/products"

	"github.com/gin-gonic/gin"
	null "gopkg.in/guregu/null.v4"
)

// SetupRouter set the routers and their handlers. It receives the dependencies which will be needed by the handlers.
func SetupRouter(dep dependencies.Dependencies) (*gin.Engine, error) {
	man, err := dep.GetDiscountsManager()
	if err != nil {
		return nil, err
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/products", func(c *gin.Context) {
		filters := make([]products.Filter, 0)
		if catCrit := c.Query("category"); catCrit != "" {
			filters = append(filters, products.GetFilterByCategory(catCrit))
		}
		if pricLesCrit := c.Query("priceLessThan"); pricLesCrit != "" {
			crit, err := strconv.Atoi(pricLesCrit)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("invalid value %q for priceLessThan, the value must be a integer", pricLesCrit),
				})
				return
			}
			filters = append(filters, products.GetFilterByPriceLessThan(crit))
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
	return r, nil
}

// PresentablePrice is a nested type, it is used to Present the price and discount in PresentableProduct
type PresentablePrice struct {
	Original           int         `json:"original"`
	Final              int         `json:"final"`
	DiscountPercentage null.String `json:"discount_percentage"`
	Currency           string      `json:"currency"`
}

// PresentableProduct represent the information of internal.discounts.Product but is adapted to the presentation
//requirements.
type PresentableProduct struct {
	SKU      string           `json:"sku"`
	Name     string           `json:"name"`
	Category string           `json:"category"`
	Price    PresentablePrice `json:"price"`
}

// present take an internal.discounts.Product and return its representation as a PresentableProduct
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
