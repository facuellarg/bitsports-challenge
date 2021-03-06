// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EncrypterI is an autogenerated mock type for the EncrypterI type
type EncrypterI struct {
	mock.Mock
}

// ComparePassword provides a mock function with given fields: password, hash
func (_m *EncrypterI) ComparePassword(password []byte, hash []byte) error {
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
func (_m *EncrypterI) EncryptPassword(password []byte) ([]byte, error) {
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
