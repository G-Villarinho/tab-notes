// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// UserHandlerMock is an autogenerated mock type for the UserHandler type
type UserHandlerMock struct {
	mock.Mock
}

type UserHandlerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *UserHandlerMock) EXPECT() *UserHandlerMock_Expecter {
	return &UserHandlerMock_Expecter{mock: &_m.Mock}
}

// GetProfile provides a mock function with given fields: w, r
func (_m *UserHandlerMock) GetProfile(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// UserHandlerMock_GetProfile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProfile'
type UserHandlerMock_GetProfile_Call struct {
	*mock.Call
}

// GetProfile is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *UserHandlerMock_Expecter) GetProfile(w interface{}, r interface{}) *UserHandlerMock_GetProfile_Call {
	return &UserHandlerMock_GetProfile_Call{Call: _e.mock.On("GetProfile", w, r)}
}

func (_c *UserHandlerMock_GetProfile_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *UserHandlerMock_GetProfile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *UserHandlerMock_GetProfile_Call) Return() *UserHandlerMock_GetProfile_Call {
	_c.Call.Return()
	return _c
}

func (_c *UserHandlerMock_GetProfile_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *UserHandlerMock_GetProfile_Call {
	_c.Run(run)
	return _c
}

// GetProfileByUsername provides a mock function with given fields: w, r
func (_m *UserHandlerMock) GetProfileByUsername(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// UserHandlerMock_GetProfileByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProfileByUsername'
type UserHandlerMock_GetProfileByUsername_Call struct {
	*mock.Call
}

// GetProfileByUsername is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *UserHandlerMock_Expecter) GetProfileByUsername(w interface{}, r interface{}) *UserHandlerMock_GetProfileByUsername_Call {
	return &UserHandlerMock_GetProfileByUsername_Call{Call: _e.mock.On("GetProfileByUsername", w, r)}
}

func (_c *UserHandlerMock_GetProfileByUsername_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *UserHandlerMock_GetProfileByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *UserHandlerMock_GetProfileByUsername_Call) Return() *UserHandlerMock_GetProfileByUsername_Call {
	_c.Call.Return()
	return _c
}

func (_c *UserHandlerMock_GetProfileByUsername_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *UserHandlerMock_GetProfileByUsername_Call {
	_c.Run(run)
	return _c
}

// SearchUsers provides a mock function with given fields: w, r
func (_m *UserHandlerMock) SearchUsers(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// UserHandlerMock_SearchUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchUsers'
type UserHandlerMock_SearchUsers_Call struct {
	*mock.Call
}

// SearchUsers is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *UserHandlerMock_Expecter) SearchUsers(w interface{}, r interface{}) *UserHandlerMock_SearchUsers_Call {
	return &UserHandlerMock_SearchUsers_Call{Call: _e.mock.On("SearchUsers", w, r)}
}

func (_c *UserHandlerMock_SearchUsers_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *UserHandlerMock_SearchUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *UserHandlerMock_SearchUsers_Call) Return() *UserHandlerMock_SearchUsers_Call {
	_c.Call.Return()
	return _c
}

func (_c *UserHandlerMock_SearchUsers_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *UserHandlerMock_SearchUsers_Call {
	_c.Run(run)
	return _c
}

// UpdateUser provides a mock function with given fields: w, r
func (_m *UserHandlerMock) UpdateUser(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// UserHandlerMock_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type UserHandlerMock_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *UserHandlerMock_Expecter) UpdateUser(w interface{}, r interface{}) *UserHandlerMock_UpdateUser_Call {
	return &UserHandlerMock_UpdateUser_Call{Call: _e.mock.On("UpdateUser", w, r)}
}

func (_c *UserHandlerMock_UpdateUser_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *UserHandlerMock_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *UserHandlerMock_UpdateUser_Call) Return() *UserHandlerMock_UpdateUser_Call {
	_c.Call.Return()
	return _c
}

func (_c *UserHandlerMock_UpdateUser_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *UserHandlerMock_UpdateUser_Call {
	_c.Run(run)
	return _c
}

// NewUserHandlerMock creates a new instance of UserHandlerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserHandlerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserHandlerMock {
	mock := &UserHandlerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
