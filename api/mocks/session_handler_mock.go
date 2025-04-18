// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// SessionHandlerMock is an autogenerated mock type for the SessionHandler type
type SessionHandlerMock struct {
	mock.Mock
}

type SessionHandlerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *SessionHandlerMock) EXPECT() *SessionHandlerMock_Expecter {
	return &SessionHandlerMock_Expecter{mock: &_m.Mock}
}

// GetUserSessions provides a mock function with given fields: w, r
func (_m *SessionHandlerMock) GetUserSessions(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// SessionHandlerMock_GetUserSessions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserSessions'
type SessionHandlerMock_GetUserSessions_Call struct {
	*mock.Call
}

// GetUserSessions is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *SessionHandlerMock_Expecter) GetUserSessions(w interface{}, r interface{}) *SessionHandlerMock_GetUserSessions_Call {
	return &SessionHandlerMock_GetUserSessions_Call{Call: _e.mock.On("GetUserSessions", w, r)}
}

func (_c *SessionHandlerMock_GetUserSessions_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *SessionHandlerMock_GetUserSessions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *SessionHandlerMock_GetUserSessions_Call) Return() *SessionHandlerMock_GetUserSessions_Call {
	_c.Call.Return()
	return _c
}

func (_c *SessionHandlerMock_GetUserSessions_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *SessionHandlerMock_GetUserSessions_Call {
	_c.Run(run)
	return _c
}

// RevokeAllSessions provides a mock function with given fields: w, r
func (_m *SessionHandlerMock) RevokeAllSessions(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// SessionHandlerMock_RevokeAllSessions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevokeAllSessions'
type SessionHandlerMock_RevokeAllSessions_Call struct {
	*mock.Call
}

// RevokeAllSessions is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *SessionHandlerMock_Expecter) RevokeAllSessions(w interface{}, r interface{}) *SessionHandlerMock_RevokeAllSessions_Call {
	return &SessionHandlerMock_RevokeAllSessions_Call{Call: _e.mock.On("RevokeAllSessions", w, r)}
}

func (_c *SessionHandlerMock_RevokeAllSessions_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *SessionHandlerMock_RevokeAllSessions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *SessionHandlerMock_RevokeAllSessions_Call) Return() *SessionHandlerMock_RevokeAllSessions_Call {
	_c.Call.Return()
	return _c
}

func (_c *SessionHandlerMock_RevokeAllSessions_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *SessionHandlerMock_RevokeAllSessions_Call {
	_c.Run(run)
	return _c
}

// RevokeSession provides a mock function with given fields: w, r
func (_m *SessionHandlerMock) RevokeSession(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// SessionHandlerMock_RevokeSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevokeSession'
type SessionHandlerMock_RevokeSession_Call struct {
	*mock.Call
}

// RevokeSession is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *SessionHandlerMock_Expecter) RevokeSession(w interface{}, r interface{}) *SessionHandlerMock_RevokeSession_Call {
	return &SessionHandlerMock_RevokeSession_Call{Call: _e.mock.On("RevokeSession", w, r)}
}

func (_c *SessionHandlerMock_RevokeSession_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *SessionHandlerMock_RevokeSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *SessionHandlerMock_RevokeSession_Call) Return() *SessionHandlerMock_RevokeSession_Call {
	_c.Call.Return()
	return _c
}

func (_c *SessionHandlerMock_RevokeSession_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *SessionHandlerMock_RevokeSession_Call {
	_c.Run(run)
	return _c
}

// NewSessionHandlerMock creates a new instance of SessionHandlerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSessionHandlerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *SessionHandlerMock {
	mock := &SessionHandlerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
