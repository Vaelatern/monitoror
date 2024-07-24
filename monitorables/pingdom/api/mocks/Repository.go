// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/Vaelatern/monitoror/monitorables/pingdom/api/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetCheck provides a mock function with given fields: checkID
func (_m *Repository) GetCheck(checkID int) (*models.Check, error) {
	ret := _m.Called(checkID)

	var r0 *models.Check
	if rf, ok := ret.Get(0).(func(int) *models.Check); ok {
		r0 = rf(checkID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Check)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(checkID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChecks provides a mock function with given fields: tags
func (_m *Repository) GetChecks(tags string) ([]models.Check, error) {
	ret := _m.Called(tags)

	var r0 []models.Check
	if rf, ok := ret.Get(0).(func(string) []models.Check); ok {
		r0 = rf(tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Check)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionCheck provides a mock function with given fields: checkID
func (_m *Repository) GetTransactionCheck(checkID int) (*models.Check, error) {
	ret := _m.Called(checkID)

	var r0 *models.Check
	if rf, ok := ret.Get(0).(func(int) *models.Check); ok {
		r0 = rf(checkID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Check)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(checkID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionChecks provides a mock function with given fields: tags
func (_m *Repository) GetTransactionChecks(tags string) ([]models.Check, error) {
	ret := _m.Called(tags)

	var r0 []models.Check
	if rf, ok := ret.Get(0).(func(string) []models.Check); ok {
		r0 = rf(tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Check)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
