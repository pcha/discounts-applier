package products

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StubMongoClient struct {
	*MockMongoClient
	originalNewFunc newClientFunc
}

type StopStub func()

func (s *StubMongoClient) StartStub(err error) StopStub {
	s.originalNewFunc = newMongoClient
	newMongoClient = func(opts ...*options.ClientOptions) (MongoClient, error) {
		if err != nil {
			return nil, err
		}
		return s, nil
	}
	return func() {
		newMongoClient = s.originalNewFunc
	}
}
