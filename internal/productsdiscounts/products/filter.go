package products

type Filter interface {
	GetFilter() interface{}
}

type FilterByCategory struct {
	criteria string
}

func (f FilterByCategory) GetFilter() interface{} {
	panic("implement me")
}

func GetFilterByCategory(criteria string) FilterByCategory {
	return FilterByCategory{criteria: criteria}
}
