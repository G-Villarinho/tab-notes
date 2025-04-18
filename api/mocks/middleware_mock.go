// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MiddlewareMock is an autogenerated mock type for the Middleware type
type MiddlewareMock struct {
	mock.Mock
}

type MiddlewareMock_Expecter struct {
	mock *mock.Mock
}

func (_m *MiddlewareMock) EXPECT() *MiddlewareMock_Expecter {
	return &MiddlewareMock_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MiddlewareMock) Execute(_a0 http.Handler) http.Handler {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(http.Handler) http.Handler); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MiddlewareMock_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MiddlewareMock_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 http.Handler
func (_e *MiddlewareMock_Expecter) Execute(_a0 interface{}) *MiddlewareMock_Execute_Call {
	return &MiddlewareMock_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MiddlewareMock_Execute_Call) Run(run func(_a0 http.Handler)) *MiddlewareMock_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.Handler))
	})
	return _c
}

func (_c *MiddlewareMock_Execute_Call) Return(_a0 http.Handler) *MiddlewareMock_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MiddlewareMock_Execute_Call) RunAndReturn(run func(http.Handler) http.Handler) *MiddlewareMock_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMiddlewareMock creates a new instance of MiddlewareMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMiddlewareMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *MiddlewareMock {
	mock := &MiddlewareMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
