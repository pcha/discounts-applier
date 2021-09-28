// Code generated by mockery 2.9.4. DO NOT EDIT.

package client

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	mongo "go.mongodb.org/mongo-driver/mongo"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

// MockMongoCollection is an autogenerated mock type for the MongoCollection type
type MockMongoCollection struct {
	mock.Mock
}

// Aggregate provides a mock function with given fields: ctx, pipeline, opts
func (_m *MockMongoCollection) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, pipeline)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.Cursor
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.AggregateOptions) *mongo.Cursor); ok {
		r0 = rf(ctx, pipeline, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Cursor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.AggregateOptions) error); ok {
		r1 = rf(ctx, pipeline, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BulkWrite provides a mock function with given fields: ctx, models, opts
func (_m *MockMongoCollection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, models)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.BulkWriteResult
	if rf, ok := ret.Get(0).(func(context.Context, []mongo.WriteModel, ...*options.BulkWriteOptions) *mongo.BulkWriteResult); ok {
		r0 = rf(ctx, models, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.BulkWriteResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []mongo.WriteModel, ...*options.BulkWriteOptions) error); ok {
		r1 = rf(ctx, models, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Clone provides a mock function with given fields: opts
func (_m *MockMongoCollection) Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.Collection
	if rf, ok := ret.Get(0).(func(...*options.CollectionOptions) *mongo.Collection); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...*options.CollectionOptions) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountDocuments provides a mock function with given fields: ctx, filter, opts
func (_m *MockMongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.CountOptions) int64); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.CountOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Database provides a mock function with given fields:
func (_m *MockMongoCollection) Database() *mongo.Database {
	ret := _m.Called()

	var r0 *mongo.Database
	if rf, ok := ret.Get(0).(func() *mongo.Database); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Database)
		}
	}

	return r0
}

// DeleteMany provides a mock function with given fields: ctx, filter, opts
func (_m *MockMongoCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.DeleteResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.DeleteOptions) *mongo.DeleteResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.DeleteResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.DeleteOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOne provides a mock function with given fields: ctx, filter, opts
func (_m *MockMongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.DeleteResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.DeleteOptions) *mongo.DeleteResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.DeleteResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.DeleteOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Distinct provides a mock function with given fields: ctx, fieldName, filter, opts
func (_m *MockMongoCollection) Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, fieldName, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, ...*options.DistinctOptions) []interface{}); ok {
		r0 = rf(ctx, fieldName, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}, ...*options.DistinctOptions) error); ok {
		r1 = rf(ctx, fieldName, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Drop provides a mock function with given fields: ctx
func (_m *MockMongoCollection) Drop(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EstimatedDocumentCount provides a mock function with given fields: ctx, opts
func (_m *MockMongoCollection) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, ...*options.EstimatedDocumentCountOptions) int64); ok {
		r0 = rf(ctx, opts...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...*options.EstimatedDocumentCountOptions) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, filter, opts
func (_m *MockMongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.Cursor
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOptions) *mongo.Cursor); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Cursor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.FindOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFive provides a mock function with given fields: filter
func (_m *MockMongoCollection) FindFive(filter interface{}) (MongoCursor, error) {
	ret := _m.Called(filter)

	var r0 MongoCursor
	if rf, ok := ret.Get(0).(func(interface{}) MongoCursor); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(MongoCursor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, filter, opts
func (_m *MockMongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOneOptions) *mongo.SingleResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.SingleResult)
		}
	}

	return r0
}

// FindOneAndDelete provides a mock function with given fields: ctx, filter, opts
func (_m *MockMongoCollection) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOneAndDeleteOptions) *mongo.SingleResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.SingleResult)
		}
	}

	return r0
}

// FindOneAndReplace provides a mock function with given fields: ctx, filter, replacement, opts
func (_m *MockMongoCollection) FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, replacement)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.FindOneAndReplaceOptions) *mongo.SingleResult); ok {
		r0 = rf(ctx, filter, replacement, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.SingleResult)
		}
	}

	return r0
}

// FindOneAndUpdate provides a mock function with given fields: ctx, filter, update, opts
func (_m *MockMongoCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, update)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.FindOneAndUpdateOptions) *mongo.SingleResult); ok {
		r0 = rf(ctx, filter, update, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.SingleResult)
		}
	}

	return r0
}

// Indexes provides a mock function with given fields:
func (_m *MockMongoCollection) Indexes() mongo.IndexView {
	ret := _m.Called()

	var r0 mongo.IndexView
	if rf, ok := ret.Get(0).(func() mongo.IndexView); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(mongo.IndexView)
	}

	return r0
}

// InsertMany provides a mock function with given fields: ctx, documents, opts
func (_m *MockMongoCollection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, documents)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.InsertManyResult
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}, ...*options.InsertManyOptions) *mongo.InsertManyResult); ok {
		r0 = rf(ctx, documents, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.InsertManyResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []interface{}, ...*options.InsertManyOptions) error); ok {
		r1 = rf(ctx, documents, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertOne provides a mock function with given fields: ctx, document, opts
func (_m *MockMongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, document)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.InsertOneResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.InsertOneOptions) *mongo.InsertOneResult); ok {
		r0 = rf(ctx, document, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.InsertOneResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.InsertOneOptions) error); ok {
		r1 = rf(ctx, document, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *MockMongoCollection) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ReplaceOne provides a mock function with given fields: ctx, filter, replacement, opts
func (_m *MockMongoCollection) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, replacement)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.ReplaceOptions) *mongo.UpdateResult); ok {
		r0 = rf(ctx, filter, replacement, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.ReplaceOptions) error); ok {
		r1 = rf(ctx, filter, replacement, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ctx, id, update, opts
func (_m *MockMongoCollection) UpdateByID(ctx context.Context, id interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id, update)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) *mongo.UpdateResult); ok {
		r0 = rf(ctx, id, update, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) error); ok {
		r1 = rf(ctx, id, update, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMany provides a mock function with given fields: ctx, filter, update, opts
func (_m *MockMongoCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, update)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) *mongo.UpdateResult); ok {
		r0 = rf(ctx, filter, update, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) error); ok {
		r1 = rf(ctx, filter, update, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOne provides a mock function with given fields: ctx, filter, update, opts
func (_m *MockMongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, update)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) *mongo.UpdateResult); ok {
		r0 = rf(ctx, filter, update, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.UpdateResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) error); ok {
		r1 = rf(ctx, filter, update, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, pipeline, opts
func (_m *MockMongoCollection) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, pipeline)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.ChangeStream
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.ChangeStreamOptions) *mongo.ChangeStream); ok {
		r0 = rf(ctx, pipeline, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.ChangeStream)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.ChangeStreamOptions) error); ok {
		r1 = rf(ctx, pipeline, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
