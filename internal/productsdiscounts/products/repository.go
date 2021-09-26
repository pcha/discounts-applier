package products

type Repository interface {
	Find(...Filter) ([]Product, error)
}

func NewRepository(connectionURL string) Repository {
	return nil
}
