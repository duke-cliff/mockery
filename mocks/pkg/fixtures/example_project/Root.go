// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	foo "github.com/duke-cliff/mockery/v2/pkg/fixtures/example_project/foo"
)

// Root is an autogenerated mock type for the Root type
type Root struct {
	mock.Mock
}

// ReturnsFoo provides a mock function with given fields:
func (_m *Root) ReturnsFoo() (foo.Foo, error) {
	ret := _m.Called()

	var r0 foo.Foo
	if rf, ok := ret.Get(0).(func() foo.Foo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(foo.Foo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TakesBaz provides a mock function with given fields: _a0
func (_m *Root) TakesBaz(_a0 *foo.Baz) {
	_m.Called(_a0)
}
