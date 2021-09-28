package client

import "go.mongodb.org/mongo-driver/mongo"

type WrappedDB struct {
	*mongo.Database
}

func (db *WrappedDB) GetCollection() MongoCollection {
	coll := db.Collection(getDBData().Collection)
	return &WrappedCollection{coll}
}
