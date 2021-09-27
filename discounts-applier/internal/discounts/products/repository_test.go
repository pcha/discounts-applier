package products

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewRepository(t *testing.T) {
	tests := []struct {
		name          string
		connectionURL string
		wantErr       error
	}{
		{
			"ok",
			"mongodb://url",
			nil,
		},
		{
			"error",
			"not_mongo_url",
			errors.New("url error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := NewStubMongoClient()
			sc.MockMongoClient.On("Connect", mock.Anything).Return(nil)
			stop := sc.StartStub(tt.wantErr)
			defer stop()

			r, err := NewRepository(tt.connectionURL)
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				if assert.IsType(t, &MongoRepository{}, r) {
					mr := r.(*MongoRepository)
					assert.Same(t, sc, mr.client.MongoClient)
				}
			}
		})
	}
}

func TestMongoRepository_Find(t *testing.T) {
	connectionURI := "mongodb+srv://unittests:s6Xt4KB1q5y28N1o@cluster0.p7omq.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	prod1 := Product{
		SKU:      "0001",
		Name:     "Product 1",
		Category: "cat2",
		Price:    10000,
	}
	prod2 := Product{
		SKU:      "0002",
		Name:     "Product2",
		Category: "cat1",
		Price:    22222,
	}
	prod3 := Product{
		SKU:      "0003",
		Name:     "Product 3",
		Category: "cat3",
		Price:    15000,
	}
	tests := []struct {
		name          string
		connectionURI string
		filters       []Filter
		expected      []Product
	}{
		{
			name:          "base",
			connectionURI: connectionURI,
			expected: []Product{
				prod1,
				prod2,
				prod3,
			},
		},
		{
			name:          "with category filter",
			connectionURI: connectionURI,
			filters: []Filter{
				GetFilterByCategory("cat1"),
			},
			expected: []Product{
				prod2,
			},
		},
		{
			name:          "with priceLessThan filter",
			connectionURI: connectionURI,
			filters: []Filter{
				GetFilterByPriceLessThan(15000),
			},
			expected: []Product{
				prod1,
				prod3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep, err := NewRepository(tt.connectionURI)
			if err != nil {
				t.Error(err)
			}
			res, err := rep.Find(tt.filters...)
			if err != nil {
				t.Error(err)
			}
			fmt.Println(res)
			assert.Equal(t, tt.expected, res)
		})
	}
}
