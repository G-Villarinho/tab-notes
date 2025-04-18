// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// AuthMiddlewareMock is an autogenerated mock type for the AuthMiddleware type
type AuthMiddlewareMock struct {
	mock.Mock
}

type AuthMiddlewareMock_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthMiddlewareMock) EXPECT() *AuthMiddlewareMock_Expecter {
	return &AuthMiddlewareMock_Expecter{mock: &_m.Mock}
}

// Authenticated provides a mock function with given fields: next
func (_m *AuthMiddlewareMock) Authenticated(next http.HandlerFunc) http.HandlerFunc {
	ret := _m.Called(next)

	if len(ret) == 0 {
		panic("no return value specified for Authenticated")
	}

	var r0 http.HandlerFunc
	if rf, ok := ret.Get(0).(func(http.HandlerFunc) http.HandlerFunc); ok {
		r0 = rf(next)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.HandlerFunc)
		}
	}

	return r0
}

// AuthMiddlewareMock_Authenticated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Authenticated'
type AuthMiddlewareMock_Authenticated_Call struct {
	*mock.Call
}

// Authenticated is a helper method to define mock.On call
//   - next http.HandlerFunc
func (_e *AuthMiddlewareMock_Expecter) Authenticated(next interface{}) *AuthMiddlewareMock_Authenticated_Call {
	return &AuthMiddlewareMock_Authenticated_Call{Call: _e.mock.On("Authenticated", next)}
}

func (_c *AuthMiddlewareMock_Authenticated_Call) Run(run func(next http.HandlerFunc)) *AuthMiddlewareMock_Authenticated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.HandlerFunc))
	})
	return _c
}

func (_c *AuthMiddlewareMock_Authenticated_Call) Return(_a0 http.HandlerFunc) *AuthMiddlewareMock_Authenticated_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthMiddlewareMock_Authenticated_Call) RunAndReturn(run func(http.HandlerFunc) http.HandlerFunc) *AuthMiddlewareMock_Authenticated_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthMiddlewareMock creates a new instance of AuthMiddlewareMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthMiddlewareMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthMiddlewareMock {
	mock := &AuthMiddlewareMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
