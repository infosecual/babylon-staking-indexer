// Code generated by mockery v2.51.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BtcInterface is an autogenerated mock type for the BtcInterface type
type BtcInterface struct {
	mock.Mock
}

// GetBlockTimestamp provides a mock function with given fields: height
func (_m *BtcInterface) GetBlockTimestamp(height uint32) (int64, error) {
	ret := _m.Called(height)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockTimestamp")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint32) (int64, error)); ok {
		return rf(height)
	}
	if rf, ok := ret.Get(0).(func(uint32) int64); ok {
		r0 = rf(height)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTipHeight provides a mock function with no fields
func (_m *BtcInterface) GetTipHeight() (uint64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTipHeight")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBtcInterface creates a new instance of BtcInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBtcInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *BtcInterface {
	mock := &BtcInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
