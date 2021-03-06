// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PasswordValidatorI is an autogenerated mock type for the PasswordValidatorI type
type PasswordValidatorI struct {
	mock.Mock
}

// ComparePassword provides a mock function with given fields: password, hash
func (_m *PasswordValidatorI) ComparePassword(password []byte, hash []byte) error {
	ret := _m.Called(password, hash)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) error); ok {
		r0 = rf(password, hash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EncryptPassword provides a mock function with given fields: password
func (_m *PasswordValidatorI) EncryptPassword(password []byte) ([]byte, error) {
	ret := _m.Called(password)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatePassword provides a mock function with given fields: s
func (_m *PasswordValidatorI) ValidatePassword(s []byte) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
