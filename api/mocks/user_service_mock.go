// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/g-villarinho/tab-notes-api/models"
	mock "github.com/stretchr/testify/mock"
)

// UserServiceMock is an autogenerated mock type for the UserService type
type UserServiceMock struct {
	mock.Mock
}

type UserServiceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *UserServiceMock) EXPECT() *UserServiceMock_Expecter {
	return &UserServiceMock_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, name, username, email
func (_m *UserServiceMock) CreateUser(ctx context.Context, name string, username string, email string) (*models.User, error) {
	ret := _m.Called(ctx, name, username, email)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*models.User, error)); ok {
		return rf(ctx, name, username, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *models.User); ok {
		r0 = rf(ctx, name, username, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, username, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserServiceMock_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type UserServiceMock_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - username string
//   - email string
func (_e *UserServiceMock_Expecter) CreateUser(ctx interface{}, name interface{}, username interface{}, email interface{}) *UserServiceMock_CreateUser_Call {
	return &UserServiceMock_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, name, username, email)}
}

func (_c *UserServiceMock_CreateUser_Call) Run(run func(ctx context.Context, name string, username string, email string)) *UserServiceMock_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *UserServiceMock_CreateUser_Call) Return(_a0 *models.User, _a1 error) *UserServiceMock_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserServiceMock_CreateUser_Call) RunAndReturn(run func(context.Context, string, string, string) (*models.User, error)) *UserServiceMock_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetProfile provides a mock function with given fields: ctx, id
func (_m *UserServiceMock) GetProfile(ctx context.Context, id string) (*models.UserResponse, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetProfile")
	}

	var r0 *models.UserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.UserResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.UserResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserServiceMock_GetProfile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProfile'
type UserServiceMock_GetProfile_Call struct {
	*mock.Call
}

// GetProfile is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *UserServiceMock_Expecter) GetProfile(ctx interface{}, id interface{}) *UserServiceMock_GetProfile_Call {
	return &UserServiceMock_GetProfile_Call{Call: _e.mock.On("GetProfile", ctx, id)}
}

func (_c *UserServiceMock_GetProfile_Call) Run(run func(ctx context.Context, id string)) *UserServiceMock_GetProfile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserServiceMock_GetProfile_Call) Return(_a0 *models.UserResponse, _a1 error) *UserServiceMock_GetProfile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserServiceMock_GetProfile_Call) RunAndReturn(run func(context.Context, string) (*models.UserResponse, error)) *UserServiceMock_GetProfile_Call {
	_c.Call.Return(run)
	return _c
}

// GetProfileByUsername provides a mock function with given fields: ctx, username, viewerID
func (_m *UserServiceMock) GetProfileByUsername(ctx context.Context, username string, viewerID string) (*models.UserProfileResponse, error) {
	ret := _m.Called(ctx, username, viewerID)

	if len(ret) == 0 {
		panic("no return value specified for GetProfileByUsername")
	}

	var r0 *models.UserProfileResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*models.UserProfileResponse, error)); ok {
		return rf(ctx, username, viewerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.UserProfileResponse); ok {
		r0 = rf(ctx, username, viewerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserProfileResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, viewerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserServiceMock_GetProfileByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProfileByUsername'
type UserServiceMock_GetProfileByUsername_Call struct {
	*mock.Call
}

// GetProfileByUsername is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - viewerID string
func (_e *UserServiceMock_Expecter) GetProfileByUsername(ctx interface{}, username interface{}, viewerID interface{}) *UserServiceMock_GetProfileByUsername_Call {
	return &UserServiceMock_GetProfileByUsername_Call{Call: _e.mock.On("GetProfileByUsername", ctx, username, viewerID)}
}

func (_c *UserServiceMock_GetProfileByUsername_Call) Run(run func(ctx context.Context, username string, viewerID string)) *UserServiceMock_GetProfileByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *UserServiceMock_GetProfileByUsername_Call) Return(_a0 *models.UserProfileResponse, _a1 error) *UserServiceMock_GetProfileByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserServiceMock_GetProfileByUsername_Call) RunAndReturn(run func(context.Context, string, string) (*models.UserProfileResponse, error)) *UserServiceMock_GetProfileByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *UserServiceMock) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserServiceMock_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type UserServiceMock_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *UserServiceMock_Expecter) GetUserByEmail(ctx interface{}, email interface{}) *UserServiceMock_GetUserByEmail_Call {
	return &UserServiceMock_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", ctx, email)}
}

func (_c *UserServiceMock_GetUserByEmail_Call) Run(run func(ctx context.Context, email string)) *UserServiceMock_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserServiceMock_GetUserByEmail_Call) Return(_a0 *models.User, _a1 error) *UserServiceMock_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserServiceMock_GetUserByEmail_Call) RunAndReturn(run func(context.Context, string) (*models.User, error)) *UserServiceMock_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// SearchUsers provides a mock function with given fields: ctx, query
func (_m *UserServiceMock) SearchUsers(ctx context.Context, query string) ([]*models.SearchUserResponse, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for SearchUsers")
	}

	var r0 []*models.SearchUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*models.SearchUserResponse, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*models.SearchUserResponse); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.SearchUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserServiceMock_SearchUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchUsers'
type UserServiceMock_SearchUsers_Call struct {
	*mock.Call
}

// SearchUsers is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
func (_e *UserServiceMock_Expecter) SearchUsers(ctx interface{}, query interface{}) *UserServiceMock_SearchUsers_Call {
	return &UserServiceMock_SearchUsers_Call{Call: _e.mock.On("SearchUsers", ctx, query)}
}

func (_c *UserServiceMock_SearchUsers_Call) Run(run func(ctx context.Context, query string)) *UserServiceMock_SearchUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserServiceMock_SearchUsers_Call) Return(_a0 []*models.SearchUserResponse, _a1 error) *UserServiceMock_SearchUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserServiceMock_SearchUsers_Call) RunAndReturn(run func(context.Context, string) ([]*models.SearchUserResponse, error)) *UserServiceMock_SearchUsers_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserServiceMock creates a new instance of UserServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserServiceMock {
	mock := &UserServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
