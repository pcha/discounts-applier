// Code generated by mockery 2.9.4. DO NOT EDIT.

package products

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	mongo "go.mongodb.org/mongo-driver/mongo"

	options "go.mongodb.org/mongo-driver/mongo/options"

	readpref "go.mongodb.org/mongo-driver/mongo/readpref"
)

// MockMongoClient is an autogenerated mock type for the MongoClient type
type MockMongoClient struct {
	mock.Mock
}

// Connect provides a mock function with given fields: ctx
func (_m *MockMongoClient) Connect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database provides a mock function with given fields: name, opts
func (_m *MockMongoClient) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.Database
	if rf, ok := ret.Get(0).(func(string, ...*options.DatabaseOptions) *mongo.Database); ok {
		r0 = rf(name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Database)
		}
	}

	return r0
}

// Disconnect provides a mock function with given fields: ctx
func (_m *MockMongoClient) Disconnect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListDatabaseNames provides a mock function with given fields: ctx, filter, opts
func (_m *MockMongoClient) ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.ListDatabasesOptions) []string); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.ListDatabasesOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDatabases provides a mock function with given fields: ctx, filter, opts
func (_m *MockMongoClient) ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongo.ListDatabasesResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.ListDatabasesOptions) mongo.ListDatabasesResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		r0 = ret.Get(0).(mongo.ListDatabasesResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.ListDatabasesOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NumberSessionsInProgress provides a mock function with given fields:
func (_m *MockMongoClient) NumberSessionsInProgress() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Ping provides a mock function with given fields: ctx, rp
func (_m *MockMongoClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	ret := _m.Called(ctx, rp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *readpref.ReadPref) error); ok {
		r0 = rf(ctx, rp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartSession provides a mock function with given fields: opts
func (_m *MockMongoClient) StartSession(opts ...*options.SessionOptions) (mongo.Session, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongo.Session
	if rf, ok := ret.Get(0).(func(...*options.SessionOptions) mongo.Session); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...*options.SessionOptions) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseSession provides a mock function with given fields: ctx, fn
func (_m *MockMongoClient) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(mongo.SessionContext) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UseSessionWithOptions provides a mock function with given fields: ctx, opts, fn
func (_m *MockMongoClient) UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error {
	ret := _m.Called(ctx, opts, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *options.SessionOptions, func(mongo.SessionContext) error) error); ok {
		r0 = rf(ctx, opts, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Watch provides a mock function with given fields: ctx, pipeline, opts
func (_m *MockMongoClient) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
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
