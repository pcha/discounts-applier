package client

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type WrappedCursor struct {
	*mongo.Cursor
}

func (c *WrappedCursor) Unmarshall(results interface{}) error {
	return c.All(context.Background(), results)
}
