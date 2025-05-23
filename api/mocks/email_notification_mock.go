// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// EmailNotificationMock is an autogenerated mock type for the EmailNotification type
type EmailNotificationMock struct {
	mock.Mock
}

type EmailNotificationMock_Expecter struct {
	mock *mock.Mock
}

func (_m *EmailNotificationMock) EXPECT() *EmailNotificationMock_Expecter {
	return &EmailNotificationMock_Expecter{mock: &_m.Mock}
}

// SendMagicLink provides a mock function with given fields: ctx, name, email, magicLink
func (_m *EmailNotificationMock) SendMagicLink(ctx context.Context, name string, email string, magicLink string) error {
	ret := _m.Called(ctx, name, email, magicLink)

	if len(ret) == 0 {
		panic("no return value specified for SendMagicLink")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, name, email, magicLink)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EmailNotificationMock_SendMagicLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMagicLink'
type EmailNotificationMock_SendMagicLink_Call struct {
	*mock.Call
}

// SendMagicLink is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - email string
//   - magicLink string
func (_e *EmailNotificationMock_Expecter) SendMagicLink(ctx interface{}, name interface{}, email interface{}, magicLink interface{}) *EmailNotificationMock_SendMagicLink_Call {
	return &EmailNotificationMock_SendMagicLink_Call{Call: _e.mock.On("SendMagicLink", ctx, name, email, magicLink)}
}

func (_c *EmailNotificationMock_SendMagicLink_Call) Run(run func(ctx context.Context, name string, email string, magicLink string)) *EmailNotificationMock_SendMagicLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *EmailNotificationMock_SendMagicLink_Call) Return(_a0 error) *EmailNotificationMock_SendMagicLink_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EmailNotificationMock_SendMagicLink_Call) RunAndReturn(run func(context.Context, string, string, string) error) *EmailNotificationMock_SendMagicLink_Call {
	_c.Call.Return(run)
	return _c
}

// SendWelcomeEmail provides a mock function with given fields: ctx, name, email, magicLink
func (_m *EmailNotificationMock) SendWelcomeEmail(ctx context.Context, name string, email string, magicLink string) error {
	ret := _m.Called(ctx, name, email, magicLink)

	if len(ret) == 0 {
		panic("no return value specified for SendWelcomeEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, name, email, magicLink)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EmailNotificationMock_SendWelcomeEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendWelcomeEmail'
type EmailNotificationMock_SendWelcomeEmail_Call struct {
	*mock.Call
}

// SendWelcomeEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - email string
//   - magicLink string
func (_e *EmailNotificationMock_Expecter) SendWelcomeEmail(ctx interface{}, name interface{}, email interface{}, magicLink interface{}) *EmailNotificationMock_SendWelcomeEmail_Call {
	return &EmailNotificationMock_SendWelcomeEmail_Call{Call: _e.mock.On("SendWelcomeEmail", ctx, name, email, magicLink)}
}

func (_c *EmailNotificationMock_SendWelcomeEmail_Call) Run(run func(ctx context.Context, name string, email string, magicLink string)) *EmailNotificationMock_SendWelcomeEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *EmailNotificationMock_SendWelcomeEmail_Call) Return(_a0 error) *EmailNotificationMock_SendWelcomeEmail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EmailNotificationMock_SendWelcomeEmail_Call) RunAndReturn(run func(context.Context, string, string, string) error) *EmailNotificationMock_SendWelcomeEmail_Call {
	_c.Call.Return(run)
	return _c
}

// NewEmailNotificationMock creates a new instance of EmailNotificationMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmailNotificationMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmailNotificationMock {
	mock := &EmailNotificationMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
