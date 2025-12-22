package types

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// Buffer represents a linear array of data, which is used for various purposes by binding it
// to a graphics or compute pipeline.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBuffer.html
type Buffer struct {
	device       driver.VkDevice
	bufferHandle driver.VkBuffer

	apiVersion common.APIVersion
}

func (b Buffer) DeviceHandle() driver.VkDevice {
	return b.device
}

func (b Buffer) Handle() driver.VkBuffer {
	return b.bufferHandle
}

func (b Buffer) APIVersion() common.APIVersion {
	return b.apiVersion
}

func InternalBuffer(device driver.VkDevice, handle driver.VkBuffer, version common.APIVersion) Buffer {
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
	device           driver.VkDevice
	bufferViewHandle driver.VkBufferView

	apiVersion common.APIVersion
}

func (b BufferView) DeviceHandle() driver.VkDevice {
	return b.device
}

func (b BufferView) Handle() driver.VkBufferView {
	return b.bufferViewHandle
}

func (b BufferView) APIVersion() common.APIVersion {
	return b.apiVersion
}

func InternalBufferView(device driver.VkDevice, handle driver.VkBufferView, version common.APIVersion) BufferView {
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
	device              driver.VkDevice
	commandPool         driver.VkCommandPool
	commandBufferHandle driver.VkCommandBuffer

	apiVersion common.APIVersion
}

func (b CommandBuffer) DeviceHandle() driver.VkDevice {
	return b.device
}

func (b CommandBuffer) CommandPoolHandle() driver.VkCommandPool {
	return b.commandPool
}

func (b CommandBuffer) Handle() driver.VkCommandBuffer {
	return b.commandBufferHandle
}

func (b CommandBuffer) APIVersion() common.APIVersion {
	return b.apiVersion
}

func InternalCommandBuffer(device driver.VkDevice, commandPool driver.VkCommandPool, handle driver.VkCommandBuffer, version common.APIVersion) CommandBuffer {
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
	commandPoolHandle driver.VkCommandPool
	device            driver.VkDevice
	apiVersion        common.APIVersion
}

func (p CommandPool) Handle() driver.VkCommandPool {
	return p.commandPoolHandle
}

func (p CommandPool) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p CommandPool) APIVersion() common.APIVersion {
	return p.apiVersion
}

func InternalCommandPool(device driver.VkDevice, handle driver.VkCommandPool, version common.APIVersion) CommandPool {
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
	descriptorPoolHandle driver.VkDescriptorPool
	device               driver.VkDevice

	apiVersion common.APIVersion
}

func (p DescriptorPool) Handle() driver.VkDescriptorPool {
	return p.descriptorPoolHandle
}

func (p DescriptorPool) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p DescriptorPool) APIVersion() common.APIVersion {
	return p.apiVersion
}

func InternalDescriptorPool(device driver.VkDevice, handle driver.VkDescriptorPool, version common.APIVersion) DescriptorPool {
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
	device                    driver.VkDevice
	descriptorSetLayoutHandle driver.VkDescriptorSetLayout

	apiVersion common.APIVersion
}

func (h DescriptorSetLayout) Handle() driver.VkDescriptorSetLayout {
	return h.descriptorSetLayoutHandle
}

func (h DescriptorSetLayout) DeviceHandle() driver.VkDevice {
	return h.device
}

func (h DescriptorSetLayout) APIVersion() common.APIVersion {
	return h.apiVersion
}

func InternalDescriptorSetLayout(device driver.VkDevice, handle driver.VkDescriptorSetLayout, version common.APIVersion) DescriptorSetLayout {
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
	descriptorSetHandle driver.VkDescriptorSet
	device              driver.VkDevice
	descriptorPool      driver.VkDescriptorPool

	apiVersion common.APIVersion
}

func (s DescriptorSet) Handle() driver.VkDescriptorSet {
	return s.descriptorSetHandle
}

func (s DescriptorSet) APIVersion() common.APIVersion {
	return s.apiVersion
}

func (s DescriptorSet) DescriptorPoolHandle() driver.VkDescriptorPool {
	return s.descriptorPool
}

func (s DescriptorSet) DeviceHandle() driver.VkDevice {
	return s.device
}

func InternalDescriptorSet(device driver.VkDevice, descriptorPool driver.VkDescriptorPool, handle driver.VkDescriptorSet, version common.APIVersion) DescriptorSet {
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
	device             driver.VkDevice
	deviceMemoryHandle driver.VkDeviceMemory

	apiVersion common.APIVersion

	size int
}

func (m DeviceMemory) Handle() driver.VkDeviceMemory {
	return m.deviceMemoryHandle
}

func (m DeviceMemory) DeviceHandle() driver.VkDevice {
	return m.device
}

func (m DeviceMemory) APIVersion() common.APIVersion {
	return m.apiVersion
}

func (m DeviceMemory) Size() int {
	return m.size
}

func InternalDeviceMemory(device driver.VkDevice, handle driver.VkDeviceMemory, version common.APIVersion, size int) DeviceMemory {
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
	deviceHandle driver.VkDevice

	apiVersion             common.APIVersion
	activeDeviceExtensions map[string]struct{}
}

func (d Device) Handle() driver.VkDevice {
	return d.deviceHandle
}

func (d Device) APIVersion() common.APIVersion {
	return d.apiVersion
}

func (d Device) IsDeviceExtensionActive(extensionName string) bool {
	_, active := d.activeDeviceExtensions[extensionName]
	return active
}

func InternalDevice(handle driver.VkDevice, version common.APIVersion, extensions []string) Device {
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
	eventHandle driver.VkEvent
	device      driver.VkDevice

	apiVersion common.APIVersion
}

func (e *Event) Handle() driver.VkEvent {
	return e.eventHandle
}

func (e *Event) DeviceHandle() driver.VkDevice {
	return e.device
}

func (e *Event) APIVersion() common.APIVersion {
	return e.apiVersion
}

func InternalEvent(device driver.VkDevice, handle driver.VkEvent, version common.APIVersion) Event {
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
	device      driver.VkDevice
	fenceHandle driver.VkFence

	apiVersion common.APIVersion
}

func (f Fence) Handle() driver.VkFence {
	return f.fenceHandle
}

func (f Fence) DeviceHandle() driver.VkDevice {
	return f.device
}

func (f Fence) APIVersion() common.APIVersion {
	return f.apiVersion
}

func InternalFence(device driver.VkDevice, handle driver.VkFence, version common.APIVersion) Fence {
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
	device            driver.VkDevice
	framebufferHandle driver.VkFramebuffer

	apiVersion common.APIVersion
}

func (b Framebuffer) Handle() driver.VkFramebuffer {
	return b.framebufferHandle
}

func (b Framebuffer) DeviceHandle() driver.VkDevice {
	return b.device
}

func (b Framebuffer) APIVersion() common.APIVersion {
	return b.apiVersion
}

func InternalFramebuffer(device driver.VkDevice, handle driver.VkFramebuffer, version common.APIVersion) Framebuffer {
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
	imageHandle driver.VkImage
	device      driver.VkDevice

	apiVersion common.APIVersion
}

func (i Image) Handle() driver.VkImage {
	return i.imageHandle
}

func (i Image) DeviceHandle() driver.VkDevice {
	return i.device
}

func (i Image) APIVersion() common.APIVersion {
	return i.apiVersion
}

func InternalImage(device driver.VkDevice, handle driver.VkImage, version common.APIVersion) Image {
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
	imageViewHandle driver.VkImageView
	device          driver.VkDevice

	apiVersion common.APIVersion
}

func (v ImageView) Handle() driver.VkImageView {
	return v.imageViewHandle
}

func (v ImageView) DeviceHandle() driver.VkDevice {
	return v.device
}

func (v ImageView) APIVersion() common.APIVersion {
	return v.apiVersion
}

func InternalImageView(device driver.VkDevice, handle driver.VkImageView, version common.APIVersion) ImageView {
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
	instanceHandle driver.VkInstance
	maximumVersion common.APIVersion

	activeInstanceExtensions map[string]struct{}
}

func (i Instance) Handle() driver.VkInstance {
	return i.instanceHandle
}

func (i Instance) APIVersion() common.APIVersion {
	return i.maximumVersion
}

func (i Instance) IsInstanceExtensionActive(extensionName string) bool {
	_, active := i.activeInstanceExtensions[extensionName]
	return active
}

func InternalInstance(handle driver.VkInstance, version common.APIVersion, extensions []string) Instance {
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
	physicalDeviceHandle driver.VkPhysicalDevice

	instanceVersion      common.APIVersion
	maximumDeviceVersion common.APIVersion
}

func (d PhysicalDevice) Handle() driver.VkPhysicalDevice {
	return d.physicalDeviceHandle
}

func (d PhysicalDevice) DeviceAPIVersion() common.APIVersion {
	return d.maximumDeviceVersion
}

func (d PhysicalDevice) InstanceAPIVersion() common.APIVersion {
	return d.instanceVersion
}

func InternalPhysicalDevice(handle driver.VkPhysicalDevice, instanceVersion common.APIVersion, deviceVersion common.APIVersion) PhysicalDevice {
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
	device         driver.VkDevice
	pipelineHandle driver.VkPipeline

	apiVersion common.APIVersion
}

func (p Pipeline) Handle() driver.VkPipeline {
	return p.pipelineHandle
}

func (p Pipeline) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p Pipeline) APIVersion() common.APIVersion {
	return p.apiVersion
}

func InternalPipeline(device driver.VkDevice, handle driver.VkPipeline, version common.APIVersion) Pipeline {
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
	device              driver.VkDevice
	pipelineCacheHandle driver.VkPipelineCache

	apiVersion common.APIVersion
}

func (c PipelineCache) Handle() driver.VkPipelineCache {
	return c.pipelineCacheHandle
}

func (c PipelineCache) DeviceHandle() driver.VkDevice {
	return c.device
}

func (c PipelineCache) APIVersion() common.APIVersion {
	return c.apiVersion
}

func InternalPipelineCache(device driver.VkDevice, handle driver.VkPipelineCache, version common.APIVersion) PipelineCache {
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
	device               driver.VkDevice
	pipelineLayoutHandle driver.VkPipelineLayout

	apiVersion common.APIVersion
}

func (l PipelineLayout) Handle() driver.VkPipelineLayout {
	return l.pipelineLayoutHandle
}

func (l PipelineLayout) DeviceHandle() driver.VkDevice {
	return l.device
}

func (l PipelineLayout) APIVersion() common.APIVersion {
	return l.apiVersion
}

func InternalPipelineLayout(device driver.VkDevice, handle driver.VkPipelineLayout, version common.APIVersion) PipelineLayout {
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
	queryPoolHandle driver.VkQueryPool
	device          driver.VkDevice

	apiVersion common.APIVersion
}

func (p QueryPool) Handle() driver.VkQueryPool {
	return p.queryPoolHandle
}

func (p QueryPool) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p QueryPool) APIVersion() common.APIVersion {
	return p.apiVersion
}

func InternalQueryPool(device driver.VkDevice, handle driver.VkQueryPool, version common.APIVersion) QueryPool {
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
	device      driver.VkDevice
	queueHandle driver.VkQueue

	apiVersion common.APIVersion
}

func (q Queue) Handle() driver.VkQueue {
	return q.queueHandle
}

func (q Queue) DeviceHandle() driver.VkDevice {
	return q.device
}

func (q Queue) APIVersion() common.APIVersion {
	return q.apiVersion
}

func InternalQueue(device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) Queue {
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
	device           driver.VkDevice
	renderPassHandle driver.VkRenderPass

	apiVersion common.APIVersion
}

func (p RenderPass) Handle() driver.VkRenderPass {
	return p.renderPassHandle
}

func (p RenderPass) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p RenderPass) APIVersion() common.APIVersion {
	return p.apiVersion
}

func InternalRenderPass(device driver.VkDevice, handle driver.VkRenderPass, version common.APIVersion) RenderPass {
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
	device        driver.VkDevice
	samplerHandle driver.VkSampler

	apiVersion common.APIVersion
}

func (s Sampler) Handle() driver.VkSampler {
	return s.samplerHandle
}

func (s Sampler) DeviceHandle() driver.VkDevice {
	return s.device
}

func (s Sampler) APIVersion() common.APIVersion {
	return s.apiVersion
}

func InternalSampler(device driver.VkDevice, handle driver.VkSampler, version common.APIVersion) Sampler {
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
	device          driver.VkDevice
	semaphoreHandle driver.VkSemaphore

	apiVersion common.APIVersion
}

func (s Semaphore) Handle() driver.VkSemaphore {
	return s.semaphoreHandle
}

func (s Semaphore) DeviceHandle() driver.VkDevice {
	return s.device
}

func (s Semaphore) APIVersion() common.APIVersion {
	return s.apiVersion
}

func InternalSemaphore(device driver.VkDevice, handle driver.VkSemaphore, version common.APIVersion) Semaphore {
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
	device             driver.VkDevice
	shaderModuleHandle driver.VkShaderModule

	apiVersion common.APIVersion
}

func (m ShaderModule) Handle() driver.VkShaderModule {
	return m.shaderModuleHandle
}

func (m ShaderModule) DeviceHandle() driver.VkDevice {
	return m.device
}

func (m ShaderModule) APIVersion() common.APIVersion {
	return m.apiVersion
}

func InternalShaderModule(device driver.VkDevice, handle driver.VkShaderModule, version common.APIVersion) ShaderModule {
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
	device      driver.VkDevice
	ycbcrHandle driver.VkSamplerYcbcrConversion

	apiVersion common.APIVersion
}

func (y SamplerYcbcrConversion) Handle() driver.VkSamplerYcbcrConversion {
	return y.ycbcrHandle
}

func (y SamplerYcbcrConversion) DeviceHandle() driver.VkDevice {
	return y.device
}

func (y SamplerYcbcrConversion) APIVersion() common.APIVersion {
	return y.apiVersion
}

func InternalSamplerYcbcrConversion(device driver.VkDevice, handle driver.VkSamplerYcbcrConversion, version common.APIVersion) SamplerYcbcrConversion {
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
	device                   driver.VkDevice
	descriptorTemplateHandle driver.VkDescriptorUpdateTemplate

	apiVersion common.APIVersion
}

func (t DescriptorUpdateTemplate) Handle() driver.VkDescriptorUpdateTemplate {
	return t.descriptorTemplateHandle
}

func (t DescriptorUpdateTemplate) DeviceHandle() driver.VkDevice {
	return t.device
}

func (t DescriptorUpdateTemplate) APIVersion() common.APIVersion {
	return t.apiVersion
}

func InternalDescriptorUpdateTemplate(device driver.VkDevice, handle driver.VkDescriptorUpdateTemplate, version common.APIVersion) DescriptorUpdateTemplate {
	return DescriptorUpdateTemplate{
		device:                   device,
		descriptorTemplateHandle: handle,
		apiVersion:               version,
	}
}
