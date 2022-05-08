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

type DeviceFeaturesOutData struct {
	Features core1_0.PhysicalDeviceFeatures

	common.HaveNext
}

func (o *DeviceFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFeatures2{})))
	}

	data := (*C.VkPhysicalDeviceFeatures2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *DeviceFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceFeatures2)(cDataPointer)

	(&o.Features).PopulateFromCPointer(unsafe.Pointer(&data.features))

	return data.pNext, nil
}

////

type DeviceFeatureOptions struct {
	Features core1_0.PhysicalDeviceFeatures

	common.HaveNext
}

func (o DeviceFeatureOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFeatures2{})))
	}

	data := (*C.VkPhysicalDeviceFeatures2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	data.pNext = next
	_, err := o.Features.PopulateCPointer(allocator, unsafe.Pointer(&data.features))

	return preallocatedPointer, err
}

func (o DeviceFeatureOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceFeatures2)(cDataPointer)

	return data.pNext, nil
}

////

type Device16BitStorageFeaturesOutData struct {
	StorageBuffer16BitAccess           bool
	UniformAndStorageBuffer16BitAccess bool
	StoragePushConstant16              bool
	StorageInputOutput16               bool

	common.HaveNext
}

func (o *Device16BitStorageFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice16BitStorageFeatures{})))
	}

	data := (*C.VkPhysicalDevice16BitStorageFeatures)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *Device16BitStorageFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDevice16BitStorageFeatures)(cDataPointer)

	o.StorageBuffer16BitAccess = data.storageBuffer16BitAccess != C.VkBool32(0)
	o.UniformAndStorageBuffer16BitAccess = data.uniformAndStorageBuffer16BitAccess != C.VkBool32(0)
	o.StoragePushConstant16 = data.storagePushConstant16 != C.VkBool32(0)
	o.StorageInputOutput16 = data.storageInputOutput16 != C.VkBool32(0)

	return data.pNext, nil
}

////

type MultiviewFeaturesOptions struct {
	Multiview                   bool
	MultiviewGeometryShader     bool
	MultiviewTessellationShader bool

	common.HaveNext
}

func (o MultiviewFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o MultiviewFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceMultiviewFeatures)(cDataPointer)
	return info.pNext, nil
}

////

type MultiviewFeaturesOutData struct {
	Multiview                   bool
	MultiviewGeometryShader     bool
	MultiviewTessellationShader bool

	common.HaveNext
}

func (o *MultiviewFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMultiviewFeatures{})))
	}
	info := (*C.VkPhysicalDeviceMultiviewFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *MultiviewFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceMultiviewFeatures)(cDataPointer)

	o.Multiview = info.multiview != C.VkBool32(0)
	o.MultiviewGeometryShader = info.multiviewGeometryShader != C.VkBool32(0)
	o.MultiviewTessellationShader = info.multiviewTessellationShader != C.VkBool32(0)

	return info.pNext, nil
}

////

type ProtectedMemoryFeaturesOptions struct {
	ProtectedMemory bool

	common.HaveNext
}

func (o ProtectedMemoryFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o ProtectedMemoryFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceProtectedMemoryFeatures)(cDataPointer)
	return info.pNext, nil
}

////

type ProtectedMemoryFeaturesOutData struct {
	ProtectedMemory bool

	common.HaveNext
}

func (o *ProtectedMemoryFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceProtectedMemoryFeatures)
	}

	info := (*C.VkPhysicalDeviceProtectedMemoryFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ProtectedMemoryFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceProtectedMemoryFeatures)(cDataPointer)

	o.ProtectedMemory = info.protectedMemory != C.VkBool32(0)

	return info.pNext, nil
}

////

type SamplerYcbcrFeaturesOutData struct {
	SamplerYcbcrConversion bool

	common.HaveNext
}

func (o *SamplerYcbcrFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerYcbcrConversionFeatures{})))
	}

	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *SamplerYcbcrFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(cDataPointer)
	o.SamplerYcbcrConversion = info.samplerYcbcrConversion != C.VkBool32(0)

	return info.pNext, nil
}

////

type SamplerYcbcrFeaturesOptions struct {
	SamplerYcbcrConversion bool

	common.HaveNext
}

func (o SamplerYcbcrFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o SamplerYcbcrFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(cDataPointer)
	return info.pNext, nil
}

////

type ShaderDrawParametersFeaturesOptions struct {
	ShaderDrawParameters bool

	common.HaveNext
}

func (o ShaderDrawParametersFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o ShaderDrawParametersFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderDrawParametersFeatures)(cDataPointer)
	return info.pNext, nil
}

////

type ShaderDrawParametersFeaturesOutData struct {
	ShaderDrawParameters bool

	common.HaveNext
}

func (o *ShaderDrawParametersFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceShaderDrawParametersFeatures)
	}

	info := (*C.VkPhysicalDeviceShaderDrawParametersFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ShaderDrawParametersFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderDrawParametersFeatures)(cDataPointer)

	o.ShaderDrawParameters = info.shaderDrawParameters != C.VkBool32(0)

	return info.pNext, nil
}

////

type VariablePointersFeatureOptions struct {
	VariablePointersStorageBuffer bool
	VariablePointers              bool

	common.HaveNext
}

func (o VariablePointersFeatureOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o VariablePointersFeatureOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPhysicalDeviceVariablePointersFeatures)(cDataPointer)
	return createInfo.pNext, nil
}

////

type VariablePointersFeatureOutData struct {
	VariablePointersStorageBuffer bool
	VariablePointers              bool

	common.HaveNext
}

func (o *VariablePointersFeatureOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVariablePointersFeatures{})))
	}

	createInfo := (*C.VkPhysicalDeviceVariablePointersFeatures)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES
	createInfo.pNext = next

	return preallocatedPointer, nil
}

func (o *VariablePointersFeatureOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
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
