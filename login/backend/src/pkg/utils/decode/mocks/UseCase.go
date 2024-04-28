// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// DecodeJWT provides a mock function with given fields: accessToken
func (_m *UseCase) DecodeJWT(accessToken string) (interface{}, error) {
	ret := _m.Called(accessToken)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(accessToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}