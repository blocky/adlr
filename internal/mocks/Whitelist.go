// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Whitelist is an autogenerated mock type for the Whitelist type
type Whitelist struct {
	mock.Mock
}

// Find provides a mock function with given fields: _a0
func (_m *Whitelist) Find(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}