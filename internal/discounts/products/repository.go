package products

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func (m MongoRepository) Find(filters ...Filter) ([]Product, error) {
	ctx := context.Background()
	err := m.client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	coll := m.client.Database(getDBData().Database).Collection(getDBData().Collection)
	fil := bson.D{}
	for _, f := range filters {
		fil = append(fil, f.GetFilter())
	}
	cur, err := coll.Find(ctx, fil, options.Find().SetLimit(5))
	if err != nil {
		return nil, err
	}
	results := &[]Product{}
	err = cur.All(ctx, results)
	return *results, err
}
