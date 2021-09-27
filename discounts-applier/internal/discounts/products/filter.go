package products

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Filter interface {
	GetFilter() bson.E
}

type FilterByCategory struct {
	criteria string
}

func (f FilterByCategory) GetFilter() bson.E {
	return bson.E{
		Key:   "category",
		Value: f.criteria,
	}
}

func GetFilterByCategory(criteria string) FilterByCategory {
	return FilterByCategory{criteria: criteria}
}

type FilterByPriceLessThan struct {
	criteria int
}

func (f FilterByPriceLessThan) GetFilter() bson.E {
	return bson.E{
		Key: "price",
		Value: bson.M{
			"$lte": f.criteria,
		},
	}
}

func GetFilterByPriceLessThan(criteria int) Filter {
	return FilterByPriceLessThan{criteria: criteria}
}
