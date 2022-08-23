package core1_2

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// PhysicalDevice8BitStorageFeatures describes features supported by khr_8bit_storage
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice8BitStorageFeatures.html
type PhysicalDevice8BitStorageFeatures struct {
	// StorageBuffer8BitAccess indicates whether objects in the StorageBuffer, ShaderRecordBufferKHR,
	// or PhysicalStorageBuffer storage class with the Block decoration can have 8-bit integer members
	StorageBuffer8BitAccess bool
	// UniformAndStorageBuffer8BitAccess indicates whether objects in the Uniform storage class
	// with the Block decoration can have 8-bit integer members
	UniformAndStorageBuffer8BitAccess bool
	// StoragePushConstant8 indicates whether objects in the PushConstant storage class can have 8-bit
	// integer members
	StoragePushConstant8 bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDevice8BitStorageFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice8BitStorageFeatures{})))
	}

	outData := (*C.VkPhysicalDevice8BitStorageFeatures)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevice8BitStorageFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDevice8BitStorageFeatures)(cDataPointer)
	o.StoragePushConstant8 = outData.storagePushConstant8 != C.VkBool32(0)
	o.UniformAndStorageBuffer8BitAccess = outData.uniformAndStorageBuffer8BitAccess != C.VkBool32(0)
	o.StorageBuffer8BitAccess = outData.storageBuffer8BitAccess != C.VkBool32(0)

	return outData.pNext, nil
}

func (o PhysicalDevice8BitStorageFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice8BitStorageFeatures{})))
	}

	info := (*C.VkPhysicalDevice8BitStorageFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES
	info.pNext = next
	info.storageBuffer8BitAccess = C.VkBool32(0)
	info.uniformAndStorageBuffer8BitAccess = C.VkBool32(0)
	info.storagePushConstant8 = C.VkBool32(0)

	if o.StorageBuffer8BitAccess {
		info.storageBuffer8BitAccess = C.VkBool32(1)
	}

	if o.UniformAndStorageBuffer8BitAccess {
		info.uniformAndStorageBuffer8BitAccess = C.VkBool32(1)
	}

	if o.StoragePushConstant8 {
		info.storagePushConstant8 = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceBufferDeviceAddressFeatures describes Buffer address features that can
// be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBufferDeviceAddressFeatures.html
type PhysicalDeviceBufferDeviceAddressFeatures struct {
	// BufferDeviceAddress indicates that the implementation supports accessing Buffer memory
	// in shaders as storage Buffer objects via an address queried from Device.GetBufferDeviceAddress
	BufferDeviceAddress bool
	// BufferDeviceAddressCaptureReplay indicates that the implementation supports saving and
	// reusing Buffer and Device addresses, e.g. for trace capture and replay
	BufferDeviceAddressCaptureReplay bool
	// BufferDeviceAddressMultiDevice indicates that the implementation supports the
	// BufferDeviceAddress, RayTracingPipeline, and RayQuery features for logical Device objects
	// created with multiple PhysicalDevice objects
	BufferDeviceAddressMultiDevice bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceBufferDeviceAddressFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceBufferDeviceAddressFeatures{})))
	}

	info := (*C.VkPhysicalDeviceBufferDeviceAddressFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceBufferDeviceAddressFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceBufferDeviceAddressFeatures)(cDataPointer)

	o.BufferDeviceAddress = info.bufferDeviceAddress != C.VkBool32(0)
	o.BufferDeviceAddressCaptureReplay = info.bufferDeviceAddressCaptureReplay != C.VkBool32(0)
	o.BufferDeviceAddressMultiDevice = info.bufferDeviceAddressMultiDevice != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceBufferDeviceAddressFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceBufferDeviceAddressFeatures{})))
	}

	info := (*C.VkPhysicalDeviceBufferDeviceAddressFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES
	info.pNext = next
	info.bufferDeviceAddress = C.VkBool32(0)
	info.bufferDeviceAddressCaptureReplay = C.VkBool32(0)
	info.bufferDeviceAddressMultiDevice = C.VkBool32(0)

	if o.BufferDeviceAddress {
		info.bufferDeviceAddress = C.VkBool32(1)
	}

	if o.BufferDeviceAddressCaptureReplay {
		info.bufferDeviceAddressCaptureReplay = C.VkBool32(1)
	}

	if o.BufferDeviceAddressMultiDevice {
		info.bufferDeviceAddressMultiDevice = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceDescriptorIndexingFeatures describes descriptor indexing
// features that can be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html
type PhysicalDeviceDescriptorIndexingFeatures struct {
	// ShaderInputAttachmentArrayDynamicIndexing indicates whether arrays of input attachments
	// can be indexed by dynamically uniform integer expressions in shader code
	ShaderInputAttachmentArrayDynamicIndexing bool
	// ShaderUniformTexelBufferArrayDynamicIndexing indicates whether arrays of uniform texel
	// Buffer objects can be indexed by dynamically uniform integer expressions in shader code
	ShaderUniformTexelBufferArrayDynamicIndexing bool
	// ShaderStorageTexelBufferArrayDynamicIndexing indicates whether arrays of storage texel
	// Buffer objects can be indexed by dynamically uniform integer expressions in shader code
	ShaderStorageTexelBufferArrayDynamicIndexing bool
	// ShaderUniformBufferArrayNonUniformIndexing indicates whether arrays of uniform Buffer objects
	// can be indexed by non-uniform integer expressions in shader code.
	ShaderUniformBufferArrayNonUniformIndexing bool
	// ShaderSampledImageArrayNonUniformIndexing indicates whether arrays of Sampler objects or sampled
	// Image objects can be indexed by non-uniform integer expressions in shader code
	ShaderSampledImageArrayNonUniformIndexing bool
	// ShaderStorageBufferArrayNonUniformIndexing indicates whether arrays of storage buffers
	// can be indexed by non-uniform integer expressions in shader code
	ShaderStorageBufferArrayNonUniformIndexing bool
	// ShaderStorageImageArrayNonUniformIndexing indicates whether arrays of storage Image objects can
	// be indexed by non-uniform integer expressions in shader code
	ShaderStorageImageArrayNonUniformIndexing bool
	// ShaderInputAttachmentArrayNonUniformIndexing indicates whether arrays of input attachments
	// can be indexed by non-uniform integer expressions in shader code
	ShaderInputAttachmentArrayNonUniformIndexing bool
	// ShaderUniformTexelBufferArrayNonUniformIndexing indicates whether arrays of uniform texel
	// Buffer objects can be indexed by non-uniform integer expressions in shader code
	ShaderUniformTexelBufferArrayNonUniformIndexing bool
	// ShaderStorageTexelBufferArrayNonUniformIndexing indicates whether arrays of storage texel
	// Buffer objects can be indexed by non-uniform integer expressions in shader code
	ShaderStorageTexelBufferArrayNonUniformIndexing bool
	// DescriptorBindingUniformBufferUpdateAfterBind indicates whether the implementation supports
	// updating uniform Buffer descriptors after a set is bound
	DescriptorBindingUniformBufferUpdateAfterBind bool
	// DescriptorBindingSampledImageUpdateAfterBind indicates whether the implementation supports
	// updating sampled Image descriptors after a set is bound
	DescriptorBindingSampledImageUpdateAfterBind bool
	// DescriptorBindingStorageImageUpdateAfterBind indicates whether the implementation supports
	// updating storage Image descriptors after a set is bound
	DescriptorBindingStorageImageUpdateAfterBind bool
	// DescriptorBindingStorageBufferUpdateAfterBind indicates whether the implementation
	// supports updating storage Buffer descriptors after a set is bound
	DescriptorBindingStorageBufferUpdateAfterBind bool
	// DescriptorBindingUniformTexelBufferUpdateAfterBind indicates whether the implementation
	// supports updating uniform texel Buffer descriptors after a set is bound
	DescriptorBindingUniformTexelBufferUpdateAfterBind bool
	// DescriptorBindingStorageTexelBufferUpdateAfterBind indicates whether the impelementation
	// supports updating storage texel Buffer descriptors after a set is bound
	DescriptorBindingStorageTexelBufferUpdateAfterBind bool
	// DescriptorBindingUpdateUnusedWhilePending indicates whether the implementation supports
	// updating descriptors while the set is in use
	DescriptorBindingUpdateUnusedWhilePending bool
	// DescriptorBindingPartiallyBound indicates whether the implementation supports statically
	// using a DescriptorSet binding in which some descriptors are not valid
	DescriptorBindingPartiallyBound bool
	// DescriptorBindingVariableDescriptorCount indicates whether the implementation supports
	// DescriptorSet object with a variable-sized last binding
	DescriptorBindingVariableDescriptorCount bool
	// RuntimeDescriptorArray indicates whether the implementation supports the SPIR-V
	// RuntimeDescriptorArray capability
	RuntimeDescriptorArray bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceDescriptorIndexingFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDescriptorIndexingFeatures{})))
	}

	info := (*C.VkPhysicalDeviceDescriptorIndexingFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDescriptorIndexingFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceDescriptorIndexingFeatures)(cDataPointer)

	o.ShaderInputAttachmentArrayDynamicIndexing = info.shaderInputAttachmentArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderUniformTexelBufferArrayDynamicIndexing = info.shaderUniformTexelBufferArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderStorageTexelBufferArrayDynamicIndexing = info.shaderStorageTexelBufferArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderUniformBufferArrayNonUniformIndexing = info.shaderUniformBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderSampledImageArrayNonUniformIndexing = info.shaderSampledImageArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageBufferArrayNonUniformIndexing = info.shaderStorageBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageImageArrayNonUniformIndexing = info.shaderStorageImageArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderInputAttachmentArrayNonUniformIndexing = info.shaderInputAttachmentArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderUniformTexelBufferArrayNonUniformIndexing = info.shaderUniformTexelBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageTexelBufferArrayNonUniformIndexing = info.shaderStorageTexelBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.DescriptorBindingUniformBufferUpdateAfterBind = info.descriptorBindingUniformBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingSampledImageUpdateAfterBind = info.descriptorBindingSampledImageUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageImageUpdateAfterBind = info.descriptorBindingStorageImageUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageBufferUpdateAfterBind = info.descriptorBindingStorageBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingUniformTexelBufferUpdateAfterBind = info.descriptorBindingUniformTexelBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageTexelBufferUpdateAfterBind = info.descriptorBindingStorageTexelBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingUpdateUnusedWhilePending = info.descriptorBindingUpdateUnusedWhilePending != C.VkBool32(0)
	o.DescriptorBindingPartiallyBound = info.descriptorBindingPartiallyBound != C.VkBool32(0)
	o.DescriptorBindingVariableDescriptorCount = info.descriptorBindingVariableDescriptorCount != C.VkBool32(0)
	o.RuntimeDescriptorArray = info.runtimeDescriptorArray != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceDescriptorIndexingFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDescriptorIndexingFeatures{})))
	}

	info := (*C.VkPhysicalDeviceDescriptorIndexingFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES
	info.pNext = next
	info.shaderInputAttachmentArrayDynamicIndexing = C.VkBool32(0)
	info.shaderUniformTexelBufferArrayDynamicIndexing = C.VkBool32(0)
	info.shaderStorageTexelBufferArrayDynamicIndexing = C.VkBool32(0)
	info.shaderUniformBufferArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderSampledImageArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderStorageBufferArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderStorageImageArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderInputAttachmentArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderUniformTexelBufferArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderStorageTexelBufferArrayNonUniformIndexing = C.VkBool32(0)
	info.descriptorBindingUniformBufferUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingSampledImageUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingStorageImageUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingStorageBufferUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingUniformTexelBufferUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingStorageTexelBufferUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingUpdateUnusedWhilePending = C.VkBool32(0)
	info.descriptorBindingPartiallyBound = C.VkBool32(0)
	info.descriptorBindingVariableDescriptorCount = C.VkBool32(0)
	info.runtimeDescriptorArray = C.VkBool32(0)

	if o.ShaderInputAttachmentArrayDynamicIndexing {
		info.shaderInputAttachmentArrayDynamicIndexing = C.VkBool32(1)
	}

	if o.ShaderUniformTexelBufferArrayDynamicIndexing {
		info.shaderUniformTexelBufferArrayDynamicIndexing = C.VkBool32(1)
	}

	if o.ShaderStorageTexelBufferArrayDynamicIndexing {
		info.shaderStorageTexelBufferArrayDynamicIndexing = C.VkBool32(1)
	}

	if o.ShaderUniformBufferArrayNonUniformIndexing {
		info.shaderUniformBufferArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderSampledImageArrayNonUniformIndexing {
		info.shaderSampledImageArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderStorageBufferArrayNonUniformIndexing {
		info.shaderStorageBufferArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderStorageImageArrayNonUniformIndexing {
		info.shaderStorageImageArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderInputAttachmentArrayNonUniformIndexing {
		info.shaderInputAttachmentArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderUniformTexelBufferArrayNonUniformIndexing {
		info.shaderUniformTexelBufferArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderStorageTexelBufferArrayNonUniformIndexing {
		info.shaderStorageTexelBufferArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.DescriptorBindingUniformBufferUpdateAfterBind {
		info.descriptorBindingUniformBufferUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingSampledImageUpdateAfterBind {
		info.descriptorBindingSampledImageUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingStorageImageUpdateAfterBind {
		info.descriptorBindingStorageImageUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingStorageBufferUpdateAfterBind {
		info.descriptorBindingStorageBufferUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingUniformTexelBufferUpdateAfterBind {
		info.descriptorBindingUniformTexelBufferUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingStorageTexelBufferUpdateAfterBind {
		info.descriptorBindingStorageTexelBufferUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingUpdateUnusedWhilePending {
		info.descriptorBindingUpdateUnusedWhilePending = C.VkBool32(1)
	}

	if o.DescriptorBindingPartiallyBound {
		info.descriptorBindingPartiallyBound = C.VkBool32(1)
	}

	if o.DescriptorBindingVariableDescriptorCount {
		info.descriptorBindingVariableDescriptorCount = C.VkBool32(1)
	}

	if o.RuntimeDescriptorArray {
		info.runtimeDescriptorArray = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceHostQueryResetFeatures describes whether queries can be reset from the host
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostQueryResetFeatures.html
type PhysicalDeviceHostQueryResetFeatures struct {
	// HostQueryReset indicates that hte implementation supports resetting queries from the host
	// with QueryPool.Reset
	HostQueryReset bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceHostQueryResetFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceHostQueryResetFeatures{})))
	}

	info := (*C.VkPhysicalDeviceHostQueryResetFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceHostQueryResetFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceHostQueryResetFeatures)(cDataPointer)
	o.HostQueryReset = info.hostQueryReset != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceHostQueryResetFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceHostQueryResetFeatures{})))
	}

	info := (*C.VkPhysicalDeviceHostQueryResetFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES
	info.pNext = next
	info.hostQueryReset = C.VkBool32(0)

	if o.HostQueryReset {
		info.hostQueryReset = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceImagelessFramebufferFeatures indicates supports for imageless Framebuffer objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImagelessFramebufferFeatures.html
type PhysicalDeviceImagelessFramebufferFeatures struct {
	// ImagelessFramebuffer indicates that the implementation supports specifying the ImageView for
	// attachments at RenderPass begin time via RenderPassAttachmentBeginInfo
	ImagelessFramebuffer bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceImagelessFramebufferFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceImagelessFramebufferFeatures{})))
	}

	info := (*C.VkPhysicalDeviceImagelessFramebufferFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceImagelessFramebufferFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceImagelessFramebufferFeatures)(cDataPointer)

	o.ImagelessFramebuffer = info.imagelessFramebuffer != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceImagelessFramebufferFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceImagelessFramebufferFeatures{})))
	}

	info := (*C.VkPhysicalDeviceImagelessFramebufferFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES
	info.pNext = next
	info.imagelessFramebuffer = C.VkBool32(0)

	if o.ImagelessFramebuffer {
		info.imagelessFramebuffer = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceScalarBlockLayoutFeatures indicates support for scalar block layouts
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceScalarBlockLayoutFeatures.html
type PhysicalDeviceScalarBlockLayoutFeatures struct {
	// ScalarBlockLayout indicates that the implementation supports the layout of resource blocks
	// in shaders using scalar alignment
	ScalarBlockLayout bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceScalarBlockLayoutFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceScalarBlockLayoutFeatures{})))
	}

	info := (*C.VkPhysicalDeviceScalarBlockLayoutFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceScalarBlockLayoutFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceScalarBlockLayoutFeatures)(cDataPointer)

	o.ScalarBlockLayout = info.scalarBlockLayout != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceScalarBlockLayoutFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceScalarBlockLayoutFeatures{})))
	}

	info := (*C.VkPhysicalDeviceScalarBlockLayoutFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES
	info.pNext = next
	info.scalarBlockLayout = C.VkBool32(0)

	if o.ScalarBlockLayout {
		info.scalarBlockLayout = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceSeparateDepthStencilLayoutsFeatures describes whether the implementation
// can do depth and stencil Image barriers separately
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures.html
type PhysicalDeviceSeparateDepthStencilLayoutsFeatures struct {
	// SeparateDepthStencilLayouts indicates whether the implementation supports an
	// ImageMemoryBarrier for a depth/stencil Image with only one of core1_0.ImageAspectDepth or
	// core1_0.ImageAspectStencil, and whether ImageLayoutDepthAttachmentOptimal,
	// ImageLayoutDepthReadOnlyOptimal, ImageLayoutStencilAttachmentOptimal, or
	// ImageLayoutStencilReadOnlyOptimal can be used
	SeparateDepthStencilLayouts bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceSeparateDepthStencilLayoutsFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderAtomicInt64Features{})))
	}

	info := (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSeparateDepthStencilLayoutsFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures)(cDataPointer)

	o.SeparateDepthStencilLayouts = info.separateDepthStencilLayouts != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceSeparateDepthStencilLayoutsFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderAtomicInt64Features{})))
	}

	info := (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES
	info.pNext = next
	info.separateDepthStencilLayouts = C.VkBool32(0)

	if o.SeparateDepthStencilLayouts {
		info.separateDepthStencilLayouts = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceShaderAtomicInt64Features describes features supported by khr_shader_atomic_int64
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicInt64Features.html
type PhysicalDeviceShaderAtomicInt64Features struct {
	// ShaderBufferInt64Atomics indicates whether shaders can perform 64-bit unsigned and signed
	// integer atomic operations on Buffer objects
	ShaderBufferInt64Atomics bool
	// ShaderSharedInt64Atomics indicates whether shaders can 64-bit unsigned and signed integer
	// atomic operations on shared memory
	ShaderSharedInt64Atomics bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceShaderAtomicInt64Features) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderAtomicInt64Features{})))
	}

	info := (*C.VkPhysicalDeviceShaderAtomicInt64Features)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES
	info.pNext = next
	info.shaderBufferInt64Atomics = C.VkBool32(0)
	info.shaderSharedInt64Atomics = C.VkBool32(0)

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderAtomicInt64Features) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderAtomicInt64Features)(cDataPointer)

	o.ShaderBufferInt64Atomics = info.shaderBufferInt64Atomics != C.VkBool32(0)
	o.ShaderSharedInt64Atomics = info.shaderSharedInt64Atomics != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceShaderAtomicInt64Features) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderAtomicInt64Features{})))
	}

	info := (*C.VkPhysicalDeviceShaderAtomicInt64Features)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES
	info.pNext = next
	info.shaderBufferInt64Atomics = C.VkBool32(0)
	info.shaderSharedInt64Atomics = C.VkBool32(0)

	if o.ShaderBufferInt64Atomics {
		info.shaderBufferInt64Atomics = C.VkBool32(1)
	}

	if o.ShaderSharedInt64Atomics {
		info.shaderSharedInt64Atomics = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceShaderFloat16Int8Features describes features supported by khr_shader_float16_int8
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderFloat16Int8Features.html
type PhysicalDeviceShaderFloat16Int8Features struct {
	// ShaderFloat16 indicates whether 16-bit floats (halfs) are supported in shader code
	ShaderFloat16 bool
	// ShaderInt8 indicates whether 8-bit integer (signed and unsigned) are supported in
	// shader code
	ShaderInt8 bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceShaderFloat16Int8Features) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderFloat16Int8Features{})))
	}

	info := (*C.VkPhysicalDeviceShaderFloat16Int8Features)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderFloat16Int8Features) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderFloat16Int8Features)(cDataPointer)

	o.ShaderFloat16 = info.shaderFloat16 != C.VkBool32(0)
	o.ShaderInt8 = info.shaderInt8 != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceShaderFloat16Int8Features) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderFloat16Int8Features{})))
	}

	info := (*C.VkPhysicalDeviceShaderFloat16Int8Features)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES
	info.pNext = next
	info.shaderFloat16 = C.VkBool32(0)
	info.shaderInt8 = C.VkBool32(0)

	if o.ShaderFloat16 {
		info.shaderFloat16 = C.VkBool32(1)
	}

	if o.ShaderInt8 {
		info.shaderInt8 = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceShaderSubgroupExtendedTypesFeatures describes the extended types subgroups
// support feature for an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures.html
type PhysicalDeviceShaderSubgroupExtendedTypesFeatures struct {
	// ShaderSubgroupExtendedTypes specifies whether subgroup operations can use 8-bit integer,
	// 16-bit integer, 64-bit integer, 16-bit floating-point, and vectors of these types
	// in group operations with subgroup scope, if the implementation supports the types
	ShaderSubgroupExtendedTypes bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceShaderSubgroupExtendedTypesFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures{})))
	}

	info := (*C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderSubgroupExtendedTypesFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures)(cDataPointer)

	o.ShaderSubgroupExtendedTypes = info.shaderSubgroupExtendedTypes != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceShaderSubgroupExtendedTypesFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures{})))
	}

	info := (*C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES
	info.pNext = next
	info.shaderSubgroupExtendedTypes = C.VkBool32(0)

	if o.ShaderSubgroupExtendedTypes {
		info.shaderSubgroupExtendedTypes = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceTimelineSemaphoreFeatures describes timeline Semaphore features that can be
// supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTimelineSemaphoreFeatures.html
type PhysicalDeviceTimelineSemaphoreFeatures struct {
	// TimelineSemaphore indicates whether Semaphore objects created with a SemaphoreType
	// of SemaphoreTypeTimeline are supported
	TimelineSemaphore bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceTimelineSemaphoreFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceTimelineSemaphoreFeatures{})))
	}

	info := (*C.VkPhysicalDeviceTimelineSemaphoreFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceTimelineSemaphoreFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceTimelineSemaphoreFeatures)(cDataPointer)

	o.TimelineSemaphore = info.timelineSemaphore != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceTimelineSemaphoreFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceTimelineSemaphoreFeatures{})))
	}

	info := (*C.VkPhysicalDeviceTimelineSemaphoreFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES
	info.pNext = next
	info.timelineSemaphore = C.VkBool32(0)

	if o.TimelineSemaphore {
		info.timelineSemaphore = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceUniformBufferStandardLayoutFeatures indicates support for std430-like
// packing in uniform Buffer objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceUniformBufferStandardLayoutFeatures.html
type PhysicalDeviceUniformBufferStandardLayoutFeatures struct {
	// UniformBufferStandardLayout indicates that the implementation supports the same layouts
	// for uniform Buffer objects as for storage and other kinds of Buffer objects
	UniformBufferStandardLayout bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceUniformBufferStandardLayoutFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceUniformBufferStandardLayoutFeatures{})))
	}

	info := (*C.VkPhysicalDeviceUniformBufferStandardLayoutFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceUniformBufferStandardLayoutFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceUniformBufferStandardLayoutFeatures)(cDataPointer)

	o.UniformBufferStandardLayout = info.uniformBufferStandardLayout != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceUniformBufferStandardLayoutFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceUniformBufferStandardLayoutFeatures{})))
	}

	info := (*C.VkPhysicalDeviceUniformBufferStandardLayoutFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES
	info.pNext = next
	info.uniformBufferStandardLayout = C.VkBool32(0)

	if o.UniformBufferStandardLayout {
		info.uniformBufferStandardLayout = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceVulkanMemoryModelFeatures describes features supported by the memory model
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkanMemoryModelFeatures.html
type PhysicalDeviceVulkanMemoryModelFeatures struct {
	// VulkanMemoryModel indicates whether the Vulkan Memory Model is supported
	VulkanMemoryModel bool
	// VulkanMemoryModelDeviceScope indicates whether the Vulkan Memory Model can use Device
	// scope synchronization
	VulkanMemoryModelDeviceScope bool
	// VulkanMemoryModelAvailabilityVisibilityChains indicates whether the Vulkan Memory Model
	// can use available and visibility chains with more than one element
	VulkanMemoryModelAvailabilityVisibilityChains bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceVulkanMemoryModelFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVulkanMemoryModelFeatures{})))
	}

	info := (*C.VkPhysicalDeviceVulkanMemoryModelFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceVulkanMemoryModelFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceVulkanMemoryModelFeatures)(cDataPointer)

	o.VulkanMemoryModel = info.vulkanMemoryModel != C.VkBool32(0)
	o.VulkanMemoryModelDeviceScope = info.vulkanMemoryModelDeviceScope != C.VkBool32(0)
	o.VulkanMemoryModelAvailabilityVisibilityChains = info.vulkanMemoryModelAvailabilityVisibilityChains != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceVulkanMemoryModelFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVulkanMemoryModelFeatures{})))
	}

	info := (*C.VkPhysicalDeviceVulkanMemoryModelFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES
	info.pNext = next
	info.vulkanMemoryModel = C.VkBool32(0)
	info.vulkanMemoryModelDeviceScope = C.VkBool32(0)
	info.vulkanMemoryModelAvailabilityVisibilityChains = C.VkBool32(0)

	if o.VulkanMemoryModel {
		info.vulkanMemoryModel = C.VkBool32(1)
	}

	if o.VulkanMemoryModelDeviceScope {
		info.vulkanMemoryModelDeviceScope = C.VkBool32(1)
	}

	if o.VulkanMemoryModelAvailabilityVisibilityChains {
		info.vulkanMemoryModelAvailabilityVisibilityChains = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceVulkan11Features describes the Vulkan 1.1 features that can be supported
// by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan11Features.html
type PhysicalDeviceVulkan11Features struct {
	// StorageBuffer16BitAccess specifies whether objects in the StorageBuffer, ShaderRecordBufferKHR,
	// or PhysicalStorageBuffer storage class with the Block decoration can have 16-bit integer
	// and 16-bit floating-point members
	StorageBuffer16BitAccess bool
	// UniformAndStorageBuffer16BitAccess specifies whether objects in the Uniform storage class
	// with the Block decoration can have 16-bit integer and 16-bit floating-point members
	UniformAndStorageBuffer16BitAccess bool
	// StoragePushConstant16 specifies whether objects in the PushConstant storage class can have
	// 16-bit integer and 16-bit floating-point members
	StoragePushConstant16 bool
	// StorageInputOutput16 specifies whether objects in the Input and Output storage classes can
	// have 16-bit integer and 16-bit floating-point members
	StorageInputOutput16 bool
	// Multiview specifies whether the implementation supports multiview rendering within a
	// render pass. If this feature is not enabled, the view mask of each subpass must always be
	// zero
	Multiview bool
	// MultiviewGeometryShader specifies whether the implementation supports multiview rendering
	// within a RenderPass, with geometry shaders
	MultiviewGeometryShader bool
	// MultiviewTessellationShader specifies whether the implementation supports multiview
	// rendering within a RenderPass, with tessellation shaders
	MultiviewTessellationShader bool
	// VariablePointersStorageBuffer specifies whether the implementation supports the SPIR-V
	// VariablePointersStorageBuffer capability
	VariablePointersStorageBuffer bool
	// VariablePointers specifies whether the implementation supports the SPIR-V
	// VariablePointers capability
	VariablePointers bool
	// ProtectedMemory specifies whether protected memory is supported
	ProtectedMemory bool
	// SamplerYcbcrConversion specifies whether the implementation supports SamplerYcbcrConversion
	SamplerYcbcrConversion bool
	// ShaderDrawParameters specifies whether the implementation supports the SPIR-V
	// DrawParameters capability
	ShaderDrawParameters bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceVulkan11Features) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceVulkan11Features)
	}

	info := (*C.VkPhysicalDeviceVulkan11Features)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_FEATURES
	info.pNext = next
	info.storageBuffer16BitAccess = C.VkBool32(0)
	info.uniformAndStorageBuffer16BitAccess = C.VkBool32(0)
	info.storagePushConstant16 = C.VkBool32(0)
	info.storageInputOutput16 = C.VkBool32(0)
	info.multiview = C.VkBool32(0)
	info.multiviewGeometryShader = C.VkBool32(0)
	info.multiviewTessellationShader = C.VkBool32(0)
	info.variablePointersStorageBuffer = C.VkBool32(0)
	info.variablePointers = C.VkBool32(0)
	info.protectedMemory = C.VkBool32(0)
	info.samplerYcbcrConversion = C.VkBool32(0)
	info.shaderDrawParameters = C.VkBool32(0)

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceVulkan11Features) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceVulkan11Features)(cDataPointer)

	o.StorageBuffer16BitAccess = info.storageBuffer16BitAccess != C.VkBool32(0)
	o.UniformAndStorageBuffer16BitAccess = info.uniformAndStorageBuffer16BitAccess != C.VkBool32(0)
	o.StoragePushConstant16 = info.storagePushConstant16 != C.VkBool32(0)
	o.StorageInputOutput16 = info.storageInputOutput16 != C.VkBool32(0)
	o.Multiview = info.multiview != C.VkBool32(0)
	o.MultiviewGeometryShader = info.multiviewGeometryShader != C.VkBool32(0)
	o.MultiviewTessellationShader = info.multiviewTessellationShader != C.VkBool32(0)
	o.VariablePointersStorageBuffer = info.variablePointersStorageBuffer != C.VkBool32(0)
	o.VariablePointers = info.variablePointers != C.VkBool32(0)
	o.ProtectedMemory = info.protectedMemory != C.VkBool32(0)
	o.SamplerYcbcrConversion = info.samplerYcbcrConversion != C.VkBool32(0)
	o.ShaderDrawParameters = info.shaderDrawParameters != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceVulkan11Features) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceVulkan11Features)
	}

	info := (*C.VkPhysicalDeviceVulkan11Features)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_FEATURES
	info.pNext = next
	info.storageBuffer16BitAccess = C.VkBool32(0)
	info.uniformAndStorageBuffer16BitAccess = C.VkBool32(0)
	info.storagePushConstant16 = C.VkBool32(0)
	info.storageInputOutput16 = C.VkBool32(0)
	info.multiview = C.VkBool32(0)
	info.multiviewGeometryShader = C.VkBool32(0)
	info.multiviewTessellationShader = C.VkBool32(0)
	info.variablePointersStorageBuffer = C.VkBool32(0)
	info.variablePointers = C.VkBool32(0)
	info.protectedMemory = C.VkBool32(0)
	info.samplerYcbcrConversion = C.VkBool32(0)
	info.shaderDrawParameters = C.VkBool32(0)

	if o.StorageBuffer16BitAccess {
		info.storageBuffer16BitAccess = C.VkBool32(1)
	}

	if o.UniformAndStorageBuffer16BitAccess {
		info.uniformAndStorageBuffer16BitAccess = C.VkBool32(0)
	}

	if o.StoragePushConstant16 {
		info.storagePushConstant16 = C.VkBool32(1)
	}

	if o.StorageInputOutput16 {
		info.storageInputOutput16 = C.VkBool32(1)
	}

	if o.Multiview {
		info.multiview = C.VkBool32(1)
	}

	if o.MultiviewGeometryShader {
		info.multiviewGeometryShader = C.VkBool32(1)
	}

	if o.MultiviewTessellationShader {
		info.multiviewTessellationShader = C.VkBool32(1)
	}

	if o.VariablePointersStorageBuffer {
		info.variablePointersStorageBuffer = C.VkBool32(1)
	}

	if o.VariablePointers {
		info.variablePointers = C.VkBool32(1)
	}

	if o.ProtectedMemory {
		info.protectedMemory = C.VkBool32(1)
	}

	if o.SamplerYcbcrConversion {
		info.samplerYcbcrConversion = C.VkBool32(1)
	}

	if o.ShaderDrawParameters {
		info.shaderDrawParameters = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

// PhysicalDeviceVulkan12Features describes the Vulkan 1.2 features that can be supported by
// an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Features.html
type PhysicalDeviceVulkan12Features struct {
	// SamplerMirrorClampToEdge indicates whether the implementation supports the
	// SamplerAddressModeMirrorClampToEdge sampler address mode
	SamplerMirrorClampToEdge bool
	// DrawIndirectCount indicates whether the implementation supports the
	// CommandBuffer.CmdDrawIndirectCount and CommandBuffer.CmdDrawIndexedIndirectCount functions
	DrawIndirectCount bool
	// StorageBuffer8BitAccess indicates whether objects in the StorageBuffer,
	// ShaderRecordBufferKHR, or PhysicalStorageBuffer storage class with the Block decoration
	// can have 8-bit integer members
	StorageBuffer8BitAccess bool
	// UniformAndStorageBuffer8BitAccess indicates whether objects in the Uniform storage class
	// with the Block decoration can have 8-bit integer members
	UniformAndStorageBuffer8BitAccess bool
	// StoragePushConstant8 indicates whether objects in the PushConstant storage class can
	// have 8-bit integer members
	StoragePushConstant8 bool
	// ShaderBufferInt64Atomics indicates whether shaders can perform 64-bit unsigned and signed
	// integer atomic operations on Buffer objects
	ShaderBufferInt64Atomics bool
	// ShaderSharedInt64Atomics indicates whether shaders can perform 64-bit unsigned and signed
	// integer atomic operations on shared memory
	ShaderSharedInt64Atomics bool
	// ShaderFloat16 indicates whether 16-bit floats (halfs) are supported in shader code
	ShaderFloat16 bool
	// ShaderInt8 indicates whether 8-bit integers (signed and unsigned) are supported in shader
	// code
	ShaderInt8 bool
	// DescriptorIndexing indicates whether the implementation supports the minimum set of
	// descriptor indexing features as described in the Feature Requirements section
	DescriptorIndexing bool

	// ShaderInputAttachmentArrayDynamicIndexing indicates whether arrays of input attachments
	// can be indexed by dynamically uniform integer expressions in shader code
	ShaderInputAttachmentArrayDynamicIndexing bool
	// ShaderUniformTexelBufferArrayDynamicIndexing indicates whether arrays of uniform texel
	// Buffer objects can be indexed by dynamically uniform integer expressions in shader code
	ShaderUniformTexelBufferArrayDynamicIndexing bool
	// ShaderStorageTexelBufferArrayDynamicIndexing indicates whether arrays of storage texel
	// Buffer objects can be indexed by dynamically uniform integer expressions in shader code
	ShaderStorageTexelBufferArrayDynamicIndexing bool
	// ShaderUniformBufferArrayNonUniformIndexing indicates whether arrays of uniform Buffer objects
	// can be indexed by non-uniform integer expressions in shader code
	ShaderUniformBufferArrayNonUniformIndexing bool
	// ShaderSampledImageArrayNonUniformIndexing indicates whether arrays of Sampler objects or sampled
	// Image objects can be indexed by non-uniform integer expressions in shader code
	ShaderSampledImageArrayNonUniformIndexing bool
	// ShaderStorageBufferArrayNonUniformIndexing indicates whether arrays of storage Buffer objects can
	// be indexed by non-uniform integer expressions in shader code
	ShaderStorageBufferArrayNonUniformIndexing bool
	// ShaderStorageImageArrayNonUniformIndexing indicates whether arrays of storage Image objects can
	// be indexed by non-uniform integer expressions in shader code
	ShaderStorageImageArrayNonUniformIndexing bool
	// ShaderInputAttachmentArrayNonUniformIndexing indicates whether arrays of input attachments
	// can be indexed by non-uniform integer expressions in shader code
	ShaderInputAttachmentArrayNonUniformIndexing bool
	// ShaderUniformTexelBufferArrayNonUniformIndexing indicates whether arrays of uniform texel
	// Buffer objects can be indexed by non-uniform integer expressions in shader code
	ShaderUniformTexelBufferArrayNonUniformIndexing bool
	// ShaderStorageTexelBufferArrayNonUniformIndexing indicates whether arrays of storage texel
	// Buffer objects can be indexed by non-uniform integer expressions in shader code
	ShaderStorageTexelBufferArrayNonUniformIndexing bool

	// DescriptorBindingUniformBufferUpdateAfterBind indicates whether the implementation
	// supports updating uniform Buffer descriptors after a set is bound
	DescriptorBindingUniformBufferUpdateAfterBind bool
	// DescriptorBindingSampledImageUpdateAfterBind indicates whether the implementation supports
	// updating sampled Image descriptors after a set is bound
	DescriptorBindingSampledImageUpdateAfterBind bool
	// DescriptorBindingStorageImageUpdateAfterBind indicates whether the implementation supports
	// updating storage Image descriptors after a set is bound
	DescriptorBindingStorageImageUpdateAfterBind bool
	// DescriptorBindingStorageBufferUpdateAfterBind indicates whether the implementation supports
	// updating storage Buffer descriptors after a set is bound
	DescriptorBindingStorageBufferUpdateAfterBind bool
	// DescriptorBindingUniformTexelBufferUpdateAfterBind indicates whether the implementation
	// supports updating uniform texel Buffer descriptors after a set is bound
	DescriptorBindingUniformTexelBufferUpdateAfterBind bool
	// DescriptorBindingStorageTexelBufferUpdateAfterBind indicates whether the implementation
	// supports updating storage texel buffer descriptors after a set is bound
	DescriptorBindingStorageTexelBufferUpdateAfterBind bool

	// DescriptorBindingUpdateUnusedWhilePending indicates whether the implementation supports
	// updating descriptors while the set is in use
	DescriptorBindingUpdateUnusedWhilePending bool
	// DescriptorBindingPartiallyBound indicates whether the implementation supports statically
	// using a DescriptorSet binding in which some descriptors are not valid
	DescriptorBindingPartiallyBound bool
	// DescriptorBindingVariableDescriptorCount indicates whether the implementation supports
	// DescriptorSet with a variable-sized last binding
	DescriptorBindingVariableDescriptorCount bool

	// RuntimeDescriptorArray indicates whether the implementation supports the SPIR-V
	// RuntimeDescriptorArray capability
	RuntimeDescriptorArray bool
	// SamplerFilterMinmax indicates whether the implementation supports a minimum set of
	// required formats supporting min/max filtering as defined by the
	// filterMinmaxSingleComponentFormats property minimum requirements
	SamplerFilterMinmax bool
	// ScalarBlockLayout indicates that the implementation supports the layout of resource blocks
	// in shaders using scalar alignment
	ScalarBlockLayout bool
	// ImagelessFramebuffer indicates that the implementation supports specifying the ImageView for
	// attachments at RenderPass begin time via RenderPassAttachmentBeginInfo
	ImagelessFramebuffer bool
	// UniformBufferStandardLayout indicates that the implementation supports the same layouts for
	// uniform Buffer objects as for storage and other kinds of Buffer objects
	UniformBufferStandardLayout bool
	// ShaderSubgroupExtendedTypes is a boolean specifying whether subgroup operations can use
	// 8-bit integer, 16-bit integer, 64-bit integer, 16-bit floating-point, and vectors of these
	// types in group operations with subgroup scope, if the implementation supports the types
	ShaderSubgroupExtendedTypes bool
	// SeparateDepthStencilLayouts indicates whether the implementation supports an ImageMemoryBarrier
	// for a depth/stencil Image with only one of core1_0.ImageAspectDepth or core1_0.ImageAspectStencil,
	// and whether ImageLayoutDepthAttachmentOptimal, ImageLayoutDepthReadOnlyOptimal,
	// ImageLayoutStencilAttachmentOptimal, ImageLayoutStencilReadOnlyOptimal can be used
	SeparateDepthStencilLayouts bool
	// HostQueryReset indicates that the implementation supports resetting queries from the host with
	// QueryPool.Reset
	HostQueryReset bool
	// TimelineSemaphore indicates whether Semaphore objects created with a SemaphoreType of
	// SemaphoreTypeTimeline are supported
	TimelineSemaphore bool
	// BufferDeviceAddress indicates that the implementation supports accessing Buffer memory in shaders
	// as storage Buffer objects via an address queries from Device.GetBufferDeviceAddress
	BufferDeviceAddress bool
	// BufferDeviceAddressCaptureReplay indicates that the implementation supports saving and
	// reusing Buffer and Device addresses, e.g. for trace capture and replay
	BufferDeviceAddressCaptureReplay bool
	// BufferDeviceAddressMultiDevice indicates that the implementation supports the
	// BufferDeviceAddress, RayTracingPipeline and RayQuery features for logical Device objects
	// created with multiple PhysicalDevice objects
	BufferDeviceAddressMultiDevice bool
	// VulkanMemoryModel indicates whether the Vulkan Memory Model is supported
	VulkanMemoryModel bool
	// VulkanMemoryModelDeviceScope indicates whether the Vulkan Memory Model can use Device
	// scope synchronization
	VulkanMemoryModelDeviceScope bool
	// VulkanMemoryModelAvailabilityVisibilityChains indicates whether the Vulkan Memory Model
	// can use availability and visibility chains with more than one element
	VulkanMemoryModelAvailabilityVisibilityChains bool
	// ShaderOutputViewportIndex indicates whether the implementation supports the
	// ShaderViewportIndex SPIR-V capability enabling variables decorated with the ViewportIndex
	// built-in to be exported from vertex or tessellation evaluation shaders
	ShaderOutputViewportIndex bool
	// ShaderOutputLayer indicates whether the implementation supports the ShaderLayer SPIR-V
	// capability enabling variables decorated with the Layer built-in to be exported from vertex
	// or tessellation evaluation shaders
	ShaderOutputLayer bool
	// SubgroupBroadcastDynamicID indicates whether the "Id" operand of OpGroupNonUniformBroadcast
	// can be dynamically uniform within a subgroup, and whether the "Index" operand of
	// OpGroupNonUniformQuadBroadcast can be dynamically uniform within the derivative group
	SubgroupBroadcastDynamicID bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceVulkan12Features) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceVulkan12Features)
	}

	info := (*C.VkPhysicalDeviceVulkan12Features)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceVulkan12Features) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceVulkan12Features)(cDataPointer)

	o.SamplerMirrorClampToEdge = info.samplerMirrorClampToEdge != C.VkBool32(0)
	o.DrawIndirectCount = info.drawIndirectCount != C.VkBool32(0)
	o.StorageBuffer8BitAccess = info.storageBuffer8BitAccess != C.VkBool32(0)
	o.UniformAndStorageBuffer8BitAccess = info.uniformAndStorageBuffer8BitAccess != C.VkBool32(0)
	o.StoragePushConstant8 = info.storagePushConstant8 != C.VkBool32(0)
	o.ShaderBufferInt64Atomics = info.shaderBufferInt64Atomics != C.VkBool32(0)
	o.ShaderSharedInt64Atomics = info.shaderSharedInt64Atomics != C.VkBool32(0)
	o.ShaderFloat16 = info.shaderFloat16 != C.VkBool32(0)
	o.ShaderInt8 = info.shaderInt8 != C.VkBool32(0)
	o.DescriptorIndexing = info.descriptorIndexing != C.VkBool32(0)

	o.ShaderInputAttachmentArrayDynamicIndexing = info.shaderInputAttachmentArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderUniformTexelBufferArrayDynamicIndexing = info.shaderUniformTexelBufferArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderStorageTexelBufferArrayDynamicIndexing = info.shaderStorageTexelBufferArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderUniformBufferArrayNonUniformIndexing = info.shaderUniformBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderSampledImageArrayNonUniformIndexing = info.shaderSampledImageArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageBufferArrayNonUniformIndexing = info.shaderStorageBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageImageArrayNonUniformIndexing = info.shaderStorageImageArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderInputAttachmentArrayNonUniformIndexing = info.shaderInputAttachmentArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderUniformTexelBufferArrayNonUniformIndexing = info.shaderUniformTexelBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageTexelBufferArrayNonUniformIndexing = info.shaderStorageTexelBufferArrayNonUniformIndexing != C.VkBool32(0)

	o.DescriptorBindingUniformBufferUpdateAfterBind = info.descriptorBindingUniformBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingSampledImageUpdateAfterBind = info.descriptorBindingSampledImageUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageImageUpdateAfterBind = info.descriptorBindingStorageImageUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageBufferUpdateAfterBind = info.descriptorBindingStorageBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingUniformTexelBufferUpdateAfterBind = info.descriptorBindingUniformTexelBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageTexelBufferUpdateAfterBind = info.descriptorBindingStorageTexelBufferUpdateAfterBind != C.VkBool32(0)

	o.DescriptorBindingUpdateUnusedWhilePending = info.descriptorBindingUpdateUnusedWhilePending != C.VkBool32(0)
	o.DescriptorBindingPartiallyBound = info.descriptorBindingPartiallyBound != C.VkBool32(0)
	o.DescriptorBindingVariableDescriptorCount = info.descriptorBindingVariableDescriptorCount != C.VkBool32(0)

	o.RuntimeDescriptorArray = info.runtimeDescriptorArray != C.VkBool32(0)
	o.SamplerFilterMinmax = info.samplerFilterMinmax != C.VkBool32(0)
	o.ScalarBlockLayout = info.scalarBlockLayout != C.VkBool32(0)
	o.ImagelessFramebuffer = info.imagelessFramebuffer != C.VkBool32(0)
	o.UniformBufferStandardLayout = info.uniformBufferStandardLayout != C.VkBool32(0)
	o.ShaderSubgroupExtendedTypes = info.shaderSubgroupExtendedTypes != C.VkBool32(0)
	o.SeparateDepthStencilLayouts = info.separateDepthStencilLayouts != C.VkBool32(0)
	o.HostQueryReset = info.hostQueryReset != C.VkBool32(0)
	o.TimelineSemaphore = info.timelineSemaphore != C.VkBool32(0)
	o.BufferDeviceAddress = info.bufferDeviceAddress != C.VkBool32(0)
	o.BufferDeviceAddressCaptureReplay = info.bufferDeviceAddressCaptureReplay != C.VkBool32(0)
	o.BufferDeviceAddressMultiDevice = info.bufferDeviceAddressMultiDevice != C.VkBool32(0)
	o.VulkanMemoryModel = info.vulkanMemoryModel != C.VkBool32(0)
	o.VulkanMemoryModelDeviceScope = info.vulkanMemoryModelDeviceScope != C.VkBool32(0)
	o.VulkanMemoryModelAvailabilityVisibilityChains = info.vulkanMemoryModelAvailabilityVisibilityChains != C.VkBool32(0)
	o.ShaderOutputViewportIndex = info.shaderOutputViewportIndex != C.VkBool32(0)
	o.ShaderOutputLayer = info.shaderOutputLayer != C.VkBool32(0)
	o.SubgroupBroadcastDynamicID = info.subgroupBroadcastDynamicId != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceVulkan12Features) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceVulkan12Features)
	}

	info := (*C.VkPhysicalDeviceVulkan12Features)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_FEATURES
	info.pNext = next
	info.samplerMirrorClampToEdge = C.VkBool32(0)
	info.drawIndirectCount = C.VkBool32(0)
	info.storageBuffer8BitAccess = C.VkBool32(0)
	info.uniformAndStorageBuffer8BitAccess = C.VkBool32(0)
	info.storagePushConstant8 = C.VkBool32(0)
	info.shaderBufferInt64Atomics = C.VkBool32(0)
	info.shaderSharedInt64Atomics = C.VkBool32(0)
	info.shaderFloat16 = C.VkBool32(0)
	info.shaderInt8 = C.VkBool32(0)
	info.descriptorIndexing = C.VkBool32(0)

	info.shaderInputAttachmentArrayDynamicIndexing = C.VkBool32(0)
	info.shaderUniformTexelBufferArrayDynamicIndexing = C.VkBool32(0)
	info.shaderStorageTexelBufferArrayDynamicIndexing = C.VkBool32(0)
	info.shaderUniformBufferArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderSampledImageArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderStorageBufferArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderStorageImageArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderInputAttachmentArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderUniformTexelBufferArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderStorageTexelBufferArrayNonUniformIndexing = C.VkBool32(0)

	info.descriptorBindingUniformBufferUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingSampledImageUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingStorageImageUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingStorageBufferUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingUniformTexelBufferUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingStorageTexelBufferUpdateAfterBind = C.VkBool32(0)

	info.descriptorBindingUpdateUnusedWhilePending = C.VkBool32(0)
	info.descriptorBindingPartiallyBound = C.VkBool32(0)
	info.descriptorBindingVariableDescriptorCount = C.VkBool32(0)

	info.runtimeDescriptorArray = C.VkBool32(0)
	info.samplerFilterMinmax = C.VkBool32(0)
	info.scalarBlockLayout = C.VkBool32(0)
	info.imagelessFramebuffer = C.VkBool32(0)
	info.uniformBufferStandardLayout = C.VkBool32(0)
	info.shaderSubgroupExtendedTypes = C.VkBool32(0)
	info.separateDepthStencilLayouts = C.VkBool32(0)
	info.hostQueryReset = C.VkBool32(0)
	info.timelineSemaphore = C.VkBool32(0)
	info.bufferDeviceAddress = C.VkBool32(0)
	info.bufferDeviceAddressCaptureReplay = C.VkBool32(0)
	info.bufferDeviceAddressMultiDevice = C.VkBool32(0)
	info.vulkanMemoryModel = C.VkBool32(0)
	info.vulkanMemoryModelDeviceScope = C.VkBool32(0)
	info.vulkanMemoryModelAvailabilityVisibilityChains = C.VkBool32(0)
	info.shaderOutputViewportIndex = C.VkBool32(0)
	info.shaderOutputLayer = C.VkBool32(0)
	info.subgroupBroadcastDynamicId = C.VkBool32(0)

	if o.SamplerMirrorClampToEdge {
		info.samplerMirrorClampToEdge = C.VkBool32(1)
	}
	if o.DrawIndirectCount {
		info.drawIndirectCount = C.VkBool32(1)
	}
	if o.StorageBuffer8BitAccess {
		info.storageBuffer8BitAccess = C.VkBool32(1)
	}
	if o.UniformAndStorageBuffer8BitAccess {
		info.uniformAndStorageBuffer8BitAccess = C.VkBool32(1)
	}
	if o.StoragePushConstant8 {
		info.storagePushConstant8 = C.VkBool32(1)
	}
	if o.ShaderBufferInt64Atomics {
		info.shaderBufferInt64Atomics = C.VkBool32(1)
	}
	if o.ShaderSharedInt64Atomics {
		info.shaderSharedInt64Atomics = C.VkBool32(1)
	}
	if o.ShaderFloat16 {
		info.shaderFloat16 = C.VkBool32(1)
	}
	if o.ShaderInt8 {
		info.shaderInt8 = C.VkBool32(1)
	}
	if o.DescriptorIndexing {
		info.descriptorIndexing = C.VkBool32(1)
	}

	if o.ShaderInputAttachmentArrayDynamicIndexing {
		info.shaderInputAttachmentArrayDynamicIndexing = C.VkBool32(1)
	}
	if o.ShaderUniformTexelBufferArrayDynamicIndexing {
		info.shaderUniformTexelBufferArrayDynamicIndexing = C.VkBool32(1)
	}
	if o.ShaderStorageTexelBufferArrayDynamicIndexing {
		info.shaderStorageTexelBufferArrayDynamicIndexing = C.VkBool32(1)
	}
	if o.ShaderUniformBufferArrayNonUniformIndexing {
		info.shaderUniformBufferArrayNonUniformIndexing = C.VkBool32(1)
	}
	if o.ShaderSampledImageArrayNonUniformIndexing {
		info.shaderSampledImageArrayNonUniformIndexing = C.VkBool32(1)
	}
	if o.ShaderStorageBufferArrayNonUniformIndexing {
		info.shaderStorageBufferArrayNonUniformIndexing = C.VkBool32(1)
	}
	if o.ShaderStorageImageArrayNonUniformIndexing {
		info.shaderStorageImageArrayNonUniformIndexing = C.VkBool32(1)
	}
	if o.ShaderInputAttachmentArrayNonUniformIndexing {
		info.shaderInputAttachmentArrayNonUniformIndexing = C.VkBool32(1)
	}
	if o.ShaderUniformTexelBufferArrayNonUniformIndexing {
		info.shaderUniformTexelBufferArrayNonUniformIndexing = C.VkBool32(1)
	}
	if o.ShaderStorageTexelBufferArrayNonUniformIndexing {
		info.shaderStorageTexelBufferArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.DescriptorBindingUniformBufferUpdateAfterBind {
		info.descriptorBindingUniformBufferUpdateAfterBind = C.VkBool32(1)
	}
	if o.DescriptorBindingSampledImageUpdateAfterBind {
		info.descriptorBindingSampledImageUpdateAfterBind = C.VkBool32(1)
	}
	if o.DescriptorBindingStorageImageUpdateAfterBind {
		info.descriptorBindingStorageImageUpdateAfterBind = C.VkBool32(1)
	}
	if o.DescriptorBindingStorageBufferUpdateAfterBind {
		info.descriptorBindingStorageBufferUpdateAfterBind = C.VkBool32(1)
	}
	if o.DescriptorBindingUniformTexelBufferUpdateAfterBind {
		info.descriptorBindingUniformTexelBufferUpdateAfterBind = C.VkBool32(1)
	}
	if o.DescriptorBindingStorageTexelBufferUpdateAfterBind {
		info.descriptorBindingStorageTexelBufferUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingUpdateUnusedWhilePending {
		info.descriptorBindingUpdateUnusedWhilePending = C.VkBool32(1)
	}
	if o.DescriptorBindingPartiallyBound {
		info.descriptorBindingPartiallyBound = C.VkBool32(1)
	}
	if o.DescriptorBindingVariableDescriptorCount {
		info.descriptorBindingVariableDescriptorCount = C.VkBool32(1)
	}

	if o.RuntimeDescriptorArray {
		info.runtimeDescriptorArray = C.VkBool32(1)
	}
	if o.SamplerFilterMinmax {
		info.samplerFilterMinmax = C.VkBool32(1)
	}
	if o.ScalarBlockLayout {
		info.scalarBlockLayout = C.VkBool32(1)
	}
	if o.ImagelessFramebuffer {
		info.imagelessFramebuffer = C.VkBool32(1)
	}
	if o.UniformBufferStandardLayout {
		info.uniformBufferStandardLayout = C.VkBool32(1)
	}
	if o.ShaderSubgroupExtendedTypes {
		info.shaderSubgroupExtendedTypes = C.VkBool32(1)
	}
	if o.SeparateDepthStencilLayouts {
		info.separateDepthStencilLayouts = C.VkBool32(1)
	}
	if o.HostQueryReset {
		info.hostQueryReset = C.VkBool32(1)
	}
	if o.TimelineSemaphore {
		info.timelineSemaphore = C.VkBool32(1)
	}
	if o.BufferDeviceAddress {
		info.bufferDeviceAddress = C.VkBool32(1)
	}
	if o.BufferDeviceAddressCaptureReplay {
		info.bufferDeviceAddressCaptureReplay = C.VkBool32(1)
	}
	if o.BufferDeviceAddressMultiDevice {
		info.bufferDeviceAddressMultiDevice = C.VkBool32(1)
	}
	if o.VulkanMemoryModel {
		info.vulkanMemoryModel = C.VkBool32(1)
	}
	if o.VulkanMemoryModelDeviceScope {
		info.vulkanMemoryModelDeviceScope = C.VkBool32(1)
	}
	if o.VulkanMemoryModelAvailabilityVisibilityChains {
		info.vulkanMemoryModelAvailabilityVisibilityChains = C.VkBool32(1)
	}
	if o.ShaderOutputViewportIndex {
		info.shaderOutputViewportIndex = C.VkBool32(1)
	}
	if o.ShaderOutputLayer {
		info.shaderOutputLayer = C.VkBool32(1)
	}
	if o.SubgroupBroadcastDynamicID {
		info.subgroupBroadcastDynamicId = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
