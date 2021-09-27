package client

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type WrappedCollection struct {
	MongoCollection
}

func (c *WrappedCollection) FindFive(filter interface{}) (MongoCursor, error) {
	cur, err := c.Find(context.Background(), filter, options.Find().SetLimit(5))
	return &WrappedCursor{cur}, err
}
