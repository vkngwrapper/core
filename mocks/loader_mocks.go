// Code generated by MockGen. DO NOT EDIT.
// Source: ./iface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	core "github.com/CannibalVox/VKng/core"
	common "github.com/CannibalVox/VKng/core/common"
	core1_0 "github.com/CannibalVox/VKng/core/core1_0"
	driver "github.com/CannibalVox/VKng/core/driver"
	gomock "github.com/golang/mock/gomock"
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

// AllocateCommandBuffers mocks base method.
func (m *MockLoader) AllocateCommandBuffers(o *core1_0.CommandBufferOptions) ([]core1_0.CommandBuffer, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllocateCommandBuffers", o)
	ret0, _ := ret[0].([]core1_0.CommandBuffer)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AllocateCommandBuffers indicates an expected call of AllocateCommandBuffers.
func (mr *MockLoaderMockRecorder) AllocateCommandBuffers(o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocateCommandBuffers", reflect.TypeOf((*MockLoader)(nil).AllocateCommandBuffers), o)
}

// AllocateDescriptorSets mocks base method.
func (m *MockLoader) AllocateDescriptorSets(o *core1_0.DescriptorSetOptions) ([]core1_0.DescriptorSet, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllocateDescriptorSets", o)
	ret0, _ := ret[0].([]core1_0.DescriptorSet)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AllocateDescriptorSets indicates an expected call of AllocateDescriptorSets.
func (mr *MockLoaderMockRecorder) AllocateDescriptorSets(o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocateDescriptorSets", reflect.TypeOf((*MockLoader)(nil).AllocateDescriptorSets), o)
}

// AvailableExtensions mocks base method.
func (m *MockLoader) AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableExtensions")
	ret0, _ := ret[0].(map[string]*common.ExtensionProperties)
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
func (m *MockLoader) AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableExtensionsForLayer", layerName)
	ret0, _ := ret[0].(map[string]*common.ExtensionProperties)
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
func (m *MockLoader) AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableLayers")
	ret0, _ := ret[0].(map[string]*common.LayerProperties)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AvailableLayers indicates an expected call of AvailableLayers.
func (mr *MockLoaderMockRecorder) AvailableLayers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableLayers", reflect.TypeOf((*MockLoader)(nil).AvailableLayers))
}

// Core1_1 mocks base method.
func (m *MockLoader) Core1_1() core.Loader1_1 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Core1_1")
	ret0, _ := ret[0].(core.Loader1_1)
	return ret0
}

// Core1_1 indicates an expected call of Core1_1.
func (mr *MockLoaderMockRecorder) Core1_1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Core1_1", reflect.TypeOf((*MockLoader)(nil).Core1_1))
}

// CreateBuffer mocks base method.
func (m *MockLoader) CreateBuffer(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.BufferOptions) (core1_0.Buffer, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBuffer", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.Buffer)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateBuffer indicates an expected call of CreateBuffer.
func (mr *MockLoaderMockRecorder) CreateBuffer(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBuffer", reflect.TypeOf((*MockLoader)(nil).CreateBuffer), device, allocationCallbacks, o)
}

// CreateBufferView mocks base method.
func (m *MockLoader) CreateBufferView(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.BufferViewOptions) (core1_0.BufferView, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBufferView", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.BufferView)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateBufferView indicates an expected call of CreateBufferView.
func (mr *MockLoaderMockRecorder) CreateBufferView(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBufferView", reflect.TypeOf((*MockLoader)(nil).CreateBufferView), device, allocationCallbacks, o)
}

// CreateCommandPool mocks base method.
func (m *MockLoader) CreateCommandPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.CommandPoolOptions) (core1_0.CommandPool, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommandPool", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.CommandPool)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateCommandPool indicates an expected call of CreateCommandPool.
func (mr *MockLoaderMockRecorder) CreateCommandPool(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommandPool", reflect.TypeOf((*MockLoader)(nil).CreateCommandPool), device, allocationCallbacks, o)
}

// CreateComputePipelines mocks base method.
func (m *MockLoader) CreateComputePipelines(device core1_0.Device, pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.ComputePipelineOptions) ([]core1_0.Pipeline, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComputePipelines", device, pipelineCache, allocationCallbacks, o)
	ret0, _ := ret[0].([]core1_0.Pipeline)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateComputePipelines indicates an expected call of CreateComputePipelines.
func (mr *MockLoaderMockRecorder) CreateComputePipelines(device, pipelineCache, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComputePipelines", reflect.TypeOf((*MockLoader)(nil).CreateComputePipelines), device, pipelineCache, allocationCallbacks, o)
}

// CreateDescriptorPool mocks base method.
func (m *MockLoader) CreateDescriptorPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.DescriptorPoolOptions) (core1_0.DescriptorPool, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDescriptorPool", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.DescriptorPool)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateDescriptorPool indicates an expected call of CreateDescriptorPool.
func (mr *MockLoaderMockRecorder) CreateDescriptorPool(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDescriptorPool", reflect.TypeOf((*MockLoader)(nil).CreateDescriptorPool), device, allocationCallbacks, o)
}

// CreateDescriptorSetLayout mocks base method.
func (m *MockLoader) CreateDescriptorSetLayout(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.DescriptorSetLayoutOptions) (core1_0.DescriptorSetLayout, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDescriptorSetLayout", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.DescriptorSetLayout)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateDescriptorSetLayout indicates an expected call of CreateDescriptorSetLayout.
func (mr *MockLoaderMockRecorder) CreateDescriptorSetLayout(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDescriptorSetLayout", reflect.TypeOf((*MockLoader)(nil).CreateDescriptorSetLayout), device, allocationCallbacks, o)
}

// CreateDevice mocks base method.
func (m *MockLoader) CreateDevice(physicalDevice core1_0.PhysicalDevice, allocationCallbacks *driver.AllocationCallbacks, options *core1_0.DeviceOptions) (core1_0.Device, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevice", physicalDevice, allocationCallbacks, options)
	ret0, _ := ret[0].(core1_0.Device)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockLoaderMockRecorder) CreateDevice(physicalDevice, allocationCallbacks, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockLoader)(nil).CreateDevice), physicalDevice, allocationCallbacks, options)
}

// CreateEvent mocks base method.
func (m *MockLoader) CreateEvent(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, options *core1_0.EventOptions) (core1_0.Event, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", device, allocationCallbacks, options)
	ret0, _ := ret[0].(core1_0.Event)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockLoaderMockRecorder) CreateEvent(device, allocationCallbacks, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockLoader)(nil).CreateEvent), device, allocationCallbacks, options)
}

// CreateFence mocks base method.
func (m *MockLoader) CreateFence(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.FenceOptions) (core1_0.Fence, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFence", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.Fence)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateFence indicates an expected call of CreateFence.
func (mr *MockLoaderMockRecorder) CreateFence(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFence", reflect.TypeOf((*MockLoader)(nil).CreateFence), device, allocationCallbacks, o)
}

// CreateFrameBuffer mocks base method.
func (m *MockLoader) CreateFrameBuffer(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.FramebufferOptions) (core1_0.Framebuffer, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFrameBuffer", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.Framebuffer)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateFrameBuffer indicates an expected call of CreateFrameBuffer.
func (mr *MockLoaderMockRecorder) CreateFrameBuffer(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFrameBuffer", reflect.TypeOf((*MockLoader)(nil).CreateFrameBuffer), device, allocationCallbacks, o)
}

// CreateGraphicsPipelines mocks base method.
func (m *MockLoader) CreateGraphicsPipelines(device core1_0.Device, pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.GraphicsPipelineOptions) ([]core1_0.Pipeline, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGraphicsPipelines", device, pipelineCache, allocationCallbacks, o)
	ret0, _ := ret[0].([]core1_0.Pipeline)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateGraphicsPipelines indicates an expected call of CreateGraphicsPipelines.
func (mr *MockLoaderMockRecorder) CreateGraphicsPipelines(device, pipelineCache, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGraphicsPipelines", reflect.TypeOf((*MockLoader)(nil).CreateGraphicsPipelines), device, pipelineCache, allocationCallbacks, o)
}

// CreateImage mocks base method.
func (m *MockLoader) CreateImage(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, options *core1_0.ImageOptions) (core1_0.Image, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImage", device, allocationCallbacks, options)
	ret0, _ := ret[0].(core1_0.Image)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateImage indicates an expected call of CreateImage.
func (mr *MockLoaderMockRecorder) CreateImage(device, allocationCallbacks, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImage", reflect.TypeOf((*MockLoader)(nil).CreateImage), device, allocationCallbacks, options)
}

// CreateImageView mocks base method.
func (m *MockLoader) CreateImageView(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.ImageViewOptions) (core1_0.ImageView, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImageView", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.ImageView)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateImageView indicates an expected call of CreateImageView.
func (mr *MockLoaderMockRecorder) CreateImageView(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImageView", reflect.TypeOf((*MockLoader)(nil).CreateImageView), device, allocationCallbacks, o)
}

// CreateInstance mocks base method.
func (m *MockLoader) CreateInstance(allocationCallbacks *driver.AllocationCallbacks, options *core1_0.InstanceOptions) (core1_0.Instance, common.VkResult, error) {
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

// CreatePipelineCache mocks base method.
func (m *MockLoader) CreatePipelineCache(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.PipelineCacheOptions) (core1_0.PipelineCache, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipelineCache", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.PipelineCache)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreatePipelineCache indicates an expected call of CreatePipelineCache.
func (mr *MockLoaderMockRecorder) CreatePipelineCache(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipelineCache", reflect.TypeOf((*MockLoader)(nil).CreatePipelineCache), device, allocationCallbacks, o)
}

// CreatePipelineLayout mocks base method.
func (m *MockLoader) CreatePipelineLayout(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.PipelineLayoutOptions) (core1_0.PipelineLayout, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipelineLayout", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.PipelineLayout)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreatePipelineLayout indicates an expected call of CreatePipelineLayout.
func (mr *MockLoaderMockRecorder) CreatePipelineLayout(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipelineLayout", reflect.TypeOf((*MockLoader)(nil).CreatePipelineLayout), device, allocationCallbacks, o)
}

// CreateQueryPool mocks base method.
func (m *MockLoader) CreateQueryPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.QueryPoolOptions) (core1_0.QueryPool, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQueryPool", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.QueryPool)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateQueryPool indicates an expected call of CreateQueryPool.
func (mr *MockLoaderMockRecorder) CreateQueryPool(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQueryPool", reflect.TypeOf((*MockLoader)(nil).CreateQueryPool), device, allocationCallbacks, o)
}

// CreateRenderPass mocks base method.
func (m *MockLoader) CreateRenderPass(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.RenderPassOptions) (core1_0.RenderPass, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRenderPass", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.RenderPass)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRenderPass indicates an expected call of CreateRenderPass.
func (mr *MockLoaderMockRecorder) CreateRenderPass(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRenderPass", reflect.TypeOf((*MockLoader)(nil).CreateRenderPass), device, allocationCallbacks, o)
}

// CreateSampler mocks base method.
func (m *MockLoader) CreateSampler(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.SamplerOptions) (core1_0.Sampler, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSampler", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.Sampler)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSampler indicates an expected call of CreateSampler.
func (mr *MockLoaderMockRecorder) CreateSampler(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSampler", reflect.TypeOf((*MockLoader)(nil).CreateSampler), device, allocationCallbacks, o)
}

// CreateSemaphore mocks base method.
func (m *MockLoader) CreateSemaphore(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.SemaphoreOptions) (core1_0.Semaphore, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSemaphore", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.Semaphore)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSemaphore indicates an expected call of CreateSemaphore.
func (mr *MockLoaderMockRecorder) CreateSemaphore(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSemaphore", reflect.TypeOf((*MockLoader)(nil).CreateSemaphore), device, allocationCallbacks, o)
}

// CreateShaderModule mocks base method.
func (m *MockLoader) CreateShaderModule(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.ShaderModuleOptions) (core1_0.ShaderModule, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShaderModule", device, allocationCallbacks, o)
	ret0, _ := ret[0].(core1_0.ShaderModule)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateShaderModule indicates an expected call of CreateShaderModule.
func (mr *MockLoaderMockRecorder) CreateShaderModule(device, allocationCallbacks, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShaderModule", reflect.TypeOf((*MockLoader)(nil).CreateShaderModule), device, allocationCallbacks, o)
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

// FreeCommandBuffers mocks base method.
func (m *MockLoader) FreeCommandBuffers(buffers []core1_0.CommandBuffer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FreeCommandBuffers", buffers)
}

// FreeCommandBuffers indicates an expected call of FreeCommandBuffers.
func (mr *MockLoaderMockRecorder) FreeCommandBuffers(buffers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreeCommandBuffers", reflect.TypeOf((*MockLoader)(nil).FreeCommandBuffers), buffers)
}

// FreeDescriptorSets mocks base method.
func (m *MockLoader) FreeDescriptorSets(sets []core1_0.DescriptorSet) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FreeDescriptorSets", sets)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FreeDescriptorSets indicates an expected call of FreeDescriptorSets.
func (mr *MockLoaderMockRecorder) FreeDescriptorSets(sets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreeDescriptorSets", reflect.TypeOf((*MockLoader)(nil).FreeDescriptorSets), sets)
}

// Version mocks base method.
func (m *MockLoader) Version() common.APIVersion {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(common.APIVersion)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockLoaderMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockLoader)(nil).Version))
}

// MockLoader1_1 is a mock of Loader1_1 interface.
type MockLoader1_1 struct {
	ctrl     *gomock.Controller
	recorder *MockLoader1_1MockRecorder
}

// MockLoader1_1MockRecorder is the mock recorder for MockLoader1_1.
type MockLoader1_1MockRecorder struct {
	mock *MockLoader1_1
}

// NewMockLoader1_1 creates a new mock instance.
func NewMockLoader1_1(ctrl *gomock.Controller) *MockLoader1_1 {
	mock := &MockLoader1_1{ctrl: ctrl}
	mock.recorder = &MockLoader1_1MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoader1_1) EXPECT() *MockLoader1_1MockRecorder {
	return m.recorder
}

// SomeOneOneMethod mocks base method.
func (m *MockLoader1_1) SomeOneOneMethod() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SomeOneOneMethod")
}

// SomeOneOneMethod indicates an expected call of SomeOneOneMethod.
func (mr *MockLoader1_1MockRecorder) SomeOneOneMethod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SomeOneOneMethod", reflect.TypeOf((*MockLoader1_1)(nil).SomeOneOneMethod))
}