// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: KeyVaultVaultsClient,KeyVaultManagedHsmsClient,KeyVaultKeysClient,KeyVaultSecretsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	keyvault "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	keyvault0 "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	keyvault1 "github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
	gomock "github.com/golang/mock/gomock"
)

// MockKeyVaultVaultsClient is a mock of KeyVaultVaultsClient interface.
type MockKeyVaultVaultsClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeyVaultVaultsClientMockRecorder
}

// MockKeyVaultVaultsClientMockRecorder is the mock recorder for MockKeyVaultVaultsClient.
type MockKeyVaultVaultsClientMockRecorder struct {
	mock *MockKeyVaultVaultsClient
}

// NewMockKeyVaultVaultsClient creates a new mock instance.
func NewMockKeyVaultVaultsClient(ctrl *gomock.Controller) *MockKeyVaultVaultsClient {
	mock := &MockKeyVaultVaultsClient{ctrl: ctrl}
	mock.recorder = &MockKeyVaultVaultsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyVaultVaultsClient) EXPECT() *MockKeyVaultVaultsClientMockRecorder {
	return m.recorder
}

// ListBySubscription mocks base method.
func (m *MockKeyVaultVaultsClient) ListBySubscription(arg0 context.Context, arg1 *int32) (keyvault.VaultListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0, arg1)
	ret0, _ := ret[0].(keyvault.VaultListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubscription indicates an expected call of ListBySubscription.
func (mr *MockKeyVaultVaultsClientMockRecorder) ListBySubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockKeyVaultVaultsClient)(nil).ListBySubscription), arg0, arg1)
}

// MockKeyVaultManagedHsmsClient is a mock of KeyVaultManagedHsmsClient interface.
type MockKeyVaultManagedHsmsClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeyVaultManagedHsmsClientMockRecorder
}

// MockKeyVaultManagedHsmsClientMockRecorder is the mock recorder for MockKeyVaultManagedHsmsClient.
type MockKeyVaultManagedHsmsClientMockRecorder struct {
	mock *MockKeyVaultManagedHsmsClient
}

// NewMockKeyVaultManagedHsmsClient creates a new mock instance.
func NewMockKeyVaultManagedHsmsClient(ctrl *gomock.Controller) *MockKeyVaultManagedHsmsClient {
	mock := &MockKeyVaultManagedHsmsClient{ctrl: ctrl}
	mock.recorder = &MockKeyVaultManagedHsmsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyVaultManagedHsmsClient) EXPECT() *MockKeyVaultManagedHsmsClientMockRecorder {
	return m.recorder
}

// ListBySubscription mocks base method.
func (m *MockKeyVaultManagedHsmsClient) ListBySubscription(arg0 context.Context, arg1 *int32) (keyvault1.ManagedHsmListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0, arg1)
	ret0, _ := ret[0].(keyvault1.ManagedHsmListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubscription indicates an expected call of ListBySubscription.
func (mr *MockKeyVaultManagedHsmsClientMockRecorder) ListBySubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockKeyVaultManagedHsmsClient)(nil).ListBySubscription), arg0, arg1)
}

// MockKeyVaultKeysClient is a mock of KeyVaultKeysClient interface.
type MockKeyVaultKeysClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeyVaultKeysClientMockRecorder
}

// MockKeyVaultKeysClientMockRecorder is the mock recorder for MockKeyVaultKeysClient.
type MockKeyVaultKeysClientMockRecorder struct {
	mock *MockKeyVaultKeysClient
}

// NewMockKeyVaultKeysClient creates a new mock instance.
func NewMockKeyVaultKeysClient(ctrl *gomock.Controller) *MockKeyVaultKeysClient {
	mock := &MockKeyVaultKeysClient{ctrl: ctrl}
	mock.recorder = &MockKeyVaultKeysClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyVaultKeysClient) EXPECT() *MockKeyVaultKeysClientMockRecorder {
	return m.recorder
}

// GetKeys mocks base method.
func (m *MockKeyVaultKeysClient) GetKeys(arg0 context.Context, arg1 string, arg2 *int32) (keyvault0.KeyListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].(keyvault0.KeyListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeys indicates an expected call of GetKeys.
func (mr *MockKeyVaultKeysClientMockRecorder) GetKeys(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeys", reflect.TypeOf((*MockKeyVaultKeysClient)(nil).GetKeys), arg0, arg1, arg2)
}

// MockKeyVaultSecretsClient is a mock of KeyVaultSecretsClient interface.
type MockKeyVaultSecretsClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeyVaultSecretsClientMockRecorder
}

// MockKeyVaultSecretsClientMockRecorder is the mock recorder for MockKeyVaultSecretsClient.
type MockKeyVaultSecretsClientMockRecorder struct {
	mock *MockKeyVaultSecretsClient
}

// NewMockKeyVaultSecretsClient creates a new mock instance.
func NewMockKeyVaultSecretsClient(ctrl *gomock.Controller) *MockKeyVaultSecretsClient {
	mock := &MockKeyVaultSecretsClient{ctrl: ctrl}
	mock.recorder = &MockKeyVaultSecretsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyVaultSecretsClient) EXPECT() *MockKeyVaultSecretsClientMockRecorder {
	return m.recorder
}

// GetSecrets mocks base method.
func (m *MockKeyVaultSecretsClient) GetSecrets(arg0 context.Context, arg1 string, arg2 *int32) (keyvault0.SecretListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecrets", arg0, arg1, arg2)
	ret0, _ := ret[0].(keyvault0.SecretListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecrets indicates an expected call of GetSecrets.
func (mr *MockKeyVaultSecretsClientMockRecorder) GetSecrets(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecrets", reflect.TypeOf((*MockKeyVaultSecretsClient)(nil).GetSecrets), arg0, arg1, arg2)
}
