// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/shekhertrivedi/tweet-service/model"
	mock "github.com/stretchr/testify/mock"
)

// TweetRepo is an autogenerated mock type for the TweetRepo type
type TweetRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *TweetRepo) Create(_a0 context.Context, _a1 *model.Tweet) (*model.Tweet, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Tweet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Tweet) (*model.Tweet, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Tweet) *model.Tweet); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Tweet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Tweet) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *TweetRepo) Delete(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *TweetRepo) Get(_a0 context.Context, _a1 string) (*model.Tweet, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Tweet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Tweet, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Tweet); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Tweet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *TweetRepo) GetAll(ctx context.Context) ([]*model.Tweet, error) {
	ret := _m.Called(ctx)

	var r0 []*model.Tweet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*model.Tweet, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*model.Tweet); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Tweet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *TweetRepo) Update(_a0 context.Context, _a1 *model.Tweet) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Tweet) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTweetRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewTweetRepo creates a new instance of TweetRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTweetRepo(t mockConstructorTestingTNewTweetRepo) *TweetRepo {
	mock := &TweetRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}