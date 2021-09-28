package client

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WrappedClient struct {
	*mongo.Client
}

func (w WrappedClient) GetDB() MongoDatabase {
	db := w.Database(getDBData().Database)
	return &WrappedDB{db}
}

func NewMongoClient(opts ...*options.ClientOptions) (MongoClient, error) {
	cli, err := mongo.NewClient(opts...)
	return MongoClient(&WrappedClient{cli}), err
}
