// Code generated by mockery v2.18.0. DO NOT EDIT.

package cachemocks

import (
	model "github.com/eminoz/go-api/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// UserCache is an autogenerated mock type for the UserCache type
type UserCache struct {
	mock.Mock
}

// DeleteUserById provides a mock function with given fields: id
func (_m *UserCache) DeleteUserById(id string) {
	_m.Called(id)
}

// GetUSerById provides a mock function with given fields: id
func (_m *UserCache) GetUSerById(id string) model.UserDto {
	ret := _m.Called(id)

	var r0 model.UserDto
	if rf, ok := ret.Get(0).(func(string) model.UserDto); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.UserDto)
	}

	return r0
}

// SaveUserByID provides a mock function with given fields: id, user
func (_m *UserCache) SaveUserByID(id string, user model.UserDto) {
	_m.Called(id, user)
}

type mockConstructorTestingTNewUserCache interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserCache creates a new instance of UserCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserCache(t mockConstructorTestingTNewUserCache) *UserCache {
	mock := &UserCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
