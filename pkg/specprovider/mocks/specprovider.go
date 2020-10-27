// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/specprovider (interfaces: SpecProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/openshift/managed-upgrade-operator/pkg/apis/upgrade/v1alpha1"
	reflect "reflect"
)

// MockSpecProvider is a mock of SpecProvider interface
type MockSpecProvider struct {
	ctrl     *gomock.Controller
	recorder *MockSpecProviderMockRecorder
}

// MockSpecProviderMockRecorder is the mock recorder for MockSpecProvider
type MockSpecProviderMockRecorder struct {
	mock *MockSpecProvider
}

// NewMockSpecProvider creates a new mock instance
func NewMockSpecProvider(ctrl *gomock.Controller) *MockSpecProvider {
	mock := &MockSpecProvider{ctrl: ctrl}
	mock.recorder = &MockSpecProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpecProvider) EXPECT() *MockSpecProviderMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockSpecProvider) Get() ([]v1alpha1.UpgradeConfigSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].([]v1alpha1.UpgradeConfigSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSpecProviderMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSpecProvider)(nil).Get))
}
