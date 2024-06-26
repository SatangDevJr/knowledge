// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	entity "esystem/src/pkg/entity"
	error "esystem/src/pkg/utils/error"

	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *UseCase) Add(_a0 entity.User) *error.ErrorCode {
	ret := _m.Called(_a0)

	var r0 *error.ErrorCode
	if rf, ok := ret.Get(0).(func(entity.User) *error.ErrorCode); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*error.ErrorCode)
		}
	}

	return r0
}

// ValidateLoginExternal provides a mock function with given fields: logIn
func (_m *UseCase) ValidateLoginExternal(logIn entity.Login) (*entity.User, *string, *error.ErrorCode) {
	ret := _m.Called(logIn)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(entity.Login) *entity.User); ok {
		r0 = rf(logIn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 *string
	if rf, ok := ret.Get(1).(func(entity.Login) *string); ok {
		r1 = rf(logIn)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	var r2 *error.ErrorCode
	if rf, ok := ret.Get(2).(func(entity.Login) *error.ErrorCode); ok {
		r2 = rf(logIn)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*error.ErrorCode)
		}
	}

	return r0, r1, r2
}
