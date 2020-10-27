// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/upgradeconfigmanager (interfaces: UpgradeConfigManager)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/openshift/managed-upgrade-operator/pkg/apis/upgrade/v1alpha1"
	reflect "reflect"
)

// MockUpgradeConfigManager is a mock of UpgradeConfigManager interface
type MockUpgradeConfigManager struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeConfigManagerMockRecorder
}

// MockUpgradeConfigManagerMockRecorder is the mock recorder for MockUpgradeConfigManager
type MockUpgradeConfigManagerMockRecorder struct {
	mock *MockUpgradeConfigManager
}

// NewMockUpgradeConfigManager creates a new mock instance
func NewMockUpgradeConfigManager(ctrl *gomock.Controller) *MockUpgradeConfigManager {
	mock := &MockUpgradeConfigManager{ctrl: ctrl}
	mock.recorder = &MockUpgradeConfigManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpgradeConfigManager) EXPECT() *MockUpgradeConfigManagerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUpgradeConfigManager) Get() (*v1alpha1.UpgradeConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*v1alpha1.UpgradeConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUpgradeConfigManagerMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUpgradeConfigManager)(nil).Get))
}

// Refresh mocks base method
func (m *MockUpgradeConfigManager) Refresh() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh
func (mr *MockUpgradeConfigManagerMockRecorder) Refresh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockUpgradeConfigManager)(nil).Refresh))
}

// StartSync mocks base method
func (m *MockUpgradeConfigManager) StartSync(arg0 <-chan struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartSync", arg0)
}

// StartSync indicates an expected call of StartSync
func (mr *MockUpgradeConfigManagerMockRecorder) StartSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSync", reflect.TypeOf((*MockUpgradeConfigManager)(nil).StartSync), arg0)
}
