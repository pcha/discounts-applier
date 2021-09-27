package products

import (
	"context"

	"discounts-applier/internal/discounts/products/client"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Find(...Filter) ([]Product, error)
	Clean() error
	Write([]Product) error
}

type newClientFunc func(opts ...*options.ClientOptions) (client.MongoClient, error)

var newMongoClient newClientFunc = func(opts ...*options.ClientOptions) (client.MongoClient, error) {
	return client.NewWrappedClient(opts...)
}

func NewRepository(connectionURL string) (Repository, error) {
	c, err := newMongoClient(options.Client().ApplyURI(connectionURL))
	if err != nil {
		return nil, err
	}
	err = c.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return &MongoRepository{
		&client.WrappedClient{MongoClient: c},
	}, nil
}

type MongoRepository struct {
	client *client.WrappedClient
}

func (m MongoRepository) Find(filters ...Filter) ([]Product, error) {
	ctx := context.Background()
	coll := m.client.GetDB().GetCollection()
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

func (m MongoRepository) Clean() error {
	ctx := context.Background()
	coll := m.client.GetDB().GetCollection()
	_, err := coll.DeleteMany(ctx, bson.D{})
	return err
}

func (m MongoRepository) Write(products []Product) error {
	ctx := context.Background()
	coll := m.client.GetDB().GetCollection()
	prods := make([]interface{}, len(products))
	for i, p := range products {
		prods[i] = p
	}
	_, err := coll.InsertMany(ctx, prods)
	return err
}
