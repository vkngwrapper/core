package core1_1

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/mocks1_1/mocks.go -package mocks1_1

type CoreInstanceDriver interface {
	core1_0.CoreInstanceDriver

	// GetPhysicalDeviceExternalFenceProperties queries external Fence capabilities
	//
	// physicalDevice - the PHysicalDevice object to query
	//
	// o - Describes the parameters that would be consumed by Device.CreateFence
	//
	// outData - A pre-allocated object in which the results will be populated. It should include
	// any desired chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalFenceProperties.html
	GetPhysicalDeviceExternalFenceProperties(physicalDevice types.PhysicalDevice, o PhysicalDeviceExternalFenceInfo, outData *ExternalFenceProperties) error
	// GetPhysicalDeviceExternalBufferProperties queries external types supported by Buffer objects
	//
	// physicalDevice - the PHysicalDevice object to query
	//
	// o - Describes the parameters that would be consumed by Device.CreateBuffer
	//
	// outData - A pre-allocated object in which the results will be populated. It should include any
	// desired chained OutData objects.
	//
	// https://www.khronos.org/registry/VulkanSC/specs/1.0-extensions/man/html/vkGetPhysicalDeviceExternalBufferProperties.html
	GetPhysicalDeviceExternalBufferProperties(physicalDevice types.PhysicalDevice, o PhysicalDeviceExternalBufferInfo, outData *ExternalBufferProperties) error
	// GetPhysicalDeviceExternalSemaphoreProperties queries external Semaphore capabilities
	//
	// physicalDevice - the PHysicalDevice object to query
	//
	// o - Describes the parameters that would be consumed by Device.CreateSemaphore
	//
	// outData - A pre-allocated object in which the results will be populated. It should include any
	// desired chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html
	GetPhysicalDeviceExternalSemaphoreProperties(physicalDevice types.PhysicalDevice, o PhysicalDeviceExternalSemaphoreInfo, outData *ExternalSemaphoreProperties) error

	// GetPhysicalDeviceFeatures2 reports capabilities of a PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice object to query
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html
	GetPhysicalDeviceFeatures2(physicalDevice types.PhysicalDevice, out *PhysicalDeviceFeatures2) error
	// GetPhysicalDeviceFormatProperties2 lists a PhysicalDevice object's format capabilities
	//
	// physicalDevice - The PhysicalDevice object to query
	//
	// format - The format whose properties are queried
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html
	GetPhysicalDeviceFormatProperties2(physicalDevice types.PhysicalDevice, format core1_0.Format, out *FormatProperties2) error
	// GetPhysicalDeviceImageFormatProperties2 lists a PhysicalDevice object's format capabilities
	//
	// physicalDevice - The PhysicalDevice object to query
	//
	// o - Describes the parameters that would be consumed by Device.CreateImage
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html
	GetPhysicalDeviceImageFormatProperties2(physicalDevice types.PhysicalDevice, o PhysicalDeviceImageFormatInfo2, out *ImageFormatProperties2) (common.VkResult, error)
	// GetPhysicalDeviceMemoryProperties2 reports memory information for a PhysicalDevice
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// physicalDevice - The PhysicalDevice object to query
	//
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2.html
	GetPhysicalDeviceMemoryProperties2(physicalDevice types.PhysicalDevice, out *PhysicalDeviceMemoryProperties2) error
	// GetPhysicalDeviceProperties2 returns properties of a PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice object to query
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html
	GetPhysicalDeviceProperties2(physicalDevice types.PhysicalDevice, out *PhysicalDeviceProperties2) error
	// GetPhysicalDeviceQueueFamilyProperties2 reports properties of the queues of a PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice object to query
	//
	// outDataFactory - This method can be provided to allocate each QueueFamilyProperties2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// QueueFamilyProperties2 will be allocated with no chained structures.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html
	GetPhysicalDeviceQueueFamilyProperties2(physicalDevice types.PhysicalDevice, outDataFactory func() *QueueFamilyProperties2) ([]*QueueFamilyProperties2, error)
	// GetPhysicalDeviceSparseImageFormatProperties2 retrieves properties of an Image format applied to sparse Image
	//
	// physicalDevice - The PhysicalDevice object to query
	//
	// o - Contains input parameters
	//
	// outDataFactory - This method can be provided to allocate each SparseImageFormatProperties2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// SparseImageFormatProperties2 will be allocated with no chained structures.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2KHR.html
	GetPhysicalDeviceSparseImageFormatProperties2(physicalDevice types.PhysicalDevice, o PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *SparseImageFormatProperties2) ([]*SparseImageFormatProperties2, error)

	// EnumeratePhysicalDeviceGroups enumerates groups of PhysicalDevice objects that can be used to
	// create a single logical Device
	//
	// instance - The Instance whose PhysicalDevice objects should be enumerated
	//
	// outDataFactory - This method can be provided to allocate each PhysicalDeviceGroupProperties object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// PhysicalDeviceGroupProperties will be allocated with no chained structures.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceGroups.html
	EnumeratePhysicalDeviceGroups(instance types.Instance, outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error)
}

type DeviceDriver interface {
	core1_0.DeviceDriver

	// CmdDispatchBase dispatches compute work items with non-zero base values for the workgroup IDs
	//
	// commandBuffer - The CommandBuffer object to record to
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
	CmdDispatchBase(commandBuffer types.CommandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int)
	// CmdSetDeviceMask modifies the device mask of a CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// deviceMask - The new value of the current Device mask
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDeviceMask.html
	CmdSetDeviceMask(commandBuffer types.CommandBuffer, deviceMask uint32)

	// TrimCommandPool trims a CommandPool
	//
	// flags - Reserved for future use
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTrimCommandPool.html
	TrimCommandPool(commandPool types.CommandPool, flags CommandPoolTrimFlags)

	// BindBufferMemory2 binds DeviceMemory to Buffer objects
	//
	// o - A slice of BindBufferMemoryInfo structures describing Buffer objects and memory to bind
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory2.html
	BindBufferMemory2(o ...BindBufferMemoryInfo) (common.VkResult, error)
	// BindImageMemory2 binds DeviceMemory to Image objects
	//
	// o - A slice of BindImageMemoryInfo structures describing Image objects and memory to bind
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory2.html
	BindImageMemory2(o ...BindImageMemoryInfo) (common.VkResult, error)

	// GetBufferMemoryRequirements2 returns the memory requirements for the specified Vulkan object
	//
	// o - Contains parameters required for the memory requirements query
	//
	// out - A pre-allocated object in which the memory requirements of the Buffer object will be
	// populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2.html
	GetBufferMemoryRequirements2(o BufferMemoryRequirementsInfo2, out *MemoryRequirements2) error
	// GetImageMemoryRequirements2 returns the memory requirements for the specified Vulkan object
	//
	// o - Contains parameters required for the memory requirements query
	//
	// out - A pre-allocated object in which the memory requirements of the Image object will be
	// populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html
	GetImageMemoryRequirements2(o ImageMemoryRequirementsInfo2, out *MemoryRequirements2) error
	// GetImageSparseMemoryRequirements2 queries the memory requirements for a sparse Image
	//
	// o - Contains parameters required for the memory requirements query
	//
	// outDataFactory - This method can be provided to allocate each SparseImageMemoryRequirements2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// SparseImageMemoryRequirements2 will be allocated with no chained structures.
	GetImageSparseMemoryRequirements2(o ImageSparseMemoryRequirementsInfo2, outDataFactory func() *SparseImageMemoryRequirements2) ([]*SparseImageMemoryRequirements2, error)

	// GetDescriptorSetLayoutSupport queries whether a DescriptorSetLayout can be created
	//
	// device - The Device that will be used to create the descriptor set
	//
	// o - Specifies the state of the DescriptorSetLayout object
	//
	// outData - A pre-allocated object in which information about support for the DescriptorSetLayout
	// object will be populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupport.html
	GetDescriptorSetLayoutSupport(device types.Device, o core1_0.DescriptorSetLayoutCreateInfo, outData *DescriptorSetLayoutSupport) error

	// GetDeviceGroupPeerMemoryFeatures queries supported peer memory features of a Device
	//
	// device - The Device object being queried
	//
	// heapIndex - The index of the memory heap from which the memory is allocated
	//
	// localDeviceIndex - The device index of the PhysicalDevice that performs the memory access
	//
	// remoteDeviceIndex - The device index of the PhysicalDevice that the memory is allocated for
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPeerMemoryFeatures.html
	GetDeviceGroupPeerMemoryFeatures(device types.Device, heapIndex, localDeviceIndex, remoteDeviceIndex int) PeerMemoryFeatureFlags

	// CreateDescriptorUpdateTemplate creates a new DescriptorUpdateTemplate
	//
	// device - The Device used to create the DescriptorUpdateTemplate
	//
	// o - Specifies the set of descriptors to update with a single call to DescriptorUpdateTemplate.UpdateDescriptorSet...
	//
	// allocator - Controls host allocation behavior
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html
	CreateDescriptorUpdateTemplate(device types.Device, o DescriptorUpdateTemplateCreateInfo, allocator *loader.AllocationCallbacks) (types.DescriptorUpdateTemplate, common.VkResult, error)
	// CreateSamplerYcbcrConversion creates a new Y'CbCr conversion
	//
	// device - The Device used to created the SamplerYcbcrConversion
	//
	// o - Specifies the requested sampler Y'CbCr conversion
	//
	// allocator - Controls host allocation behavior
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSamplerYcbcrConversion.html
	CreateSamplerYcbcrConversion(device types.Device, o SamplerYcbcrConversionCreateInfo, allocator *loader.AllocationCallbacks) (types.SamplerYcbcrConversion, common.VkResult, error)

	// GetDeviceQueue2 gets a Queue object from a Device
	//
	// device - The Device to retrieve the queue from
	//
	// o - Describes parameters of the Device Queue to be retrieved
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceQueue2.html
	GetDeviceQueue2(device types.Device, o DeviceQueueInfo2) (types.Queue, error)

	// DestroyDescriptorUpdateTemplate destroys a DescriptorUpdateTemplate object and the
	// underlying structures. **Warning** after destruction, the object will continue to exist,
	// but the Vulkan object handle that backs it will be invalid. Do not call further methods
	// with the DescriptorUpdateTemplate
	//
	// template - The DescriptorUpdateTemplate to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorUpdateTemplate.html
	DestroyDescriptorUpdateTemplate(template types.DescriptorUpdateTemplate, allocator *loader.AllocationCallbacks)
	// UpdateDescriptorSetWithTemplateFromBuffer updates the contents of a DescriptorSet object
	// with a template and a Buffer
	//
	// template - The DescriptorUpdateTemplate used to perform the update
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - Information and a Buffer used to write the descriptor
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetWithTemplateFromBuffer(descriptorSet types.DescriptorSet, template types.DescriptorUpdateTemplate, data core1_0.DescriptorBufferInfo)
	// UpdateDescriptorSetWithTemplateFromImage updates the contents of a DescriptorSet object with
	// a template and an Image
	//
	// template - The DescriptorUpdateTemplate used to perform the update
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - Information and an Image used to write the descriptor
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetWithTemplateFromImage(descriptorSet types.DescriptorSet, template types.DescriptorUpdateTemplate, data core1_0.DescriptorImageInfo)
	// UpdateDescriptorSetWithTemplateFromObjectHandle updates the contents of a DescriptorSet object with
	// a template and an arbitrary handle
	//
	// template - The DescriptorUpdateTemplate used to perform the update
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - A Vulkan object handle used to write the descriptor. Can be a BufferView handle or
	// perhaps an acceleration structure.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetWithTemplateFromObjectHandle(descriptorSet types.DescriptorSet, template types.DescriptorUpdateTemplate, data loader.VulkanHandle)

	// DestroySamplerYcbcrConversion destroys a SamplerYcbcrConversion object and the underlying
	// structures. **Warning** after destruction, tye object will continue to exist, but the Vulkan
	// object handle that backs it will be invalid. Do not call further methods with the
	// SamplerYcbcrConversion
	//
	// conversion - The SamplerYcbcrConversion to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySamplerYcbcrConversion.html
	DestroySamplerYcbcrConversion(conversion types.SamplerYcbcrConversion, allocator *loader.AllocationCallbacks)
}

type CoreDeviceDriver interface {
	CoreInstanceDriver
	DeviceDriver
}
