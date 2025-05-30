// Code generated by mockery v2.53.3. DO NOT EDIT.

package mock_walimpls

import (
	context "context"

	message "github.com/milvus-io/milvus/pkg/v2/streaming/util/message"
	mock "github.com/stretchr/testify/mock"

	types "github.com/milvus-io/milvus/pkg/v2/streaming/util/types"

	walimpls "github.com/milvus-io/milvus/pkg/v2/streaming/walimpls"
)

// MockWALImpls is an autogenerated mock type for the WALImpls type
type MockWALImpls struct {
	mock.Mock
}

type MockWALImpls_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWALImpls) EXPECT() *MockWALImpls_Expecter {
	return &MockWALImpls_Expecter{mock: &_m.Mock}
}

// Append provides a mock function with given fields: ctx, msg
func (_m *MockWALImpls) Append(ctx context.Context, msg message.MutableMessage) (message.MessageID, error) {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for Append")
	}

	var r0 message.MessageID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, message.MutableMessage) (message.MessageID, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, message.MutableMessage) message.MessageID); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(message.MessageID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, message.MutableMessage) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWALImpls_Append_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Append'
type MockWALImpls_Append_Call struct {
	*mock.Call
}

// Append is a helper method to define mock.On call
//   - ctx context.Context
//   - msg message.MutableMessage
func (_e *MockWALImpls_Expecter) Append(ctx interface{}, msg interface{}) *MockWALImpls_Append_Call {
	return &MockWALImpls_Append_Call{Call: _e.mock.On("Append", ctx, msg)}
}

func (_c *MockWALImpls_Append_Call) Run(run func(ctx context.Context, msg message.MutableMessage)) *MockWALImpls_Append_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(message.MutableMessage))
	})
	return _c
}

func (_c *MockWALImpls_Append_Call) Return(_a0 message.MessageID, _a1 error) *MockWALImpls_Append_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWALImpls_Append_Call) RunAndReturn(run func(context.Context, message.MutableMessage) (message.MessageID, error)) *MockWALImpls_Append_Call {
	_c.Call.Return(run)
	return _c
}

// Channel provides a mock function with no fields
func (_m *MockWALImpls) Channel() types.PChannelInfo {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Channel")
	}

	var r0 types.PChannelInfo
	if rf, ok := ret.Get(0).(func() types.PChannelInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.PChannelInfo)
	}

	return r0
}

// MockWALImpls_Channel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Channel'
type MockWALImpls_Channel_Call struct {
	*mock.Call
}

// Channel is a helper method to define mock.On call
func (_e *MockWALImpls_Expecter) Channel() *MockWALImpls_Channel_Call {
	return &MockWALImpls_Channel_Call{Call: _e.mock.On("Channel")}
}

func (_c *MockWALImpls_Channel_Call) Run(run func()) *MockWALImpls_Channel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWALImpls_Channel_Call) Return(_a0 types.PChannelInfo) *MockWALImpls_Channel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWALImpls_Channel_Call) RunAndReturn(run func() types.PChannelInfo) *MockWALImpls_Channel_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with no fields
func (_m *MockWALImpls) Close() {
	_m.Called()
}

// MockWALImpls_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockWALImpls_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockWALImpls_Expecter) Close() *MockWALImpls_Close_Call {
	return &MockWALImpls_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockWALImpls_Close_Call) Run(run func()) *MockWALImpls_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWALImpls_Close_Call) Return() *MockWALImpls_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockWALImpls_Close_Call) RunAndReturn(run func()) *MockWALImpls_Close_Call {
	_c.Run(run)
	return _c
}

// Read provides a mock function with given fields: ctx, opts
func (_m *MockWALImpls) Read(ctx context.Context, opts walimpls.ReadOption) (walimpls.ScannerImpls, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 walimpls.ScannerImpls
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, walimpls.ReadOption) (walimpls.ScannerImpls, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, walimpls.ReadOption) walimpls.ScannerImpls); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(walimpls.ScannerImpls)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, walimpls.ReadOption) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWALImpls_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockWALImpls_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - opts walimpls.ReadOption
func (_e *MockWALImpls_Expecter) Read(ctx interface{}, opts interface{}) *MockWALImpls_Read_Call {
	return &MockWALImpls_Read_Call{Call: _e.mock.On("Read", ctx, opts)}
}

func (_c *MockWALImpls_Read_Call) Run(run func(ctx context.Context, opts walimpls.ReadOption)) *MockWALImpls_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(walimpls.ReadOption))
	})
	return _c
}

func (_c *MockWALImpls_Read_Call) Return(_a0 walimpls.ScannerImpls, _a1 error) *MockWALImpls_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWALImpls_Read_Call) RunAndReturn(run func(context.Context, walimpls.ReadOption) (walimpls.ScannerImpls, error)) *MockWALImpls_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Truncate provides a mock function with given fields: ctx, id
func (_m *MockWALImpls) Truncate(ctx context.Context, id message.MessageID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Truncate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, message.MessageID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWALImpls_Truncate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Truncate'
type MockWALImpls_Truncate_Call struct {
	*mock.Call
}

// Truncate is a helper method to define mock.On call
//   - ctx context.Context
//   - id message.MessageID
func (_e *MockWALImpls_Expecter) Truncate(ctx interface{}, id interface{}) *MockWALImpls_Truncate_Call {
	return &MockWALImpls_Truncate_Call{Call: _e.mock.On("Truncate", ctx, id)}
}

func (_c *MockWALImpls_Truncate_Call) Run(run func(ctx context.Context, id message.MessageID)) *MockWALImpls_Truncate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(message.MessageID))
	})
	return _c
}

func (_c *MockWALImpls_Truncate_Call) Return(_a0 error) *MockWALImpls_Truncate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWALImpls_Truncate_Call) RunAndReturn(run func(context.Context, message.MessageID) error) *MockWALImpls_Truncate_Call {
	_c.Call.Return(run)
	return _c
}

// WALName provides a mock function with no fields
func (_m *MockWALImpls) WALName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for WALName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockWALImpls_WALName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WALName'
type MockWALImpls_WALName_Call struct {
	*mock.Call
}

// WALName is a helper method to define mock.On call
func (_e *MockWALImpls_Expecter) WALName() *MockWALImpls_WALName_Call {
	return &MockWALImpls_WALName_Call{Call: _e.mock.On("WALName")}
}

func (_c *MockWALImpls_WALName_Call) Run(run func()) *MockWALImpls_WALName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWALImpls_WALName_Call) Return(_a0 string) *MockWALImpls_WALName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWALImpls_WALName_Call) RunAndReturn(run func() string) *MockWALImpls_WALName_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWALImpls creates a new instance of MockWALImpls. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWALImpls(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWALImpls {
	mock := &MockWALImpls{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
