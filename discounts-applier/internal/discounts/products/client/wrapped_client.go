package client

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WrappedClient struct {
	MongoClient
}

func (w WrappedClient) GetDB() *WrappedDB {
	db := w.Database(getDBData().Database)
	return &WrappedDB{db}
}

func NewWrappedClient(opts ...*options.ClientOptions) (*WrappedClient, error) {
	cli, err := mongo.NewClient(opts...)
	return &WrappedClient{cli}, err
}
