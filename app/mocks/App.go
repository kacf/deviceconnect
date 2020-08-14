// Copyright 2020 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/mendersoftware/deviceconnect/model"
	mock "github.com/stretchr/testify/mock"
)

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// DeleteDevice provides a mock function with given fields: ctx, tenantID, deviceID
func (_m *App) DeleteDevice(ctx context.Context, tenantID string, deviceID string) error {
	ret := _m.Called(ctx, tenantID, deviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenantID, deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HealthCheck provides a mock function with given fields: ctx
func (_m *App) HealthCheck(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvisionDevice provides a mock function with given fields: ctx, tenantID, device
func (_m *App) ProvisionDevice(ctx context.Context, tenantID string, device *model.Device) error {
	ret := _m.Called(ctx, tenantID, device)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Device) error); ok {
		r0 = rf(ctx, tenantID, device)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvisionTenant provides a mock function with given fields: ctx, tenant
func (_m *App) ProvisionTenant(ctx context.Context, tenant *model.Tenant) error {
	ret := _m.Called(ctx, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Tenant) error); ok {
		r0 = rf(ctx, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}