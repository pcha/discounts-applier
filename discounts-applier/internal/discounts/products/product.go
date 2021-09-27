package products

type Product struct {
	SKU      string `bson:"sku"`
	Name     string `bson:"name"`
	Category string `bson:"category"`
	Price    int    `bson:"price"`
}
