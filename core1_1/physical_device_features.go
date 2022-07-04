package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DeviceFeatures struct {
	Features core1_0.PhysicalDeviceFeatures

	common.NextOptions
	common.NextOutData
}

func (o *DeviceFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFeatures2{})))
	}

	data := (*C.VkPhysicalDeviceFeatures2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *DeviceFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceFeatures2)(cDataPointer)

	(&o.Features).PopulateFromCPointer(unsafe.Pointer(&data.features))

	return data.pNext, nil
}

func (o DeviceFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFeatures2{})))
	}

	data := (*C.VkPhysicalDeviceFeatures2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	data.pNext = next
	_, err := o.Features.PopulateCPointer(allocator, unsafe.Pointer(&data.features))

	return preallocatedPointer, err
}

////

type PhysicalDevice16BitStorageFeatures struct {
	StorageBuffer16BitAccess           bool
	UniformAndStorageBuffer16BitAccess bool
	StoragePushConstant16              bool
	StorageInputOutput16               bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDevice16BitStorageFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice16BitStorageFeatures{})))
	}

	data := (*C.VkPhysicalDevice16BitStorageFeatures)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevice16BitStorageFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDevice16BitStorageFeatures)(cDataPointer)

	o.StorageBuffer16BitAccess = data.storageBuffer16BitAccess != C.VkBool32(0)
	o.UniformAndStorageBuffer16BitAccess = data.uniformAndStorageBuffer16BitAccess != C.VkBool32(0)
	o.StoragePushConstant16 = data.storagePushConstant16 != C.VkBool32(0)
	o.StorageInputOutput16 = data.storageInputOutput16 != C.VkBool32(0)

	return data.pNext, nil
}

func (o PhysicalDevice16BitStorageFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice16BitStorageFeatures{})))
	}

	data := (*C.VkPhysicalDevice16BitStorageFeatures)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES
	data.pNext = next

	data.storageBuffer16BitAccess = C.VkBool32(0)
	data.uniformAndStorageBuffer16BitAccess = C.VkBool32(0)
	data.storagePushConstant16 = C.VkBool32(0)
	data.storageInputOutput16 = C.VkBool32(0)

	if o.StorageBuffer16BitAccess {
		data.storageBuffer16BitAccess = C.VkBool32(1)
	}

	if o.UniformAndStorageBuffer16BitAccess {
		data.uniformAndStorageBuffer16BitAccess = C.VkBool32(1)
	}

	if o.StoragePushConstant16 {
		data.storagePushConstant16 = C.VkBool32(1)
	}

	if o.StorageInputOutput16 {
		data.storageInputOutput16 = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

type PhysicalDeviceMultiviewFeatures struct {
	Multiview                   bool
	MultiviewGeometryShader     bool
	MultiviewTessellationShader bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceMultiviewFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMultiviewFeatures{})))
	}
	info := (*C.VkPhysicalDeviceMultiviewFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMultiviewFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceMultiviewFeatures)(cDataPointer)

	o.Multiview = info.multiview != C.VkBool32(0)
	o.MultiviewGeometryShader = info.multiviewGeometryShader != C.VkBool32(0)
	o.MultiviewTessellationShader = info.multiviewTessellationShader != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceMultiviewFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMultiviewFeatures{})))
	}
	info := (*C.VkPhysicalDeviceMultiviewFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
	info.pNext = next
	info.multiview = C.VkBool32(0)
	info.multiviewGeometryShader = C.VkBool32(0)
	info.multiviewTessellationShader = C.VkBool32(0)

	if o.Multiview {
		info.multiview = C.VkBool32(1)
	}

	if o.MultiviewGeometryShader {
		info.multiviewGeometryShader = C.VkBool32(1)
	}

	if o.MultiviewTessellationShader {
		info.multiviewTessellationShader = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

type PhysicalDeviceProtectedMemoryFeatures struct {
	ProtectedMemory bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceProtectedMemoryFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceProtectedMemoryFeatures)
	}

	info := (*C.VkPhysicalDeviceProtectedMemoryFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceProtectedMemoryFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceProtectedMemoryFeatures)(cDataPointer)

	o.ProtectedMemory = info.protectedMemory != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceProtectedMemoryFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceProtectedMemoryFeatures)
	}

	info := (*C.VkPhysicalDeviceProtectedMemoryFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES
	info.pNext = next
	info.protectedMemory = C.VkBool32(0)

	if o.ProtectedMemory {
		info.protectedMemory = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

type PhysicalDeviceSamplerYcbcrFeatures struct {
	SamplerYcbcrConversion bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceSamplerYcbcrFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerYcbcrConversionFeatures{})))
	}

	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSamplerYcbcrFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(cDataPointer)
	o.SamplerYcbcrConversion = info.samplerYcbcrConversion != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceSamplerYcbcrFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerYcbcrConversionFeatures{})))
	}

	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
	info.pNext = next
	info.samplerYcbcrConversion = C.VkBool32(0)

	if o.SamplerYcbcrConversion {
		info.samplerYcbcrConversion = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

type PhysicalDeviceShaderDrawParametersFeatures struct {
	ShaderDrawParameters bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceShaderDrawParametersFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceShaderDrawParametersFeatures)
	}

	info := (*C.VkPhysicalDeviceShaderDrawParametersFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderDrawParametersFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderDrawParametersFeatures)(cDataPointer)

	o.ShaderDrawParameters = info.shaderDrawParameters != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceShaderDrawParametersFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceShaderDrawParametersFeatures)
	}

	info := (*C.VkPhysicalDeviceShaderDrawParametersFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES
	info.pNext = next
	info.shaderDrawParameters = C.VkBool32(0)

	if o.ShaderDrawParameters {
		info.shaderDrawParameters = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

type PhysicalDeviceVariablePointersFeatures struct {
	VariablePointersStorageBuffer bool
	VariablePointers              bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceVariablePointersFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVariablePointersFeatures{})))
	}

	createInfo := (*C.VkPhysicalDeviceVariablePointersFeatures)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES
	createInfo.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceVariablePointersFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPhysicalDeviceVariablePointersFeatures)(cDataPointer)
	o.VariablePointers = false
	o.VariablePointersStorageBuffer = false

	if createInfo.variablePointersStorageBuffer != C.VkBool32(0) {
		o.VariablePointersStorageBuffer = true
	}
	if createInfo.variablePointers != C.VkBool32(0) {
		o.VariablePointers = true
	}

	return createInfo.pNext, nil
}

func (o PhysicalDeviceVariablePointersFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVariablePointersFeatures{})))
	}

	createInfo := (*C.VkPhysicalDeviceVariablePointersFeatures)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES
	createInfo.pNext = next
	createInfo.variablePointersStorageBuffer = C.VkBool32(0)
	createInfo.variablePointers = C.VkBool32(0)

	if o.VariablePointersStorageBuffer {
		createInfo.variablePointersStorageBuffer = C.VkBool32(1)
	}
	if o.VariablePointers {
		createInfo.variablePointers = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
