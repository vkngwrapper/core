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

type DeviceFeaturesOptions struct {
	Features core1_0.PhysicalDeviceFeatures

	common.HaveNext
}

func (o DeviceFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFeatures2{})))
	}

	data := (*C.VkPhysicalDeviceFeatures2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	data.pNext = next
	_, err := o.Features.PopulateCPointer(allocator, unsafe.Pointer(&data.features))

	return preallocatedPointer, err
}

func (o DeviceFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceFeatures2)(cDataPointer)

	return data.pNext, nil
}

////

type PhysicalDevice16BitStorageFeaturesOutData struct {
	StorageBuffer16BitAccess           bool
	UniformAndStorageBuffer16BitAccess bool
	StoragePushConstant16              bool
	StorageInputOutput16               bool

	common.HaveNext
}

func (o *PhysicalDevice16BitStorageFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice16BitStorageFeatures{})))
	}

	data := (*C.VkPhysicalDevice16BitStorageFeatures)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevice16BitStorageFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDevice16BitStorageFeatures)(cDataPointer)

	o.StorageBuffer16BitAccess = data.storageBuffer16BitAccess != C.VkBool32(0)
	o.UniformAndStorageBuffer16BitAccess = data.uniformAndStorageBuffer16BitAccess != C.VkBool32(0)
	o.StoragePushConstant16 = data.storagePushConstant16 != C.VkBool32(0)
	o.StorageInputOutput16 = data.storageInputOutput16 != C.VkBool32(0)

	return data.pNext, nil
}

////

type PhysicalDeviceMultiviewFeaturesOptions struct {
	Multiview                   bool
	MultiviewGeometryShader     bool
	MultiviewTessellationShader bool

	common.HaveNext
}

func (o PhysicalDeviceMultiviewFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o PhysicalDeviceMultiviewFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceMultiviewFeatures)(cDataPointer)
	return info.pNext, nil
}

////

type PhysicalDeviceMultiviewFeaturesOutData struct {
	Multiview                   bool
	MultiviewGeometryShader     bool
	MultiviewTessellationShader bool

	common.HaveNext
}

func (o *PhysicalDeviceMultiviewFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMultiviewFeatures{})))
	}
	info := (*C.VkPhysicalDeviceMultiviewFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMultiviewFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceMultiviewFeatures)(cDataPointer)

	o.Multiview = info.multiview != C.VkBool32(0)
	o.MultiviewGeometryShader = info.multiviewGeometryShader != C.VkBool32(0)
	o.MultiviewTessellationShader = info.multiviewTessellationShader != C.VkBool32(0)

	return info.pNext, nil
}

////

type PhysicalDeviceProtectedMemoryFeaturesOptions struct {
	ProtectedMemory bool

	common.HaveNext
}

func (o PhysicalDeviceProtectedMemoryFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o PhysicalDeviceProtectedMemoryFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceProtectedMemoryFeatures)(cDataPointer)
	return info.pNext, nil
}

////

type PhysicalDeviceProtectedMemoryFeaturesOutData struct {
	ProtectedMemory bool

	common.HaveNext
}

func (o *PhysicalDeviceProtectedMemoryFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceProtectedMemoryFeatures)
	}

	info := (*C.VkPhysicalDeviceProtectedMemoryFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceProtectedMemoryFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceProtectedMemoryFeatures)(cDataPointer)

	o.ProtectedMemory = info.protectedMemory != C.VkBool32(0)

	return info.pNext, nil
}

////

type PhysicalDeviceSamplerYcbcrFeaturesOutData struct {
	SamplerYcbcrConversion bool

	common.HaveNext
}

func (o *PhysicalDeviceSamplerYcbcrFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerYcbcrConversionFeatures{})))
	}

	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSamplerYcbcrFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(cDataPointer)
	o.SamplerYcbcrConversion = info.samplerYcbcrConversion != C.VkBool32(0)

	return info.pNext, nil
}

////

type PhysicalDeviceSamplerYcbcrFeaturesOptions struct {
	SamplerYcbcrConversion bool

	common.HaveNext
}

func (o PhysicalDeviceSamplerYcbcrFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o PhysicalDeviceSamplerYcbcrFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(cDataPointer)
	return info.pNext, nil
}

////

type PhysicalDeviceShaderDrawParametersFeaturesOptions struct {
	ShaderDrawParameters bool

	common.HaveNext
}

func (o PhysicalDeviceShaderDrawParametersFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o PhysicalDeviceShaderDrawParametersFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderDrawParametersFeatures)(cDataPointer)
	return info.pNext, nil
}

////

type PhysicalDeviceShaderDrawParametersFeaturesOutData struct {
	ShaderDrawParameters bool

	common.HaveNext
}

func (o *PhysicalDeviceShaderDrawParametersFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceShaderDrawParametersFeatures)
	}

	info := (*C.VkPhysicalDeviceShaderDrawParametersFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderDrawParametersFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderDrawParametersFeatures)(cDataPointer)

	o.ShaderDrawParameters = info.shaderDrawParameters != C.VkBool32(0)

	return info.pNext, nil
}

////

type PhysicalDeviceVariablePointersFeaturesOptions struct {
	VariablePointersStorageBuffer bool
	VariablePointers              bool

	common.HaveNext
}

func (o PhysicalDeviceVariablePointersFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o PhysicalDeviceVariablePointersFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPhysicalDeviceVariablePointersFeatures)(cDataPointer)
	return createInfo.pNext, nil
}

////

type PhysicalDeviceVariablePointersFeaturesOutData struct {
	VariablePointersStorageBuffer bool
	VariablePointers              bool

	common.HaveNext
}

func (o *PhysicalDeviceVariablePointersFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVariablePointersFeatures{})))
	}

	createInfo := (*C.VkPhysicalDeviceVariablePointersFeatures)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES
	createInfo.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceVariablePointersFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
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
