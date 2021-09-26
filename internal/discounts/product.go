package discounts

// Price id nested in Product to group the price and discount information
type Price struct {
	Original           int `json:"original"`
	Final              int `json:"final"`
	DiscountPercentage int `json:"discount_percentage"`
}

// Product has te information of the product and the price with the discount, if any, applied
type Product struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    Price  `json:"price"`
}
