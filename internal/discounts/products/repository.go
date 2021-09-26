package products

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Find(...Filter) ([]Product, error)
}

type newClientFunc func(opts ...*options.ClientOptions) (MongoClient, error)

var newMongoClient newClientFunc = func(opts ...*options.ClientOptions) (MongoClient, error) {
	return mongo.NewClient(opts...)
}

func NewRepository(connectionURL string) (Repository, error) {
	c, err := newMongoClient(options.Client().ApplyURI(connectionURL))
	if err != nil {
		return nil, err
	}
	return &MongoRepository{
		c,
	}, nil
}

type MongoRepository struct {
	client MongoClient
}

func (m MongoRepository) Find(filter ...Filter) ([]Product, error) {
	panic("implement me")
}
