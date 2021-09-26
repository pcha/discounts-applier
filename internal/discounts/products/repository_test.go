package products

import (
	"errors"
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
