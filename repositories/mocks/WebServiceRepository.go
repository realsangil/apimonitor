// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/realsangil/apimonitor/models"
	mock "github.com/stretchr/testify/mock"

	rsdb "github.com/realsangil/apimonitor/pkg/rsdb"

	rsmodel "github.com/realsangil/apimonitor/pkg/rsmodels"
)

// WebServiceRepository is an autogenerated mock type for the WebServiceRepository type
type WebServiceRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: tx, src
func (_m *WebServiceRepository) Create(tx rsdb.Connection, src rsmodel.ValidatedObject) error {
	ret := _m.Called(tx, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodel.ValidatedObject) error); ok {
		r0 = rf(tx, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTable provides a mock function with given fields: tx
func (_m *WebServiceRepository) CreateTable(tx rsdb.Connection) error {
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
func (_m *WebServiceRepository) DeleteById(tx rsdb.Connection, id rsmodel.ValidatedObject) error {
	ret := _m.Called(tx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodel.ValidatedObject) error); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstOrCreate provides a mock function with given fields: tx, src
func (_m *WebServiceRepository) FirstOrCreate(tx rsdb.Connection, src rsmodel.ValidatedObject) error {
	ret := _m.Called(tx, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodel.ValidatedObject) error); ok {
		r0 = rf(tx, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllWebServicesWithTests provides a mock function with given fields: conn
func (_m *WebServiceRepository) GetAllWebServicesWithTests(conn rsdb.Connection) ([]models.WebService, error) {
	ret := _m.Called(conn)

	var r0 []models.WebService
	if rf, ok := ret.Get(0).(func(rsdb.Connection) []models.WebService); ok {
		r0 = rf(conn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.WebService)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(rsdb.Connection) error); ok {
		r1 = rf(conn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: tx, id
func (_m *WebServiceRepository) GetById(tx rsdb.Connection, id rsmodel.ValidatedObject) error {
	ret := _m.Called(tx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodel.ValidatedObject) error); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: tx, items, filter, orders
func (_m *WebServiceRepository) List(tx rsdb.Connection, items interface{}, filter rsdb.ListFilter, orders rsdb.Orders) (int, error) {
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
func (_m *WebServiceRepository) Patch(tx rsdb.Connection, src rsmodel.ValidatedObject, data rsmodel.ValidatedObject) error {
	ret := _m.Called(tx, src, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodel.ValidatedObject, rsmodel.ValidatedObject) error); ok {
		r0 = rf(tx, src, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: tx, src
func (_m *WebServiceRepository) Save(tx rsdb.Connection, src rsmodel.ValidatedObject) error {
	ret := _m.Called(tx, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(rsdb.Connection, rsmodel.ValidatedObject) error); ok {
		r0 = rf(tx, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
