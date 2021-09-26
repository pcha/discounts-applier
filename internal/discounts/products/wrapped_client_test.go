package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestWrappedClient_getDB(t *testing.T) {
	cli := new(MockMongoClient)
	mockedDB := new(mongo.Database)
	cli.On("Database", getDBData().Database).Return(mockedDB)
	wc := WrappedClient{cli}

	db := wc.getDB()
	assert.Same(t, mockedDB, db)
}
