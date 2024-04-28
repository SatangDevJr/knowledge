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

// Login provides a mock function with given fields: login
func (_m *UseCase) Login(login entity.Login) (entity.UserInfo, *error.ErrorCode) {
	ret := _m.Called(login)

	var r0 entity.UserInfo
	if rf, ok := ret.Get(0).(func(entity.Login) entity.UserInfo); ok {
		r0 = rf(login)
	} else {
		r0 = ret.Get(0).(entity.UserInfo)
	}

	var r1 *error.ErrorCode
	if rf, ok := ret.Get(1).(func(entity.Login) *error.ErrorCode); ok {
		r1 = rf(login)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error.ErrorCode)
		}
	}

	return r0, r1
}