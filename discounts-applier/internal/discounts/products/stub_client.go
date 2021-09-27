package products

import (
	"discounts-applier/internal/discounts/products/client"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type StubMongoClient struct {
	*client.MockMongoClient
	originalNewFunc newClientFunc
}

func NewStubMongoClient() *StubMongoClient {
	return &StubMongoClient{
		MockMongoClient: new(client.MockMongoClient),
	}
}

type StopStub func()

func (s *StubMongoClient) StartStub(err error) StopStub {
	s.originalNewFunc = newMongoClient
	newMongoClient = func(opts ...*options.ClientOptions) (client.MongoClient, error) {
		if err != nil {
			return nil, err
		}
		return s, nil
	}
	return func() {
		newMongoClient = s.originalNewFunc
	}
}
