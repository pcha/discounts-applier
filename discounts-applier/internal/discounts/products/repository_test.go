package products

import (
	"errors"
	"testing"

	"discounts-applier/internal/discounts/products/client"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
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
			sc.On("Connect", mock.Anything).Return(nil)
			stop := sc.StartStub(tt.wantErr)
			defer stop()

			r, err := NewRepository(tt.connectionURL)
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				if assert.IsType(t, &MongoRepository{}, r) {
					mr := r.(*MongoRepository)
					assert.Same(t, sc, mr.client)
				}
			}
		})
	}
}

func TestMongoRepository_Find(t *testing.T) {
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
	connectionURI := "mongodb://uri"
	tests := []struct {
		name          string
		connectionURI string
		filters       []Filter
		filerRep      interface{}
		expected      []Product
		expectedErr   error
	}{
		{
			name:          "base",
			connectionURI: connectionURI,
			filerRep:      bson.D{},
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
			filerRep: bson.D{
				bson.E{
					Key:   "category",
					Value: "cat1",
				},
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
			filerRep: bson.D{
				bson.E{
					Key: "price",
					Value: bson.M{
						"$lte": 15000,
					},
				},
			},
			expected: []Product{
				prod1,
				prod3,
			},
		},
		{
			name:        "returns err",
			filerRep:    bson.D{},
			expectedErr: errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mCur := new(client.MockMongoCursor)
			mCur.On("All", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
				arg := args.Get(1).(*[]Product)
				*arg = tt.expected
			}).Return(nil)
			mColl := new(client.MockMongoCollection)
			mColl.On("FindFive", tt.filerRep).Return(mCur, tt.expectedErr)
			mDb := new(client.MockMongoDatabase)
			mDb.On("GetCollection").Return(mColl)
			cli := new(client.MockMongoClient)
			cli.On("GetDB").Return(mDb)

			rep := &MongoRepository{client: cli}
			res, err := rep.Find(tt.filters...)
			if tt.expectedErr != nil {
				assert.Same(t, tt.expectedErr, err)
			} else {
				assert.Equal(t, tt.expected, res)
			}
		})
	}
}

func TestMongoRepository_Clean(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "without error",
		},
		{
			name: "with error",
			err:  errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := new(client.MockMongoClient)
			mDb := new(client.MockMongoDatabase)
			cli.On("GetDB").Return(mDb)
			mColl := new(client.MockMongoCollection)
			mDb.On("GetCollection").Return(mColl)
			mColl.On("DeleteMany", mock.Anything, bson.D{}).Return(nil, tt.err)
			rep := MongoRepository{cli}

			err := rep.Clean()
			if tt.err != nil {
				assert.Same(t, tt.err, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestMongoRepository_Write(t *testing.T) {
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
		name       string
		products   []Product
		interfaces []interface{}
		err        error
	}{
		{
			name: "write OK",
			products: []Product{
				prod1,
				prod2,
				prod3,
			},
			interfaces: []interface{}{
				prod1,
				prod2,
				prod3,
			},
		},
		{
			name: "launch error",
			products: []Product{
				prod1,
				prod2,
			},
			interfaces: []interface{}{
				prod1,
				prod2,
			},
			err: errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := new(client.MockMongoClient)
			mDb := new(client.MockMongoDatabase)
			cli.On("GetDB").Return(mDb)
			mColl := new(client.MockMongoCollection)
			mDb.On("GetCollection").Return(mColl)
			mColl.On("InsertMany", mock.Anything, tt.interfaces).Return(nil, tt.err)
			rep := MongoRepository{cli}

			err := rep.Write(tt.products)
			if tt.err != nil {
				assert.Same(t, tt.err, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
