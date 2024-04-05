// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ServerInterface is an autogenerated mock type for the ServerInterface type
type ServerInterface struct {
	mock.Mock
}

// Start provides a mock function with given fields: ctx, port
func (_m *ServerInterface) Start(ctx context.Context, port int) error {
	ret := _m.Called(ctx, port)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServerInterface creates a new instance of ServerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServerInterface {
	mock := &ServerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
