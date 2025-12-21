package core1_1

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/mocks1_1/mocks.go -package mocks1_1

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
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferView.html
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

	// CmdDispatchBase dispatches compute work items with non-zero base values for the workgroup IDs
	//
	// baseGroupX - The start value of the X component of WorkgroupId
	//
	// baseGroupY - The start value of the Y component of WorkGroupId
	//
	// baseGroupZ - The start value of the Z component of WorkGroupId
	//
	// groupCountX - The number of local workgroups to dispatch in the X dimension
	//
	// groupCountY - The number of local workgroups to dispatch in the Y dimension
	//
	// groupCountZ - The number of local workgroups to dispatch in the Z dimension
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchBase.html
	CmdDispatchBase(baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int)
	// CmdSetDeviceMask modifies the device mask of a CommandBuffer
	//
	// deviceMask - The new value of the current Device mask
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDeviceMask.html
	CmdSetDeviceMask(deviceMask uint32)
}

// CommandPool is an opaque object that CommandBuffer memory is allocated from
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html
type CommandPool interface {
	core1_0.CommandPool

	// TrimCommandPool trims a CommandPool
	//
	// flags - Reserved for future use
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTrimCommandPool.html
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

	// BindBufferMemory2 binds DeviceMemory to Buffer objects
	//
	// o - A slice of BindBufferMemoryInfo structures describing Buffer objects and memory to bind
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory2.html
	BindBufferMemory2(o []BindBufferMemoryInfo) (common.VkResult, error)
	// BindImageMemory2 binds DeviceMemory to Image objects
	//
	// o - A slice of BindImageMemoryInfo structures describing Image objects and memory to bind
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory2.html
	BindImageMemory2(o []BindImageMemoryInfo) (common.VkResult, error)

	// BufferMemoryRequirements2 returns the memory requirements for the specified Vulkan object
	//
	// o - Contains parameters required for the memory requirements query
	//
	// out - A pre-allocated object in which the memory requirements of the Buffer object will be
	// populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2.html
	BufferMemoryRequirements2(o BufferMemoryRequirementsInfo2, out *MemoryRequirements2) error
	// ImageMemoryRequirements2 returns the memory requirements for the specified Vulkan object
	//
	// o - Contains parameters required for the memory requirements query
	//
	// out - A pre-allocated object in which the memory requirements of the Image object will be
	// populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html
	ImageMemoryRequirements2(o ImageMemoryRequirementsInfo2, out *MemoryRequirements2) error
	// ImageSparseMemoryRequirements2 queries the memory requirements for a sparse Image
	//
	// o - Contains parameters required for the memory requirements query
	//
	// outDataFactory - This method can be provided to allocate each SparseImageMemoryRequirements2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// SparseImageMemoryRequirements2 will be allocated with no chained structures.
	ImageSparseMemoryRequirements2(o ImageSparseMemoryRequirementsInfo2, outDataFactory func() *SparseImageMemoryRequirements2) ([]*SparseImageMemoryRequirements2, error)

	// DescriptorSetLayoutSupport queries whether a DescriptorSetLayout can be created
	//
	// o - Specifies the state of the DescriptorSetLayout object
	//
	// outData - A pre-allocated object in which information about support for the DescriptorSetLayout
	// object will be populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupport.html
	DescriptorSetLayoutSupport(o core1_0.DescriptorSetLayoutCreateInfo, outData *DescriptorSetLayoutSupport) error

	// DeviceGroupPeerMemoryFeatures queries supported peer memory features of a Device
	//
	// heapIndex - The index of the memory heap from which the memory is allocated
	//
	// localDeviceIndex - The device index of the PhysicalDevice that performs the memory access
	//
	// remoteDeviceIndex - The device index of the PhysicalDevice that the memory is allocated for
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPeerMemoryFeatures.html
	DeviceGroupPeerMemoryFeatures(heapIndex, localDeviceIndex, remoteDeviceIndex int) PeerMemoryFeatureFlags

	// CreateDescriptorUpdateTemplate creates a new DescriptorUpdateTemplate
	//
	// o - Specifies the set of descriptors to update with a single call to DescriptorUpdateTemplate.UpdateDescriptorSet...
	//
	// allocator - Controls host allocation behavior
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html
	CreateDescriptorUpdateTemplate(o DescriptorUpdateTemplateCreateInfo, allocator *driver.AllocationCallbacks) (DescriptorUpdateTemplate, common.VkResult, error)
	// CreateSamplerYcbcrConversion creates a new Y'CbCr conversion
	//
	// o - Specifies the requested sampler Y'CbCr conversion
	//
	// allocator - Controls host allocation behavior
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSamplerYcbcrConversion.html
	CreateSamplerYcbcrConversion(o SamplerYcbcrConversionCreateInfo, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error)

	// GetQueue2 gets a Queue object from this Device
	//
	// o - Describes parameters of the Device Queue to be retrieved
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceQueue2.html
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

	// EnumeratePhysicalDeviceGroups enumerates groups of PhysicalDevice objects that can be used to
	// create a single logical Device
	//
	// outDataFactory - This method can be provided to allocate each PhysicalDeviceGroupProperties object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// PhysicalDeviceGroupProperties will be allocated with no chained structures.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceGroups.html
	EnumeratePhysicalDeviceGroups(outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error)
}

// PhysicalDevice represents the device-scoped functionality of a single complete
// implementation of Vulkan available to the host, of which there are a finite number.
//
// This interface includes all commands included in Vulkan 1.1.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html
type PhysicalDevice interface {
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
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2KHR.html
	SparseImageFormatProperties2(o PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *SparseImageFormatProperties2) ([]*SparseImageFormatProperties2, error)
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

	// Destroy destroys the SamplerYcbcrConversion object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySamplerYcbcrConversion.html
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

// DeviceObjectBuilder is an internal type exposed by Device to allow objects to be create from
// vulkan handles.  This used by extensions and should not be used by most consumers.
type DeviceObjectBuilder interface {
	CreateDescriptorUpdateTemplate(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorUpdateTemplate, version common.APIVersion) DescriptorUpdateTemplate
	CreateSamplerYcbcrConversion(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSamplerYcbcrConversion, version common.APIVersion) SamplerYcbcrConversion
}
