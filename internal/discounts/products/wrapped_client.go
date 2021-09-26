package products

type WrappedClient struct {
	MongoClient
}

func (w WrappedClient) getDB() MongoDB {
	db := w.Database(getDBData().Database)
	return MongoDB(db)
}
