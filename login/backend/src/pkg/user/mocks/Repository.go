// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	entity "esystem/src/pkg/entity"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FindUserByUserName provides a mock function with given fields: username
func (_m *Repository) FindUserByUserName(username string) (*entity.User, error) {
	ret := _m.Called(username)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string) *entity.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertUser provides a mock function with given fields: _a0
func (_m *Repository) UpsertUser(_a0 entity.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
