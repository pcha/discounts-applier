package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFilterByCategory_GetFilter(t *testing.T) {
	cf := GetFilterByCategory("cat1")
	bf := cf.GetFilter()

	assert.Equal(t, bson.E{
		"category",
		"cat1",
	}, bf)
}
