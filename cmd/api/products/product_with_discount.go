package products

type PriceWithDiscount struct {
	Original           int    `json:"original"`
	Final              int    `json:"final"`
	DiscountPercentage string `json:"discount_percentage"`
	Currency           string `json:"currency"`
}

type ProductWithDiscount struct {
	SKU      string            `json:"sku"`
	Name     string            `json:"name"`
	Category string            `json:"category"`
	Price    PriceWithDiscount `json:"price"`
}
