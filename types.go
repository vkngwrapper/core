package core

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
)

// Buffer represents a linear array of data, which is used for various purposes by binding it
// to a graphics or compute pipeline.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBuffer.html
type Buffer struct {
	device       loader.VkDevice
	bufferHandle loader.VkBuffer

	apiVersion common.APIVersion
}

func (b Buffer) DeviceHandle() loader.VkDevice {
	return b.device
}

func (b Buffer) Handle() loader.VkBuffer {
	return b.bufferHandle
}

func (b Buffer) APIVersion() common.APIVersion {
	return b.apiVersion
}

func InternalBuffer(device loader.VkDevice, handle loader.VkBuffer, version common.APIVersion) Buffer {
	return Buffer{
		device:       device,
		bufferHandle: handle,
		apiVersion:   version,
	}
}

// BufferView represents a contiguous range of a buffer and a specific format to be used to
// interpret the data.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferView.html
type BufferView struct {
	device           loader.VkDevice
	bufferViewHandle loader.VkBufferView

	apiVersion common.APIVersion
}

func (b BufferView) DeviceHandle() loader.VkDevice {
	return b.device
}

func (b BufferView) Handle() loader.VkBufferView {
	return b.bufferViewHandle
}

func (b BufferView) APIVersion() common.APIVersion {
	return b.apiVersion
}

func InternalBufferView(device loader.VkDevice, handle loader.VkBufferView, version common.APIVersion) BufferView {
	return BufferView{
		device:           device,
		bufferViewHandle: handle,
		apiVersion:       version,
	}
}

// CommandBuffer is an object used to record commands which can be subsequently submitted to
// a device queue for execution.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html
type CommandBuffer struct {
	device              loader.VkDevice
	commandPool         loader.VkCommandPool
	commandBufferHandle loader.VkCommandBuffer

	apiVersion common.APIVersion
}

func (b CommandBuffer) DeviceHandle() loader.VkDevice {
	return b.device
}

func (b CommandBuffer) CommandPoolHandle() loader.VkCommandPool {
	return b.commandPool
}

func (b CommandBuffer) Handle() loader.VkCommandBuffer {
	return b.commandBufferHandle
}

func (b CommandBuffer) APIVersion() common.APIVersion {
	return b.apiVersion
}

func InternalCommandBuffer(device loader.VkDevice, commandPool loader.VkCommandPool, handle loader.VkCommandBuffer, version common.APIVersion) CommandBuffer {
	return CommandBuffer{
		device:              device,
		commandPool:         commandPool,
		commandBufferHandle: handle,
		apiVersion:          version,
	}
}

// CommandPool is an opaque object that CommandBuffer memory is allocated from
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html
type CommandPool struct {
	commandPoolHandle loader.VkCommandPool
	device            loader.VkDevice
	apiVersion        common.APIVersion
}

func (p CommandPool) Handle() loader.VkCommandPool {
	return p.commandPoolHandle
}

func (p CommandPool) DeviceHandle() loader.VkDevice {
	return p.device
}

func (p CommandPool) APIVersion() common.APIVersion {
	return p.apiVersion
}

func InternalCommandPool(device loader.VkDevice, handle loader.VkCommandPool, version common.APIVersion) CommandPool {
	return CommandPool{
		device:            device,
		commandPoolHandle: handle,
		apiVersion:        version,
	}
}

// DescriptorPool maintains a pool of descriptors, from which DescriptorSet objects are allocated.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html
type DescriptorPool struct {
	descriptorPoolHandle loader.VkDescriptorPool
	device               loader.VkDevice

	apiVersion common.APIVersion
}

func (p DescriptorPool) Handle() loader.VkDescriptorPool {
	return p.descriptorPoolHandle
}

func (p DescriptorPool) DeviceHandle() loader.VkDevice {
	return p.device
}

func (p DescriptorPool) APIVersion() common.APIVersion {
	return p.apiVersion
}

func InternalDescriptorPool(device loader.VkDevice, handle loader.VkDescriptorPool, version common.APIVersion) DescriptorPool {
	return DescriptorPool{
		descriptorPoolHandle: handle,
		device:               device,
		apiVersion:           version,
	}
}

// DescriptorSetLayout is a group of zero or more descriptor bindings definitions.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html
type DescriptorSetLayout struct {
	device                    loader.VkDevice
	descriptorSetLayoutHandle loader.VkDescriptorSetLayout

	apiVersion common.APIVersion
}

func (h DescriptorSetLayout) Handle() loader.VkDescriptorSetLayout {
	return h.descriptorSetLayoutHandle
}

func (h DescriptorSetLayout) DeviceHandle() loader.VkDevice {
	return h.device
}

func (h DescriptorSetLayout) APIVersion() common.APIVersion {
	return h.apiVersion
}

func InternalDescriptorSetLayout(device loader.VkDevice, handle loader.VkDescriptorSetLayout, version common.APIVersion) DescriptorSetLayout {
	return DescriptorSetLayout{
		device:                    device,
		descriptorSetLayoutHandle: handle,
		apiVersion:                version,
	}
}

// DescriptorSet is an opaque object allocated from a DescriptorPool
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetDescriptorPool.html
type DescriptorSet struct {
	descriptorSetHandle loader.VkDescriptorSet
	device              loader.VkDevice
	descriptorPool      loader.VkDescriptorPool

	apiVersion common.APIVersion
}

func (s DescriptorSet) Handle() loader.VkDescriptorSet {
	return s.descriptorSetHandle
}

func (s DescriptorSet) APIVersion() common.APIVersion {
	return s.apiVersion
}

func (s DescriptorSet) DescriptorPoolHandle() loader.VkDescriptorPool {
	return s.descriptorPool
}

func (s DescriptorSet) DeviceHandle() loader.VkDevice {
	return s.device
}

func InternalDescriptorSet(device loader.VkDevice, descriptorPool loader.VkDescriptorPool, handle loader.VkDescriptorSet, version common.APIVersion) DescriptorSet {
	return DescriptorSet{
		device:              device,
		descriptorPool:      descriptorPool,
		descriptorSetHandle: handle,
		apiVersion:          version,
	}
}

// DeviceMemory represents a block of memory on the device
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html
type DeviceMemory struct {
	device             loader.VkDevice
	deviceMemoryHandle loader.VkDeviceMemory

	apiVersion common.APIVersion

	size int
}

func (m DeviceMemory) Handle() loader.VkDeviceMemory {
	return m.deviceMemoryHandle
}

func (m DeviceMemory) DeviceHandle() loader.VkDevice {
	return m.device
}

func (m DeviceMemory) APIVersion() common.APIVersion {
	return m.apiVersion
}

func (m DeviceMemory) Size() int {
	return m.size
}

func InternalDeviceMemory(device loader.VkDevice, handle loader.VkDeviceMemory, version common.APIVersion, size int) DeviceMemory {
	return DeviceMemory{
		device:             device,
		deviceMemoryHandle: handle,
		apiVersion:         version,
		size:               size,
	}
}

// Device represents a logical device on the host
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDevice.html
type Device struct {
	deviceHandle loader.VkDevice

	apiVersion             common.APIVersion
	activeDeviceExtensions map[string]struct{}
}

func (d Device) Handle() loader.VkDevice {
	return d.deviceHandle
}

func (d Device) APIVersion() common.APIVersion {
	return d.apiVersion
}

func (d Device) IsDeviceExtensionActive(extensionName string) bool {
	_, active := d.activeDeviceExtensions[extensionName]
	return active
}

func InternalDevice(handle loader.VkDevice, version common.APIVersion, extensions []string) Device {
	device := Device{
		deviceHandle:           handle,
		apiVersion:             version,
		activeDeviceExtensions: make(map[string]struct{}),
	}

	for _, extension := range extensions {
		device.activeDeviceExtensions[extension] = struct{}{}
	}

	return device
}

// Event is a synchronization primitive that can be used to insert fine-grained dependencies between
// commands submitted to the same queue, or between the host and a queue.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkEvent.html
type Event struct {
	eventHandle loader.VkEvent
	device      loader.VkDevice

	apiVersion common.APIVersion
}

func (e *Event) Handle() loader.VkEvent {
	return e.eventHandle
}

func (e *Event) DeviceHandle() loader.VkDevice {
	return e.device
}

func (e *Event) APIVersion() common.APIVersion {
	return e.apiVersion
}

func InternalEvent(device loader.VkDevice, handle loader.VkEvent, version common.APIVersion) Event {
	return Event{
		eventHandle: handle,
		device:      device,
		apiVersion:  version,
	}
}

// Fence is a synchronization primitive that can be used to insert a dependency from a queue to
// the host.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkFence.html
type Fence struct {
	device      loader.VkDevice
	fenceHandle loader.VkFence

	apiVersion common.APIVersion
}

func (f Fence) Handle() loader.VkFence {
	return f.fenceHandle
}

func (f Fence) DeviceHandle() loader.VkDevice {
	return f.device
}

func (f Fence) APIVersion() common.APIVersion {
	return f.apiVersion
}

func InternalFence(device loader.VkDevice, handle loader.VkFence, version common.APIVersion) Fence {
	return Fence{
		device:      device,
		fenceHandle: handle,
		apiVersion:  version,
	}
}

// Framebuffer represents a collection of specific memory attachments that a RenderPass uses
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html
type Framebuffer struct {
	device            loader.VkDevice
	framebufferHandle loader.VkFramebuffer

	apiVersion common.APIVersion
}

func (b Framebuffer) Handle() loader.VkFramebuffer {
	return b.framebufferHandle
}

func (b Framebuffer) DeviceHandle() loader.VkDevice {
	return b.device
}

func (b Framebuffer) APIVersion() common.APIVersion {
	return b.apiVersion
}

func InternalFramebuffer(device loader.VkDevice, handle loader.VkFramebuffer, version common.APIVersion) Framebuffer {
	return Framebuffer{
		device:            device,
		framebufferHandle: handle,
		apiVersion:        version,
	}
}

// Image represents multidimensional arrays of data which can be used for various purposes.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImage.html
type Image struct {
	imageHandle loader.VkImage
	device      loader.VkDevice

	apiVersion common.APIVersion
}

func (i Image) Handle() loader.VkImage {
	return i.imageHandle
}

func (i Image) DeviceHandle() loader.VkDevice {
	return i.device
}

func (i Image) APIVersion() common.APIVersion {
	return i.apiVersion
}

func InternalImage(device loader.VkDevice, handle loader.VkImage, version common.APIVersion) Image {
	return Image{
		device:      device,
		imageHandle: handle,
		apiVersion:  version,
	}
}

// ImageView represents contiguous ranges of Image subresources and contains additional metadata
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImageView.html
type ImageView struct {
	imageViewHandle loader.VkImageView
	device          loader.VkDevice

	apiVersion common.APIVersion
}

func (v ImageView) Handle() loader.VkImageView {
	return v.imageViewHandle
}

func (v ImageView) DeviceHandle() loader.VkDevice {
	return v.device
}

func (v ImageView) APIVersion() common.APIVersion {
	return v.apiVersion
}

func InternalImageView(device loader.VkDevice, handle loader.VkImageView, version common.APIVersion) ImageView {
	return ImageView{
		device:          device,
		imageViewHandle: handle,
		apiVersion:      version,
	}
}

// Instance stores per-application state for Vulkan
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkInstance.html
type Instance struct {
	instanceHandle loader.VkInstance
	maximumVersion common.APIVersion

	activeInstanceExtensions map[string]struct{}
}

func (i Instance) Handle() loader.VkInstance {
	return i.instanceHandle
}

func (i Instance) APIVersion() common.APIVersion {
	return i.maximumVersion
}

func (i Instance) IsInstanceExtensionActive(extensionName string) bool {
	_, active := i.activeInstanceExtensions[extensionName]
	return active
}

func InternalInstance(handle loader.VkInstance, version common.APIVersion, extensions []string) Instance {
	instance := Instance{
		instanceHandle:           handle,
		maximumVersion:           version,
		activeInstanceExtensions: make(map[string]struct{}),
	}

	for _, extension := range extensions {
		instance.activeInstanceExtensions[extension] = struct{}{}
	}

	return instance
}

// PhysicalDevice represents a single complete implementation of Vulkan available to the host, of which
// there are a finite number.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html
type PhysicalDevice struct {
	physicalDeviceHandle loader.VkPhysicalDevice

	instanceVersion      common.APIVersion
	maximumDeviceVersion common.APIVersion
}

func (d PhysicalDevice) Handle() loader.VkPhysicalDevice {
	return d.physicalDeviceHandle
}

func (d PhysicalDevice) DeviceAPIVersion() common.APIVersion {
	return d.maximumDeviceVersion
}

func (d PhysicalDevice) InstanceAPIVersion() common.APIVersion {
	return d.instanceVersion
}

func InternalPhysicalDevice(handle loader.VkPhysicalDevice, instanceVersion common.APIVersion, deviceVersion common.APIVersion) PhysicalDevice {
	return PhysicalDevice{
		physicalDeviceHandle: handle,
		instanceVersion:      instanceVersion,
		maximumDeviceVersion: deviceVersion,
	}
}

// Pipeline represents compute, ray tracing, and graphics pipelines
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipeline.html
type Pipeline struct {
	device         loader.VkDevice
	pipelineHandle loader.VkPipeline

	apiVersion common.APIVersion
}

func (p Pipeline) Handle() loader.VkPipeline {
	return p.pipelineHandle
}

func (p Pipeline) DeviceHandle() loader.VkDevice {
	return p.device
}

func (p Pipeline) APIVersion() common.APIVersion {
	return p.apiVersion
}

func InternalPipeline(device loader.VkDevice, handle loader.VkPipeline, version common.APIVersion) Pipeline {
	return Pipeline{
		device:         device,
		pipelineHandle: handle,
		apiVersion:     version,
	}
}

// PipelineCache allows the result of Pipeline construction to be reused between Pipeline objects
// and between runs of an application.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html
type PipelineCache struct {
	device              loader.VkDevice
	pipelineCacheHandle loader.VkPipelineCache

	apiVersion common.APIVersion
}

func (c PipelineCache) Handle() loader.VkPipelineCache {
	return c.pipelineCacheHandle
}

func (c PipelineCache) DeviceHandle() loader.VkDevice {
	return c.device
}

func (c PipelineCache) APIVersion() common.APIVersion {
	return c.apiVersion
}

func InternalPipelineCache(device loader.VkDevice, handle loader.VkPipelineCache, version common.APIVersion) PipelineCache {
	return PipelineCache{
		device:              device,
		pipelineCacheHandle: handle,
		apiVersion:          version,
	}
}

// PipelineLayout provides access to descriptor sets to Pipeline objects by combining zero or more
// descriptor sets and zero or more push constant ranges.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html
type PipelineLayout struct {
	device               loader.VkDevice
	pipelineLayoutHandle loader.VkPipelineLayout

	apiVersion common.APIVersion
}

func (l PipelineLayout) Handle() loader.VkPipelineLayout {
	return l.pipelineLayoutHandle
}

func (l PipelineLayout) DeviceHandle() loader.VkDevice {
	return l.device
}

func (l PipelineLayout) APIVersion() common.APIVersion {
	return l.apiVersion
}

func InternalPipelineLayout(device loader.VkDevice, handle loader.VkPipelineLayout, version common.APIVersion) PipelineLayout {
	return PipelineLayout{
		device:               device,
		pipelineLayoutHandle: handle,
		apiVersion:           version,
	}
}

// QueryPool is a collection of a specific number of queries of a particular type.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html
type QueryPool struct {
	queryPoolHandle loader.VkQueryPool
	device          loader.VkDevice

	apiVersion common.APIVersion
}

func (p QueryPool) Handle() loader.VkQueryPool {
	return p.queryPoolHandle
}

func (p QueryPool) DeviceHandle() loader.VkDevice {
	return p.device
}

func (p QueryPool) APIVersion() common.APIVersion {
	return p.apiVersion
}

func InternalQueryPool(device loader.VkDevice, handle loader.VkQueryPool, version common.APIVersion) QueryPool {
	return QueryPool{
		device:          device,
		queryPoolHandle: handle,
		apiVersion:      version,
	}
}

// Queue represents a Device resource on which work is performed
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkQueue.html
type Queue struct {
	device      loader.VkDevice
	queueHandle loader.VkQueue

	apiVersion common.APIVersion
}

func (q Queue) Handle() loader.VkQueue {
	return q.queueHandle
}

func (q Queue) DeviceHandle() loader.VkDevice {
	return q.device
}

func (q Queue) APIVersion() common.APIVersion {
	return q.apiVersion
}

func InternalQueue(device loader.VkDevice, handle loader.VkQueue, version common.APIVersion) Queue {
	return Queue{
		device:      device,
		queueHandle: handle,
		apiVersion:  version,
	}
}

// RenderPass represents a collection of attachments, subpasses, and dependencies between the subpasses
// and describes how the attachments are used over the course of the subpasses
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html
type RenderPass struct {
	device           loader.VkDevice
	renderPassHandle loader.VkRenderPass

	apiVersion common.APIVersion
}

func (p RenderPass) Handle() loader.VkRenderPass {
	return p.renderPassHandle
}

func (p RenderPass) DeviceHandle() loader.VkDevice {
	return p.device
}

func (p RenderPass) APIVersion() common.APIVersion {
	return p.apiVersion
}

func InternalRenderPass(device loader.VkDevice, handle loader.VkRenderPass, version common.APIVersion) RenderPass {
	return RenderPass{
		device:           device,
		renderPassHandle: handle,
		apiVersion:       version,
	}
}

// Sampler represents the state of an Image sampler, which is used by the implementation to read Image data
// and apply filtering and other transformations for the shader.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSampler.html
type Sampler struct {
	device        loader.VkDevice
	samplerHandle loader.VkSampler

	apiVersion common.APIVersion
}

func (s Sampler) Handle() loader.VkSampler {
	return s.samplerHandle
}

func (s Sampler) DeviceHandle() loader.VkDevice {
	return s.device
}

func (s Sampler) APIVersion() common.APIVersion {
	return s.apiVersion
}

func InternalSampler(device loader.VkDevice, handle loader.VkSampler, version common.APIVersion) Sampler {
	return Sampler{
		device:        device,
		samplerHandle: handle,
		apiVersion:    version,
	}
}

// Semaphore is a synchronization primitive that can be used to insert a dependency between Queue operations
// or between a Queue operation and the host.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html
type Semaphore struct {
	device          loader.VkDevice
	semaphoreHandle loader.VkSemaphore

	apiVersion common.APIVersion
}

func (s Semaphore) Handle() loader.VkSemaphore {
	return s.semaphoreHandle
}

func (s Semaphore) DeviceHandle() loader.VkDevice {
	return s.device
}

func (s Semaphore) APIVersion() common.APIVersion {
	return s.apiVersion
}

func InternalSemaphore(device loader.VkDevice, handle loader.VkSemaphore, version common.APIVersion) Semaphore {
	return Semaphore{
		device:          device,
		semaphoreHandle: handle,
		apiVersion:      version,
	}
}

// ShaderModule objects contain shader code and one or more entry points.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html
type ShaderModule struct {
	device             loader.VkDevice
	shaderModuleHandle loader.VkShaderModule

	apiVersion common.APIVersion
}

func (m ShaderModule) Handle() loader.VkShaderModule {
	return m.shaderModuleHandle
}

func (m ShaderModule) DeviceHandle() loader.VkDevice {
	return m.device
}

func (m ShaderModule) APIVersion() common.APIVersion {
	return m.apiVersion
}

func InternalShaderModule(device loader.VkDevice, handle loader.VkShaderModule, version common.APIVersion) ShaderModule {
	return ShaderModule{
		device:             device,
		shaderModuleHandle: handle,
		apiVersion:         version,
	}
}

// SamplerYcbcrConversion is an opaque representation of a device-specific sampler YCbCr conversion
// description.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html
type SamplerYcbcrConversion struct {
	device      loader.VkDevice
	ycbcrHandle loader.VkSamplerYcbcrConversion

	apiVersion common.APIVersion
}

func (y SamplerYcbcrConversion) Handle() loader.VkSamplerYcbcrConversion {
	return y.ycbcrHandle
}

func (y SamplerYcbcrConversion) DeviceHandle() loader.VkDevice {
	return y.device
}

func (y SamplerYcbcrConversion) APIVersion() common.APIVersion {
	return y.apiVersion
}

func InternalSamplerYcbcrConversion(device loader.VkDevice, handle loader.VkSamplerYcbcrConversion, version common.APIVersion) SamplerYcbcrConversion {
	return SamplerYcbcrConversion{
		device:      device,
		ycbcrHandle: handle,
		apiVersion:  version,
	}
}

// DescriptorUpdateTemplate specifies a mapping from descriptor update information in host memory to
// descriptors in a DescriptorSet
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html
type DescriptorUpdateTemplate struct {
	device                   loader.VkDevice
	descriptorTemplateHandle loader.VkDescriptorUpdateTemplate

	apiVersion common.APIVersion
}

func (t DescriptorUpdateTemplate) Handle() loader.VkDescriptorUpdateTemplate {
	return t.descriptorTemplateHandle
}

func (t DescriptorUpdateTemplate) DeviceHandle() loader.VkDevice {
	return t.device
}

func (t DescriptorUpdateTemplate) APIVersion() common.APIVersion {
	return t.apiVersion
}

func InternalDescriptorUpdateTemplate(device loader.VkDevice, handle loader.VkDescriptorUpdateTemplate, version common.APIVersion) DescriptorUpdateTemplate {
	return DescriptorUpdateTemplate{
		device:                   device,
		descriptorTemplateHandle: handle,
		apiVersion:               version,
	}
}
