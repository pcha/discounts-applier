package main

import (
	"log"
	"net/http"

	"discounts-applier/cmd/api/products"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, []products.ProductWithDiscount{
			{
				SKU:      "000001",
				Name:     "BV Lean leather ankle boots",
				Category: "boots",
				Price: products.PriceWithDiscount{
					Original:           89000,
					Final:              62300,
					DiscountPercentage: "30%",
					Currency:           "EUR",
				},
			},
		})
	})
	return r
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	log.Fatal(err)
}
