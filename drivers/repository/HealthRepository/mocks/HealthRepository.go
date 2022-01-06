// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	records "vaccine-app-be/drivers/records"

	mock "github.com/stretchr/testify/mock"
)

// HealthRepository is an autogenerated mock type for the HealthRepository type
type HealthRepository struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *HealthRepository) FindByEmail(ctx context.Context, email string) (records.HealthFacilitator, error) {
	ret := _m.Called(ctx, email)

	var r0 records.HealthFacilitator
	if rf, ok := ret.Get(0).(func(context.Context, string) records.HealthFacilitator); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(records.HealthFacilitator)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllHealthFacilitator provides a mock function with given fields: ctx
func (_m *HealthRepository) GetAllHealthFacilitator(ctx context.Context) ([]records.HealthFacilitator, error) {
	ret := _m.Called(ctx)

	var r0 []records.HealthFacilitator
	if rf, ok := ret.Get(0).(func(context.Context) []records.HealthFacilitator); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]records.HealthFacilitator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, healthF
func (_m *HealthRepository) Register(ctx context.Context, healthF records.HealthFacilitator) (records.HealthFacilitator, error) {
	ret := _m.Called(ctx, healthF)

	var r0 records.HealthFacilitator
	if rf, ok := ret.Get(0).(func(context.Context, records.HealthFacilitator) records.HealthFacilitator); ok {
		r0 = rf(ctx, healthF)
	} else {
		r0 = ret.Get(0).(records.HealthFacilitator)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, records.HealthFacilitator) error); ok {
		r1 = rf(ctx, healthF)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
