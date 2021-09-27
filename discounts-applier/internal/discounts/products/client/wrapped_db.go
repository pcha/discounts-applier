package client

type WrappedDB struct {
	MongoDatabase
}

func (db *WrappedDB) GetCollection() *WrappedCollection {
	coll := db.Collection(getDBData().Collection)
	return &WrappedCollection{coll}
}
