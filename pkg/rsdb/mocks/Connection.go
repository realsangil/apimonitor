// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	gorm "github.com/jinzhu/gorm"
	mock "github.com/stretchr/testify/mock"

	rsdb "github.com/realsangil/apimonitor/pkg/rsdb"
)

// Connection is an autogenerated mock type for the Connection type
type Connection struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *Connection) Begin() (rsdb.Connection, error) {
	ret := _m.Called()

	var r0 rsdb.Connection
	if rf, ok := ret.Get(0).(func() rsdb.Connection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rsdb.Connection)
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

// Close provides a mock function with given fields:
func (_m *Connection) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Commit provides a mock function with given fields:
func (_m *Connection) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Conn provides a mock function with given fields:
func (_m *Connection) Conn() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *Connection) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
