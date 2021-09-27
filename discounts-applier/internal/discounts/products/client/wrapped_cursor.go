package client

import "context"

type WrappedCursor struct {
	MongoCursor
}

func (c *WrappedCursor) Unmarshall(results interface{}) error {
	return c.All(context.Background(), results)
}
