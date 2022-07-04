package core1_2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type PhysicalDevice8BitStorageFeatures struct {
	StorageBuffer8BitAccess           bool
	UniformAndStorageBuffer8BitAccess bool
	StoragePushConstant8              bool

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

type PhysicalDeviceBufferAddressFeatures struct {
	BufferDeviceAddress              bool
	BufferDeviceAddressCaptureReplay bool
	BufferDeviceAddressMultiDevice   bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceBufferAddressFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceBufferDeviceAddressFeatures{})))
	}

	info := (*C.VkPhysicalDeviceBufferDeviceAddressFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceBufferAddressFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceBufferDeviceAddressFeatures)(cDataPointer)

	o.BufferDeviceAddress = info.bufferDeviceAddress != C.VkBool32(0)
	o.BufferDeviceAddressCaptureReplay = info.bufferDeviceAddressCaptureReplay != C.VkBool32(0)
	o.BufferDeviceAddressMultiDevice = info.bufferDeviceAddressMultiDevice != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceBufferAddressFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

type PhysicalDeviceDescriptorIndexingFeatures struct {
	ShaderInputAttachmentArrayDynamicIndexing          bool
	ShaderUniformTexelBufferArrayDynamicIndexing       bool
	ShaderStorageTexelBufferArrayDynamicIndexing       bool
	ShaderUniformBufferArrayNonUniformIndexing         bool
	ShaderSampledImageArrayNonUniformIndexing          bool
	ShaderStorageBufferArrayNonUniformIndexing         bool
	ShaderStorageImageArrayNonUniformIndexing          bool
	ShaderInputAttachmentArrayNonUniformIndexing       bool
	ShaderUniformTexelBufferArrayNonUniformIndexing    bool
	ShaderStorageTexelBufferArrayNonUniformIndexing    bool
	DescriptorBindingUniformBufferUpdateAfterBind      bool
	DescriptorBindingSampledImageUpdateAfterBind       bool
	DescriptorBindingStorageImageUpdateAfterBind       bool
	DescriptorBindingStorageBufferUpdateAfterBind      bool
	DescriptorBindingUniformTexelBufferUpdateAfterBind bool
	DescriptorBindingStorageTexelBufferUpdateAfterBind bool
	DescriptorBindingUpdateUnusedWhilePending          bool
	DescriptorBindingPartiallyBound                    bool
	DescriptorBindingVariableDescriptorCount           bool
	RuntimeDescriptorArray                             bool

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

type PhysicalDeviceHostQueryResetFeatures struct {
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

type PhysicalDeviceImagelessFramebufferFeatures struct {
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

type PhysicalDeviceScalarBlockLayoutFeatures struct {
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

type PhysicalDeviceSeparateDepthStencilLayoutsFeatures struct {
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

type PhysicalDeviceShaderAtomicInt64Features struct {
	ShaderBufferInt64Atomics bool
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

type PhysicalDeviceShaderFloat16Int8Features struct {
	ShaderFloat16 bool
	ShaderInt8    bool

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

type PhysicalDeviceShaderSubgroupExtendedTypesFeatures struct {
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

type PhysicalDeviceTimelineSemaphoreFeatures struct {
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

type PhysicalDeviceUniformBufferStandardLayoutFeatures struct {
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

type PhysicalDeviceVulkanMemoryModelFeatures struct {
	VulkanMemoryModel                             bool
	VulkanMemoryModelDeviceScope                  bool
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

type PhysicalDeviceVulkan11Features struct {
	StorageBuffer16BitAccess           bool
	UniformAndStorageBuffer16BitAccess bool
	StoragePushConstant16              bool
	StorageInputOutput16               bool
	Multiview                          bool
	MultiviewGeometryShader            bool
	MultiviewTessellationShader        bool
	VariablePointersStorageBuffer      bool
	VariablePointers                   bool
	ProtectedMemory                    bool
	SamplerYcbcrConversion             bool
	ShaderDrawParameters               bool

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

type PhysicalDeviceVulkan12Features struct {
	SamplerMirrorClampToEdge          bool
	DrawIndirectCount                 bool
	StorageBuffer8BitAccess           bool
	UniformAndStorageBuffer8BitAccess bool
	StoragePushConstant8              bool
	ShaderBufferInt64Atomics          bool
	ShaderSharedInt64Atomics          bool
	ShaderFloat16                     bool
	ShaderInt8                        bool
	DescriptorIndexing                bool

	ShaderInputAttachmentArrayDynamicIndexing       bool
	ShaderUniformTexelBufferArrayDynamicIndexing    bool
	ShaderStorageTexelBufferArrayDynamicIndexing    bool
	ShaderUniformBufferArrayNonUniformIndexing      bool
	ShaderSampledImageArrayNonUniformIndexing       bool
	ShaderStorageBufferArrayNonUniformIndexing      bool
	ShaderStorageImageArrayNonUniformIndexing       bool
	ShaderInputAttachmentArrayNonUniformIndexing    bool
	ShaderUniformTexelBufferArrayNonUniformIndexing bool
	ShaderStorageTexelBufferArrayNonUniformIndexing bool

	DescriptorBindingUniformBufferUpdateAfterBind      bool
	DescriptorBindingSampledImageUpdateAfterBind       bool
	DescriptorBindingStorageImageUpdateAfterBind       bool
	DescriptorBindingStorageBufferUpdateAfterBind      bool
	DescriptorBindingUniformTexelBufferUpdateAfterBind bool
	DescriptorBindingStorageTexelBufferUpdateAfterBind bool

	DescriptorBindingUpdateUnusedWhilePending bool
	DescriptorBindingPartiallyBound           bool
	DescriptorBindingVariableDescriptorCount  bool

	RuntimeDescriptorArray                        bool
	SamplerFilterMinmax                           bool
	ScalarBlockLayout                             bool
	ImagelessFramebuffer                          bool
	UniformBufferStandardLayout                   bool
	ShaderSubgroupExtendedTypes                   bool
	SeparateDepthStencilLayouts                   bool
	HostQueryReset                                bool
	TimelineSemaphore                             bool
	BufferDeviceAddress                           bool
	BufferDeviceAddressCaptureReplay              bool
	BufferDeviceAddressMultiDevice                bool
	VulkanMemoryModel                             bool
	VulkanMemoryModelDeviceScope                  bool
	VulkanMemoryModelAvailabilityVisibilityChains bool
	ShaderOutputViewportIndex                     bool
	ShaderOutputLayer                             bool
	SubgroupBroadcastDynamicID                    bool

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
