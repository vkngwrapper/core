package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/core1_1_mocks.go -package mocks -mock_names CommandBuffer=CommandBuffer1_1,CommandPool=CommandPool1_1,Device=Device1_1,DescriptorUpdateTemplate=DescriptorUpdateTemplate1_1,Instance=Instance1_1,PhysicalDevice=PhysicalDevice1_1,Buffer=Buffer1_1,BufferView=BufferView1_1,DescriptorPool=DescriptorPool1_1,DescriptorSet=DescriptorSet1_1,DescriptorSetLayout=DescriptorSetLayout1_1,DeviceMemory=DeviceMemory1_1,Event=Event1_1,Fence=Fence1_1,Framebuffer=Framebuffer1_1,Image=Image1_1,ImageView=ImageView1_1,Pipeline=Pipeline1_1,PipelineCache=PipelineCache1_1,PipelineLayout=PipelineLayout1_1,QueryPool=QueryPool1_1,Queue=Queue1_1,RenderPass=RenderPass1_1,Sampler=Sampler1_1,Semaphore=Semaphore1_1,ShaderModule=ShaderModule1_1

// Buffer represents a linear array of data, which is used for various purposes by binding it
// to a graphics or compute pipeline.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBuffer.html
type Buffer interface {
	core1_0.Buffer
}

// BufferView represents a contiguous range of a buffer and a specific format to be used to
// interpret the data.
//
// This interface includes all commands included in Vulkan 1.1.
//
//https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferView.html
type BufferView interface {
	core1_0.BufferView
}

// CommandBuffer is an object used to record commands which can be subsequently submitted to
// a device queue for execution.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html
type CommandBuffer interface {
	core1_0.CommandBuffer

	CmdDispatchBase(baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int)
	CmdSetDeviceMask(deviceMask uint32)
}

// CommandPool is an opaque object that CommandBuffer memory is allocated from
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html
type CommandPool interface {
	core1_0.CommandPool

	TrimCommandPool(flags CommandPoolTrimFlags)
}

// DescriptorPool maintains a pool of descriptors, from which DescriptorSet objects are allocated.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html
type DescriptorPool interface {
	core1_0.DescriptorPool
}

// DescriptorSet is an opaque object allocated from a DescriptorPool
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetDescriptorPool.html
type DescriptorSet interface {
	core1_0.DescriptorSet
}

// DescriptorSetLayout is a group of zero or more descriptor bindings definitions.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html
type DescriptorSetLayout interface {
	core1_0.DescriptorSetLayout
}

// Device represents a logical device on the host
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDevice.html
type Device interface {
	core1_0.Device

	BindBufferMemory2(o []BindBufferMemoryInfo) (common.VkResult, error)
	BindImageMemory2(o []BindImageMemoryInfo) (common.VkResult, error)

	BufferMemoryRequirements2(o BufferMemoryRequirementsInfo2, out *MemoryRequirements2) error
	ImageMemoryRequirements2(o ImageMemoryRequirementsInfo2, out *MemoryRequirements2) error
	ImageSparseMemoryRequirements2(o ImageSparseMemoryRequirementsInfo2, outDataFactory func() *SparseImageMemoryRequirements2) ([]*SparseImageMemoryRequirements2, error)

	DescriptorSetLayoutSupport(o core1_0.DescriptorSetLayoutCreateInfo, outData *DescriptorSetLayoutSupport) error

	DeviceGroupPeerMemoryFeatures(heapIndex, localDeviceIndex, remoteDeviceIndex int) PeerMemoryFeatureFlags

	CreateDescriptorUpdateTemplate(o DescriptorUpdateTemplateCreateInfo, allocator *driver.AllocationCallbacks) (DescriptorUpdateTemplate, common.VkResult, error)
	CreateSamplerYcbcrConversion(o SamplerYcbcrConversionCreateInfo, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error)

	GetQueue2(o DeviceQueueInfo2) (core1_0.Queue, error)
}

// DeviceMemory represents a block of memory on the device
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html
type DeviceMemory interface {
	core1_0.DeviceMemory
}

// DescriptorUpdateTemplate specifies a mapping from descriptor update information in host memory to
// descriptors in a DescriptorSet
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html
type DescriptorUpdateTemplate interface {
	// Handle is the internal Vulkan object handle for this DescriptorUpdateTemplate
	Handle() driver.VkDescriptorUpdateTemplate
	// DeviceHandle is the internal Vulkan object handle for the Device this DescriptorUpdateTemplate belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this DescriptorUpdateTemplate
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this DescriptorUpdateTemplate. If it is at least
	// Vulkan 1.2, core1_2.PromoteDescriptorUpdateTemplate can be used to promote this to a core1_2.DescriptorUpdateTemplate,
	// etc.
	APIVersion() common.APIVersion

	// Destroy destroys the DescriptorUpdateTemplate object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorUpdateTemplate.html
	Destroy(allocator *driver.AllocationCallbacks)
	// UpdateDescriptorSetFromBuffer updates the contents of a DescriptorSet object with this template
	// and a Buffer
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - Information and a Buffer used to write the descriptor
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetFromBuffer(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorBufferInfo)
	// UpdateDescriptorSetFromImage updates the contents of a DescriptorSet object with this template and
	// an Image
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - Information and an Image used to write the descriptor
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetFromImage(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorImageInfo)
	// UpdateDescriptorSetFromObjectHandle updates the contents of a DescriptorSet object with this template
	// and an arbitrary handle
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - A Vulkan object handle used to write the descriptor. Can be a BufferView handle or
	// perhaps an acceleration structure.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetFromObjectHandle(descriptorSet core1_0.DescriptorSet, data driver.VulkanHandle)
}

// Event is a synchronization primitive that can be used to insert fine-grained dependencies between
// commands submitted to the same queue, or between the host and a queue.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkEvent.html
type Event interface {
	core1_0.Event
}

// Fence is a synchronization primitive that can be used to insert a dependency from a queue to
// the host.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkFence.html
type Fence interface {
	core1_0.Fence
}

// Framebuffer represents a collection of specific memory attachments that a RenderPass uses
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html
type Framebuffer interface {
	core1_0.Framebuffer
}

// Image represents multidimensional arrays of data which can be used for various purposes.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImage.html
type Image interface {
	core1_0.Image
}

// ImageView represents contiguous ranges of Image subresources and contains additional metadata
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImageView.html
type ImageView interface {
	core1_0.ImageView
}

// Instance stores per-application state for Vulkan
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkInstance.html
type Instance interface {
	core1_0.Instance

	// EnumeratePhysicalDeviceGroups enumerates groups of PHysicalDevice objects that can be used to
	// create a single logical Device
	//
	// outDataFactory - This method can be provided to allocate each PhysicalDeviceGroupProperties object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// PhysicalDeviceGroupProperties will be allocated with no chained structures.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceGroups.html
	EnumeratePhysicalDeviceGroups(outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error)
}

// InstanceScopedPhysicalDevice represents the instance-scoped functionality of a single complete
// implementation of Vulkan available to the host, of which there are a finite number.
//
// This interface includes all instance-scoped commands included in Vulkan 1.1.
//
// PhysicalDevice objects are unusual in that they exist between the Instance and (logical) Device level.
// As a result, PhysicalDevice is the only object that can be extended by both Instance and Device
// extensions. Consequently, there are some unusual cases in which a higher core version may be available
// for some PhysicalDevice functionality but not others. In order to represent this, physical devices
// are split into two objects at core1.1+, the PhysicalDevice and the "instance-scoped" PhysicalDevice.
//
// The InstanceScopedPhysicalDevice is usually available at the same core versions as PhysicalDevice, but
// in rare cases, a higher core version may be available.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html
type InstanceScopedPhysicalDevice interface {
	core1_0.PhysicalDevice

	// ExternalFenceProperties queries external Fence capabilities
	//
	// o - Describes the parameters that would be consumed by Device.CreateFence
	//
	// outData - A pre-allocated object in which the results will be populated. It should include
	// any desired chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalFenceProperties.html
	ExternalFenceProperties(o PhysicalDeviceExternalFenceInfo, outData *ExternalFenceProperties) error
	// ExternalBufferProperties queries external types supported by Buffer objects
	//
	// o - Describes the parameters that would be consumed by Device.CreateBuffer
	//
	// outData - A pre-allocated object in which the results will be populated. It should include any
	// desired chained OutData objects.
	//
	// https://www.khronos.org/registry/VulkanSC/specs/1.0-extensions/man/html/vkGetPhysicalDeviceExternalBufferProperties.html
	ExternalBufferProperties(o PhysicalDeviceExternalBufferInfo, outData *ExternalBufferProperties) error
	// ExternalSemaphoreProperties queries external Semaphore capabilities
	//
	// o - Describes the parameters that would be consumed by Device.CreateSemaphore
	//
	// outData - A pre-allocated object in which the results will be populated. It should include any
	// desired chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html
	ExternalSemaphoreProperties(o PhysicalDeviceExternalSemaphoreInfo, outData *ExternalSemaphoreProperties) error

	// Features2 reports capabilities of a PhysicalDevice
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html
	Features2(out *PhysicalDeviceFeatures2) error
	// FormatProperties2 lists the PhysicalDevice object's format capabilities
	//
	// format - The format whose properties are queried
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html
	FormatProperties2(format core1_0.Format, out *FormatProperties2) error
	// ImageFormatProperties2 lists the PhysicalDevice object's format capabilities
	//
	// o - Describes the parameters that would be consumed by Device.CreateImage
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html
	ImageFormatProperties2(o PhysicalDeviceImageFormatInfo2, out *ImageFormatProperties2) (common.VkResult, error)
	// MemoryProperties2 reports memory information for this PhysicalDevice
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2.html
	MemoryProperties2(out *PhysicalDeviceMemoryProperties2) error
	// Properties2 returns properties of this PhysicalDevice
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html
	Properties2(out *PhysicalDeviceProperties2) error
	// QueueFamilyProperties2 reports properties of the queues of this PhysicalDevice
	//
	// outDataFactory - This method can be provided to allocate each QueueFamilyProperties2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// QueueFamilyProperties2 will be allocated with no chained structures.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html
	QueueFamilyProperties2(outDataFactory func() *QueueFamilyProperties2) ([]*QueueFamilyProperties2, error)
	// SparseImageFormatProperties2 retrieves properties of an Image format applied to sparse Image
	//
	// o - Contains input parameters
	//
	// outDataFactory - This method can be provided to allocate each SparseImageFormatProperties2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// SparseImageFormatProperties2 will be allocated with no chained structures.
	SparseImageFormatProperties2(o PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *SparseImageFormatProperties2) ([]*SparseImageFormatProperties2, error)
}

// PhysicalDevice represents the device-scoped functionality of a single complete
// implementation of Vulkan available to the host, of which there are a finite number.
//
// This interface includes all commands included in Vulkan 1.1.
//
// PhysicalDevice objects are unusual in that they exist between the Instance and (logical) Device level.
// As a result, PhysicalDevice is the only object that can be extended by both Instance and Device
// extensions. Consequently, there are some unusual cases in which a higher core version may be available
// for some PhysicalDevice functionality but not others. In order to represent this, physical devices
// are split into two objects at core1.1+, the PhysicalDevice and the "instance-scoped" PhysicalDevice.
//
// The InstanceScopedPhysicalDevice is usually available at the same core versions as PhysicalDevice, but
// in rare cases, a higher core version may be available.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html
type PhysicalDevice interface {
	core1_0.PhysicalDevice

	// InstanceScopedPhysicalDevice1_1 returns the InstanceScopedPhysicalDevice that represents the
	// instance-scoped portion of this PhysicalDevice object's functionality. Since the instance-scoped
	// support is always equal-to-or-greater-than the device-scoped support, this method will always
	// return a functioning InstanceScopedPhysicalDevice
	InstanceScopedPhysicalDevice1_1() InstanceScopedPhysicalDevice
}

// Pipeline represents compute, ray tracing, and graphics pipelines
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipeline.html
type Pipeline interface {
	core1_0.Pipeline
}

// PipelineCache allows the result of Pipeline construction to be reused between Pipeline objects
// and between runs of an application.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html
type PipelineCache interface {
	core1_0.PipelineCache
}

// PipelineLayout provides access to descriptor sets to Pipeline objects by combining zero or more
// descriptor sets and zero or more push constant ranges.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html
type PipelineLayout interface {
	core1_0.PipelineLayout
}

// QueryPool is a collection of a specific number of queries of a particular type.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html
type QueryPool interface {
	core1_0.QueryPool
}

// Queue represents a Device resource on which work is performed
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkQueue.html
type Queue interface {
	core1_0.Queue
}

// RenderPass represents a collection of attachments, subpasses, and dependencies between the subpasses
// and describes how the attachments are used over the course of the subpasses
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html
type RenderPass interface {
	core1_0.RenderPass
}

// Sampler represents the state of an Image sampler, which is used by the implementation to read Image data
// and apply filtering and other transformations for the shader.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSampler.html
type Sampler interface {
	core1_0.Sampler
}

// SamplerYcbcrConversion is an opaque representation of a device-specific sampler YCbCr conversion
// description.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html
type SamplerYcbcrConversion interface {
	// Handle is the internal Vulkan object handle for this SamplerYcbcrConversion
	Handle() driver.VkSamplerYcbcrConversion
	// DeviceHandle is the internal Vulkan object handle for the Device this SamplerYcbcrConversion
	// belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this SamplerYcbcrConversion
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this SamplerYcbcrConversion. If it is at
	// least Vulkan 1.2, core1_2.PromoteSamplerYcbcrConversion can be used to promote this to a
	// core1_2.SamplerYcbcrConversion, etc.
	APIVersion() common.APIVersion

	Destroy(allocator *driver.AllocationCallbacks)
}

// Semaphore is a synchronization primitive that can be used to insert a dependency between Queue operations
// or between a Queue operation and the host.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html
type Semaphore interface {
	core1_0.Semaphore
}

// ShaderModule objects contain shader code and one or more entry points.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html
type ShaderModule interface {
	core1_0.ShaderModule
}
