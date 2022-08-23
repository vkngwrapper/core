// Code generated by MockGen. DO NOT EDIT.
// Source: ./iface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/vkngwrapper/core/v2/common"
	core1_0 "github.com/vkngwrapper/core/v2/core1_0"
	driver "github.com/vkngwrapper/core/v2/driver"
)

// MockLoader is a mock of Loader interface.
type MockLoader struct {
	ctrl     *gomock.Controller
	recorder *MockLoaderMockRecorder
}

// MockLoaderMockRecorder is the mock recorder for MockLoader.
type MockLoaderMockRecorder struct {
	mock *MockLoader
}

// NewMockLoader creates a new mock instance.
func NewMockLoader(ctrl *gomock.Controller) *MockLoader {
	mock := &MockLoader{ctrl: ctrl}
	mock.recorder = &MockLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoader) EXPECT() *MockLoaderMockRecorder {
	return m.recorder
}

// APIVersion mocks base method.
func (m *MockLoader) APIVersion() common.APIVersion {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIVersion")
	ret0, _ := ret[0].(common.APIVersion)
	return ret0
}

// APIVersion indicates an expected call of APIVersion.
func (mr *MockLoaderMockRecorder) APIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIVersion", reflect.TypeOf((*MockLoader)(nil).APIVersion))
}

// AvailableExtensions mocks base method.
func (m *MockLoader) AvailableExtensions() (map[string]*core1_0.ExtensionProperties, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableExtensions")
	ret0, _ := ret[0].(map[string]*core1_0.ExtensionProperties)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AvailableExtensions indicates an expected call of AvailableExtensions.
func (mr *MockLoaderMockRecorder) AvailableExtensions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableExtensions", reflect.TypeOf((*MockLoader)(nil).AvailableExtensions))
}

// AvailableExtensionsForLayer mocks base method.
func (m *MockLoader) AvailableExtensionsForLayer(layerName string) (map[string]*core1_0.ExtensionProperties, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableExtensionsForLayer", layerName)
	ret0, _ := ret[0].(map[string]*core1_0.ExtensionProperties)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AvailableExtensionsForLayer indicates an expected call of AvailableExtensionsForLayer.
func (mr *MockLoaderMockRecorder) AvailableExtensionsForLayer(layerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableExtensionsForLayer", reflect.TypeOf((*MockLoader)(nil).AvailableExtensionsForLayer), layerName)
}

// AvailableLayers mocks base method.
func (m *MockLoader) AvailableLayers() (map[string]*core1_0.LayerProperties, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableLayers")
	ret0, _ := ret[0].(map[string]*core1_0.LayerProperties)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AvailableLayers indicates an expected call of AvailableLayers.
func (mr *MockLoaderMockRecorder) AvailableLayers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableLayers", reflect.TypeOf((*MockLoader)(nil).AvailableLayers))
}

// CreateInstance mocks base method.
func (m *MockLoader) CreateInstance(allocationCallbacks *driver.AllocationCallbacks, options core1_0.InstanceCreateInfo) (core1_0.Instance, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstance", allocationCallbacks, options)
	ret0, _ := ret[0].(core1_0.Instance)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateInstance indicates an expected call of CreateInstance.
func (mr *MockLoaderMockRecorder) CreateInstance(allocationCallbacks, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstance", reflect.TypeOf((*MockLoader)(nil).CreateInstance), allocationCallbacks, options)
}

// Driver mocks base method.
func (m *MockLoader) Driver() driver.Driver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Driver")
	ret0, _ := ret[0].(driver.Driver)
	return ret0
}

// Driver indicates an expected call of Driver.
func (mr *MockLoaderMockRecorder) Driver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Driver", reflect.TypeOf((*MockLoader)(nil).Driver))
}
