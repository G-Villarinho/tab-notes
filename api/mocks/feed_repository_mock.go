// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/g-villarinho/tab-notes-api/models"
	mock "github.com/stretchr/testify/mock"
)

// FeedRepositoryMock is an autogenerated mock type for the FeedRepository type
type FeedRepositoryMock struct {
	mock.Mock
}

type FeedRepositoryMock_Expecter struct {
	mock *mock.Mock
}

func (_m *FeedRepositoryMock) EXPECT() *FeedRepositoryMock_Expecter {
	return &FeedRepositoryMock_Expecter{mock: &_m.Mock}
}

// GetFeed provides a mock function with given fields: ctx, userID, limit, offset
func (_m *FeedRepositoryMock) GetFeed(ctx context.Context, userID string, limit int, offset int) ([]*models.FeedPostResponse, error) {
	ret := _m.Called(ctx, userID, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetFeed")
	}

	var r0 []*models.FeedPostResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) ([]*models.FeedPostResponse, error)); ok {
		return rf(ctx, userID, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) []*models.FeedPostResponse); ok {
		r0 = rf(ctx, userID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.FeedPostResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) error); ok {
		r1 = rf(ctx, userID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FeedRepositoryMock_GetFeed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFeed'
type FeedRepositoryMock_GetFeed_Call struct {
	*mock.Call
}

// GetFeed is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - limit int
//   - offset int
func (_e *FeedRepositoryMock_Expecter) GetFeed(ctx interface{}, userID interface{}, limit interface{}, offset interface{}) *FeedRepositoryMock_GetFeed_Call {
	return &FeedRepositoryMock_GetFeed_Call{Call: _e.mock.On("GetFeed", ctx, userID, limit, offset)}
}

func (_c *FeedRepositoryMock_GetFeed_Call) Run(run func(ctx context.Context, userID string, limit int, offset int)) *FeedRepositoryMock_GetFeed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *FeedRepositoryMock_GetFeed_Call) Return(_a0 []*models.FeedPostResponse, _a1 error) *FeedRepositoryMock_GetFeed_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FeedRepositoryMock_GetFeed_Call) RunAndReturn(run func(context.Context, string, int, int) ([]*models.FeedPostResponse, error)) *FeedRepositoryMock_GetFeed_Call {
	_c.Call.Return(run)
	return _c
}

// NewFeedRepositoryMock creates a new instance of FeedRepositoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFeedRepositoryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *FeedRepositoryMock {
	mock := &FeedRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
