package discounts

import "discounts-applier/internal/discounts/products"

const bootsCategory = "boots"
const sku3 = "000003"

var availableDiscounts = []Discount{
	{
		condition: func(p products.Product) bool {
			return p.Category == bootsCategory
		},
		percentage: 30,
	},
	{
		condition: func(p products.Product) bool {
			return p.SKU == sku3
		},
		percentage: 15,
	},
}
