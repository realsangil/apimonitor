// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/realsangil/apimonitor/models"

import rsdb "github.com/realsangil/apimonitor/pkg/rsdb"
import rsmodels "github.com/realsangil/apimonitor/pkg/rsmodels"

// TestResultRepository is an autogenerated mock type for the TestResultRepository type
type TestResultRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: tx, src
func (_m *TestResultRepository) Create(tx rsdb.Connection, src rsmodels.ValidatedObject) error {
	ret := _m.Called(tx, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodels.ValidatedObject) error); ok {
		r0 = rf(tx, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTable provides a mock function with given fields: tx
func (_m *TestResultRepository) CreateTable(tx rsdb.Connection) error {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: tx, id
func (_m *TestResultRepository) DeleteById(tx rsdb.Connection, id rsmodels.ValidatedObject) error {
	ret := _m.Called(tx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodels.ValidatedObject) error); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstOrCreate provides a mock function with given fields: tx, src
func (_m *TestResultRepository) FirstOrCreate(tx rsdb.Connection, src rsmodels.ValidatedObject) error {
	ret := _m.Called(tx, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodels.ValidatedObject) error); ok {
		r0 = rf(tx, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetById provides a mock function with given fields: tx, id
func (_m *TestResultRepository) GetById(tx rsdb.Connection, id rsmodels.ValidatedObject) error {
	ret := _m.Called(tx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodels.ValidatedObject) error); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetResultList provides a mock function with given fields: conn, test, request
func (_m *TestResultRepository) GetResultList(conn rsdb.Connection, test *models.Test, request models.TestResultListRequest) (*rsmodels.PaginatedList, error) {
	ret := _m.Called(conn, test, request)

	var r0 *rsmodels.PaginatedList
	if rf, ok := ret.Get(0).(func(rsdb.Connection, *models.Test, models.TestResultListRequest) *rsmodels.PaginatedList); ok {
		r0 = rf(conn, test, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rsmodels.PaginatedList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(rsdb.Connection, *models.Test, models.TestResultListRequest) error); ok {
		r1 = rf(conn, test, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: tx, items, filter, orders
func (_m *TestResultRepository) List(tx rsdb.Connection, items interface{}, filter rsdb.ListFilter, orders rsdb.Orders) (int, error) {
	ret := _m.Called(tx, items, filter, orders)

	var r0 int
	if rf, ok := ret.Get(0).(func(rsdb.Connection, interface{}, rsdb.ListFilter, rsdb.Orders) int); ok {
		r0 = rf(tx, items, filter, orders)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(rsdb.Connection, interface{}, rsdb.ListFilter, rsdb.Orders) error); ok {
		r1 = rf(tx, items, filter, orders)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: tx, src, data
func (_m *TestResultRepository) Patch(tx rsdb.Connection, src rsmodels.ValidatedObject, data rsmodels.ValidatedObject) error {
	ret := _m.Called(tx, src, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodels.ValidatedObject, rsmodels.ValidatedObject) error); ok {
		r0 = rf(tx, src, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: tx, src
func (_m *TestResultRepository) Save(tx rsdb.Connection, src rsmodels.ValidatedObject) error {
	ret := _m.Called(tx, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodels.ValidatedObject) error); ok {
		r0 = rf(tx, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
