// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/azure-devops-go-api/azuredevops/v7/security (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	security "github.com/microsoft/azure-devops-go-api/azuredevops/v7/security"
)

// MockSecurityClient is a mock of Client interface.
type MockSecurityClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityClientMockRecorder
}

// MockSecurityClientMockRecorder is the mock recorder for MockSecurityClient.
type MockSecurityClientMockRecorder struct {
	mock *MockSecurityClient
}

// NewMockSecurityClient creates a new mock instance.
func NewMockSecurityClient(ctrl *gomock.Controller) *MockSecurityClient {
	mock := &MockSecurityClient{ctrl: ctrl}
	mock.recorder = &MockSecurityClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecurityClient) EXPECT() *MockSecurityClientMockRecorder {
	return m.recorder
}

// HasPermissions mocks base method.
func (m *MockSecurityClient) HasPermissions(arg0 context.Context, arg1 security.HasPermissionsArgs) (*[]bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermissions", arg0, arg1)
	ret0, _ := ret[0].(*[]bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasPermissions indicates an expected call of HasPermissions.
func (mr *MockSecurityClientMockRecorder) HasPermissions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermissions", reflect.TypeOf((*MockSecurityClient)(nil).HasPermissions), arg0, arg1)
}

// HasPermissionsBatch mocks base method.
func (m *MockSecurityClient) HasPermissionsBatch(arg0 context.Context, arg1 security.HasPermissionsBatchArgs) (*security.PermissionEvaluationBatch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermissionsBatch", arg0, arg1)
	ret0, _ := ret[0].(*security.PermissionEvaluationBatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasPermissionsBatch indicates an expected call of HasPermissionsBatch.
func (mr *MockSecurityClientMockRecorder) HasPermissionsBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermissionsBatch", reflect.TypeOf((*MockSecurityClient)(nil).HasPermissionsBatch), arg0, arg1)
}

// QueryAccessControlLists mocks base method.
func (m *MockSecurityClient) QueryAccessControlLists(arg0 context.Context, arg1 security.QueryAccessControlListsArgs) (*[]security.AccessControlList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAccessControlLists", arg0, arg1)
	ret0, _ := ret[0].(*[]security.AccessControlList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAccessControlLists indicates an expected call of QueryAccessControlLists.
func (mr *MockSecurityClientMockRecorder) QueryAccessControlLists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAccessControlLists", reflect.TypeOf((*MockSecurityClient)(nil).QueryAccessControlLists), arg0, arg1)
}

// QuerySecurityNamespaces mocks base method.
func (m *MockSecurityClient) QuerySecurityNamespaces(arg0 context.Context, arg1 security.QuerySecurityNamespacesArgs) (*[]security.SecurityNamespaceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerySecurityNamespaces", arg0, arg1)
	ret0, _ := ret[0].(*[]security.SecurityNamespaceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySecurityNamespaces indicates an expected call of QuerySecurityNamespaces.
func (mr *MockSecurityClientMockRecorder) QuerySecurityNamespaces(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySecurityNamespaces", reflect.TypeOf((*MockSecurityClient)(nil).QuerySecurityNamespaces), arg0, arg1)
}

// RemoveAccessControlEntries mocks base method.
func (m *MockSecurityClient) RemoveAccessControlEntries(arg0 context.Context, arg1 security.RemoveAccessControlEntriesArgs) (*bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAccessControlEntries", arg0, arg1)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAccessControlEntries indicates an expected call of RemoveAccessControlEntries.
func (mr *MockSecurityClientMockRecorder) RemoveAccessControlEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAccessControlEntries", reflect.TypeOf((*MockSecurityClient)(nil).RemoveAccessControlEntries), arg0, arg1)
}

// RemoveAccessControlLists mocks base method.
func (m *MockSecurityClient) RemoveAccessControlLists(arg0 context.Context, arg1 security.RemoveAccessControlListsArgs) (*bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAccessControlLists", arg0, arg1)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAccessControlLists indicates an expected call of RemoveAccessControlLists.
func (mr *MockSecurityClientMockRecorder) RemoveAccessControlLists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAccessControlLists", reflect.TypeOf((*MockSecurityClient)(nil).RemoveAccessControlLists), arg0, arg1)
}

// RemovePermission mocks base method.
func (m *MockSecurityClient) RemovePermission(arg0 context.Context, arg1 security.RemovePermissionArgs) (*security.AccessControlEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePermission", arg0, arg1)
	ret0, _ := ret[0].(*security.AccessControlEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemovePermission indicates an expected call of RemovePermission.
func (mr *MockSecurityClientMockRecorder) RemovePermission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePermission", reflect.TypeOf((*MockSecurityClient)(nil).RemovePermission), arg0, arg1)
}

// SetAccessControlEntries mocks base method.
func (m *MockSecurityClient) SetAccessControlEntries(arg0 context.Context, arg1 security.SetAccessControlEntriesArgs) (*[]security.AccessControlEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccessControlEntries", arg0, arg1)
	ret0, _ := ret[0].(*[]security.AccessControlEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetAccessControlEntries indicates an expected call of SetAccessControlEntries.
func (mr *MockSecurityClientMockRecorder) SetAccessControlEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccessControlEntries", reflect.TypeOf((*MockSecurityClient)(nil).SetAccessControlEntries), arg0, arg1)
}

// SetAccessControlLists mocks base method.
func (m *MockSecurityClient) SetAccessControlLists(arg0 context.Context, arg1 security.SetAccessControlListsArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccessControlLists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAccessControlLists indicates an expected call of SetAccessControlLists.
func (mr *MockSecurityClientMockRecorder) SetAccessControlLists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccessControlLists", reflect.TypeOf((*MockSecurityClient)(nil).SetAccessControlLists), arg0, arg1)
}
