package discounts

type Price struct {
	Original           int `json:"original"`
	Final              int `json:"final"`
	DiscountPercentage int `json:"discount_percentage"`
}

type Product struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    Price  `json:"price"`
}
