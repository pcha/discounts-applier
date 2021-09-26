package products

import "go.mongodb.org/mongo-driver/bson"

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
