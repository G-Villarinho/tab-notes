// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RegisterHandlerMock is an autogenerated mock type for the RegisterHandler type
type RegisterHandlerMock struct {
	mock.Mock
}

type RegisterHandlerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *RegisterHandlerMock) EXPECT() *RegisterHandlerMock_Expecter {
	return &RegisterHandlerMock_Expecter{mock: &_m.Mock}
}

// RegisterUser provides a mock function with given fields: w, r
func (_m *RegisterHandlerMock) RegisterUser(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// RegisterHandlerMock_RegisterUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterUser'
type RegisterHandlerMock_RegisterUser_Call struct {
	*mock.Call
}

// RegisterUser is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *RegisterHandlerMock_Expecter) RegisterUser(w interface{}, r interface{}) *RegisterHandlerMock_RegisterUser_Call {
	return &RegisterHandlerMock_RegisterUser_Call{Call: _e.mock.On("RegisterUser", w, r)}
}

func (_c *RegisterHandlerMock_RegisterUser_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *RegisterHandlerMock_RegisterUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *RegisterHandlerMock_RegisterUser_Call) Return() *RegisterHandlerMock_RegisterUser_Call {
	_c.Call.Return()
	return _c
}

func (_c *RegisterHandlerMock_RegisterUser_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *RegisterHandlerMock_RegisterUser_Call {
	_c.Run(run)
	return _c
}

// NewRegisterHandlerMock creates a new instance of RegisterHandlerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRegisterHandlerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *RegisterHandlerMock {
	mock := &RegisterHandlerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
