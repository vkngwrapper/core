package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

const (
	// FormatFeatureTransferDst specifies that an Image can be used as a destination Image for copy
	// commands and clear commands
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureTransferDst core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_TRANSFER_DST_BIT
	// FormatFeatureTransferSrc specifies that an Image can be used as a source Image for copy commands
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureTransferSrc core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_TRANSFER_SRC_BIT
)

func init() {
	FormatFeatureTransferDst.Register("Transfer Destination")
	FormatFeatureTransferSrc.Register("Transfer Source")
}

// PhysicalDeviceFeatures2 describes the fine-grained features that can be supported
// by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html
type PhysicalDeviceFeatures2 struct {
	// Features describes the fine-grained features of the Vulkan 1.0 API
	Features core1_0.PhysicalDeviceFeatures

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceFeatures2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFeatures2{})))
	}

	data := (*C.VkPhysicalDeviceFeatures2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceFeatures2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceFeatures2)(cDataPointer)

	(&o.Features).PopulateFromCPointer(unsafe.Pointer(&data.features))

	return data.pNext, nil
}

func (o PhysicalDeviceFeatures2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

// PhysicalDevice16BitStorageFeatures describes features supported by khr_16bit_storage
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice16BitStorageFeatures.html
type PhysicalDevice16BitStorageFeatures struct {
	// StorageBuffer16BitAccess specifies whether objects in the StorageBuffer, ShaderRecordBufferKHR,
	// or PhysicalStorageBuffer storage class with the Block decoration can have 16-bit integer and
	// 16-bit floating-point members
	StorageBuffer16BitAccess bool
	// UniformAndStorageBuffer16BitAccess specifies whether objects in the Uniform storage class with
	// the Block decoration can have 16-bit integer and 16-bit floating-point members
	UniformAndStorageBuffer16BitAccess bool
	// StoragePushConstant16 specifies whether objects in the PushConstant storage class can have
	// 16-bit integer and 16-bit floating-point members
	StoragePushConstant16 bool
	// StorageInputOutput16 specifies whether objects in the Input and Output storage classes can
	// have 16-bit integer and 16-bit floating-point members
	StorageInputOutput16 bool

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

// PhysicalDeviceMultiviewFeatures describes multiview features that can be supported by
// an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewFeatures.html
type PhysicalDeviceMultiviewFeatures struct {
	// Multiview specifies whether the implementation supports multiview rendering within a
	// RenderPass. If this feature is not enabled, the view mask of each subpass must always
	// be zero
	Multiview bool
	// MultiviewGeometryShader specifies whether the implementation supports multiview rendering
	// within a RenderPass, with geometry shaders
	MultiviewGeometryShader bool
	// MultiviewTessellationShader specifies whether the implementation supports multiview rendering
	// within a RenderPass, with tessellation shaders
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

// PhysicalDeviceProtectedMemoryFeatures describes protected memory features that can be supported
// by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProtectedMemoryFeatures.html
type PhysicalDeviceProtectedMemoryFeatures struct {
	// ProtectedMemory specifies whether protected memory is supported
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

// PhysicalDeviceSamplerYcbcrConversionFeatures describes Y'CbCr conversion features that can
// be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSamplerYcbcrConversionFeatures.html
type PhysicalDeviceSamplerYcbcrConversionFeatures struct {
	// SamplerYcbcrConversion specifies whether the implementation support sampler Y'CbCr conversion.
	SamplerYcbcrConversion bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceSamplerYcbcrConversionFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerYcbcrConversionFeatures{})))
	}

	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSamplerYcbcrConversionFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(cDataPointer)
	o.SamplerYcbcrConversion = info.samplerYcbcrConversion != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceSamplerYcbcrConversionFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

// PhysicalDeviceShaderDrawParametersFeatures describes shader draw parameter features that
// can be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderDrawParametersFeatures.html
type PhysicalDeviceShaderDrawParametersFeatures struct {
	// ShaderDrawParameters specifies whether the implementation supports the SPIR-V DrawParameters
	// capability
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

// PhysicalDeviceVariablePointersFeatures describes variable pointer features that can be
// supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVariablePointersFeatures.html
type PhysicalDeviceVariablePointersFeatures struct {
	// VariablePointersStorageBuffer specifies whether the implementation supports the SPIR-V
	// VariablePointersStorageBuffer capability
	VariablePointersStorageBuffer bool
	// VariablePointers specifies whether the implementation supports the SPIR-V VariablePointers
	// capability
	VariablePointers bool

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
