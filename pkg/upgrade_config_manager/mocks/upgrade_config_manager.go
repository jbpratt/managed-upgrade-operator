// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/upgrade_config_manager (interfaces: UpgradeConfigManager)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUpgradeConfigManager is a mock of UpgradeConfigManager interface.
type MockUpgradeConfigManager struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeConfigManagerMockRecorder
}

// MockUpgradeConfigManagerMockRecorder is the mock recorder for MockUpgradeConfigManager.
type MockUpgradeConfigManagerMockRecorder struct {
	mock *MockUpgradeConfigManager
}

// NewMockUpgradeConfigManager creates a new mock instance.
func NewMockUpgradeConfigManager(ctrl *gomock.Controller) *MockUpgradeConfigManager {
	mock := &MockUpgradeConfigManager{ctrl: ctrl}
	mock.recorder = &MockUpgradeConfigManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeConfigManager) EXPECT() *MockUpgradeConfigManagerMockRecorder {
	return m.recorder
}

// RefreshUpgradeConfig mocks base method.
func (m *MockUpgradeConfigManager) RefreshUpgradeConfig() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshUpgradeConfig")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshUpgradeConfig indicates an expected call of RefreshUpgradeConfig.
func (mr *MockUpgradeConfigManagerMockRecorder) RefreshUpgradeConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshUpgradeConfig", reflect.TypeOf((*MockUpgradeConfigManager)(nil).RefreshUpgradeConfig))
}

// Start mocks base method.
func (m *MockUpgradeConfigManager) Start(arg0 <-chan struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start.
func (mr *MockUpgradeConfigManagerMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockUpgradeConfigManager)(nil).Start), arg0)
}
