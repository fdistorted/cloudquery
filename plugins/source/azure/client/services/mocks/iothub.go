// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: IotHubClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	devices "github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-07-02/devices"
	gomock "github.com/golang/mock/gomock"
)

// MockIotHubClient is a mock of IotHubClient interface.
type MockIotHubClient struct {
	ctrl     *gomock.Controller
	recorder *MockIotHubClientMockRecorder
}

// MockIotHubClientMockRecorder is the mock recorder for MockIotHubClient.
type MockIotHubClientMockRecorder struct {
	mock *MockIotHubClient
}

// NewMockIotHubClient creates a new mock instance.
func NewMockIotHubClient(ctrl *gomock.Controller) *MockIotHubClient {
	mock := &MockIotHubClient{ctrl: ctrl}
	mock.recorder = &MockIotHubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIotHubClient) EXPECT() *MockIotHubClientMockRecorder {
	return m.recorder
}

// ListBySubscription mocks base method.
func (m *MockIotHubClient) ListBySubscription(arg0 context.Context) (devices.IotHubDescriptionListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0)
	ret0, _ := ret[0].(devices.IotHubDescriptionListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubscription indicates an expected call of ListBySubscription.
func (mr *MockIotHubClientMockRecorder) ListBySubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockIotHubClient)(nil).ListBySubscription), arg0)
}
