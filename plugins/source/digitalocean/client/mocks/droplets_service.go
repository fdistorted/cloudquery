// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/digitalocean/client (interfaces: DropletsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockDropletsService is a mock of DropletsService interface.
type MockDropletsService struct {
	ctrl     *gomock.Controller
	recorder *MockDropletsServiceMockRecorder
}

// MockDropletsServiceMockRecorder is the mock recorder for MockDropletsService.
type MockDropletsServiceMockRecorder struct {
	mock *MockDropletsService
}

// NewMockDropletsService creates a new mock instance.
func NewMockDropletsService(ctrl *gomock.Controller) *MockDropletsService {
	mock := &MockDropletsService{ctrl: ctrl}
	mock.recorder = &MockDropletsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletsService) EXPECT() *MockDropletsServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockDropletsService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Droplet, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]godo.Droplet)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockDropletsServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDropletsService)(nil).List), arg0, arg1)
}

// Neighbors mocks base method.
func (m *MockDropletsService) Neighbors(arg0 context.Context, arg1 int) ([]godo.Droplet, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Neighbors", arg0, arg1)
	ret0, _ := ret[0].([]godo.Droplet)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Neighbors indicates an expected call of Neighbors.
func (mr *MockDropletsServiceMockRecorder) Neighbors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Neighbors", reflect.TypeOf((*MockDropletsService)(nil).Neighbors), arg0, arg1)
}
