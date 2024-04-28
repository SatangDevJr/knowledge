// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AppCrypto is an autogenerated mock type for the AppCrypto type
type AppCrypto struct {
	mock.Mock
}

// GenerateFromPassword provides a mock function with given fields: pwd
func (_m *AppCrypto) GenerateFromPassword(pwd []byte) (string, error) {
	ret := _m.Called(pwd)

	var r0 string
	if rf, ok := ret.Get(0).(func([]byte) string); ok {
		r0 = rf(pwd)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(pwd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
