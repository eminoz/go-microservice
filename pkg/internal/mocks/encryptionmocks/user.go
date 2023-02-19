// Code generated by mockery v2.18.0. DO NOT EDIT.

package encryptionmocks

import mock "github.com/stretchr/testify/mock"

// Encryption is an autogenerated mock type for the Encryption type
type Encryption struct {
	mock.Mock
}

// CheckPasswordHash provides a mock function with given fields: password, hash
func (_m *Encryption) CheckPasswordHash(password string, hash string) bool {
	ret := _m.Called(password, hash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(password, hash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GenerateHashPassword provides a mock function with given fields: password
func (_m *Encryption) GenerateHashPassword(password string) (string, error) {
	ret := _m.Called(password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEncryption interface {
	mock.TestingT
	Cleanup(func())
}

// NewEncryption creates a new instance of Encryption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEncryption(t mockConstructorTestingTNewEncryption) *Encryption {
	mock := &Encryption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}