package products

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
			sc := new(StubMongoClient)
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

//func TestMongoRepository_Find(t *testing.T) {
//	product1 := Product{
//		SKU:      "00001",
//		Name:     "Product 1",
//		Category: "cat1",
//		Price:    100000,
//	}
//	product2 := Product{
//		SKU:      "00002",
//		Name:     "Product 2",
//		Category: "cat2",
//		Price:    20000,
//	}
//	product3 := Product{
//		SKU:      "00003",
//		Name:     "Product 3",
//		Category: "Product 3",
//		Price:    300000,
//	}
//
//	type errs struct {
//		ConnectErr error
//	}
//	type args struct {
//		filter []Filter
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    []Product
//		wantErr errs
//	}{
//		{
//			name: "find without filters",
//			args: args{
//				nil,
//			},
//			want: []Product{
//				product1,
//				product2,
//				product3,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			mt := mtest.New(t,
//				mtest.NewOptions().ClientType(mtest.Mock),
//				mtest.NewOptions().DatabaseName(getDBData().Database),
//				mtest.NewOptions().CreateCollection(true),
//				mtest.NewOptions().CollectionName(getDBData().Collection),
//				)
//			defer mt.Close()
//			mt.Run("test name", func(t *mtest.T) {
//				cli := new(MockMongoClient)
//				cli.On("Connect").Return(nil)
//				cli.On("Database", getDBData().Database).Return(mt.DB)
//
//				err := mt.Client.Connect(context.Background())
//				assert.Nil(t, err)
//				//assert.Equal(t, mt.Coll, mt.DB.Collection("products"))
//
//				//
//				//rep := &MongoRepository{
//				//	cli,
//				//}
//				//res, _ := rep.Find(tt.args.filter...)
//				//assert.Equal(t, tt.wantErr, res)
//			})
//		})
//	}
//}

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
