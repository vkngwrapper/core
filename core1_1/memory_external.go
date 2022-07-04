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

type ExternalMemoryFeatures int32

var externalMemoryFeaturesMapping = common.NewFlagStringMapping[ExternalMemoryFeatures]()

func (f ExternalMemoryFeatures) Register(str string) {
	externalMemoryFeaturesMapping.Register(f, str)
}

func (f ExternalMemoryFeatures) String() string {
	return externalMemoryFeaturesMapping.FlagsToString(f)
}

////

type ExternalMemoryHandleTypes int32

var externalMemoryHandleTypesMapping = common.NewFlagStringMapping[ExternalMemoryHandleTypes]()

func (f ExternalMemoryHandleTypes) Register(str string) {
	externalMemoryHandleTypesMapping.Register(f, str)
}

func (f ExternalMemoryHandleTypes) String() string {
	return externalMemoryHandleTypesMapping.FlagsToString(f)
}

////

const (
	ExternalMemoryFeatureDedicatedOnly ExternalMemoryFeatures = C.VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT
	ExternalMemoryFeatureExportable    ExternalMemoryFeatures = C.VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT
	ExternalMemoryFeatureImportable    ExternalMemoryFeatures = C.VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT

	ExternalMemoryHandleTypeD3D11Texture    ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT
	ExternalMemoryHandleTypeD3D11TextureKMT ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT
	ExternalMemoryHandleTypeD3D12Heap       ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT
	ExternalMemoryHandleTypeD3D12Resource   ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT
	ExternalMemoryHandleTypeOpaqueFD        ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT
	ExternalMemoryHandleTypeOpaqueWin32     ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT
	ExternalMemoryHandleTypeOpaqueWin32KMT  ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
)

func init() {
	ExternalMemoryFeatureDedicatedOnly.Register("Dedicated Only")
	ExternalMemoryFeatureExportable.Register("Exportable")
	ExternalMemoryFeatureImportable.Register("Importable")

	ExternalMemoryHandleTypeD3D11Texture.Register("D3D11 Texture")
	ExternalMemoryHandleTypeD3D11TextureKMT.Register("D3D11 Texture (Kernel-Mode)")
	ExternalMemoryHandleTypeD3D12Heap.Register("D3D12 Heap")
	ExternalMemoryHandleTypeD3D12Resource.Register("D3D12 Resource")
	ExternalMemoryHandleTypeOpaqueFD.Register("Opaque File Descriptor")
	ExternalMemoryHandleTypeOpaqueWin32.Register("Opaque Win32")
	ExternalMemoryHandleTypeOpaqueWin32KMT.Register("Opaque Win32 (Kernel-Mode)")
}

////

type ExternalMemoryProperties struct {
	ExternalMemoryFeatures        ExternalMemoryFeatures
	ExportFromImportedHandleTypes ExternalMemoryHandleTypes
	CompatibleHandleTypes         ExternalMemoryHandleTypes
}

func (o ExternalMemoryProperties) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalMemoryProperties{})))
	}

	info := (*C.VkExternalMemoryProperties)(preallocatedPointer)
	info.externalMemoryFeatures = C.VkExternalMemoryFeatureFlags(o.ExternalMemoryFeatures)
	info.exportFromImportedHandleTypes = C.VkExternalMemoryHandleTypeFlags(o.ExportFromImportedHandleTypes)
	info.compatibleHandleTypes = C.VkExternalMemoryHandleTypeFlags(o.CompatibleHandleTypes)

	return preallocatedPointer, nil
}

func (o *ExternalMemoryProperties) PopulateOutData(cDataPointer unsafe.Pointer) error {
	info := (*C.VkExternalMemoryProperties)(cDataPointer)
	o.ExternalMemoryFeatures = ExternalMemoryFeatures(info.externalMemoryFeatures)
	o.ExportFromImportedHandleTypes = ExternalMemoryHandleTypes(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalMemoryHandleTypes(info.compatibleHandleTypes)

	return nil
}

////

type ExternalBufferOptions struct {
	Flags      core1_0.BufferCreateFlags
	Usage      core1_0.BufferUsages
	HandleType ExternalMemoryHandleTypes

	common.NextOptions
}

func (o ExternalBufferOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalBufferInfo{})))
	}

	info := (*C.VkPhysicalDeviceExternalBufferInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO
	info.pNext = next
	info.flags = (C.VkBufferCreateFlags)(o.Flags)
	info.usage = (C.VkBufferUsageFlags)(o.Usage)
	info.handleType = (C.VkExternalMemoryHandleTypeFlagBits)(o.HandleType)

	return preallocatedPointer, nil
}

////

type ExternalBufferOutData struct {
	ExternalMemoryProperties ExternalMemoryProperties

	common.NextOutData
}

func (o *ExternalBufferOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalBufferProperties{})))
	}
	info := (*C.VkExternalBufferProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalBufferOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalBufferProperties)(cDataPointer)

	err = (&o.ExternalMemoryProperties).PopulateOutData(unsafe.Pointer(&info.externalMemoryProperties))
	return info.pNext, nil
}

////

type ExternalMemoryBufferOptions struct {
	HandleTypes ExternalMemoryHandleTypes

	common.NextOptions
}

func (o ExternalMemoryBufferOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalMemoryBufferCreateInfo{})))
	}

	info := (*C.VkExternalMemoryBufferCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO
	info.pNext = next
	info.handleTypes = C.VkExternalMemoryHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}

////

type ExternalMemoryImageOptions struct {
	HandleTypes ExternalMemoryHandleTypes

	common.NextOptions
}

func (o ExternalMemoryImageOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalMemoryImageCreateInfo{})))
	}

	info := (*C.VkExternalMemoryImageCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO
	info.pNext = next
	info.handleTypes = C.VkExternalMemoryHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}

////

type PhysicalDeviceExternalImageFormatOptions struct {
	HandleType ExternalMemoryHandleTypes

	common.NextOptions
}

func (o PhysicalDeviceExternalImageFormatOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalImageFormatInfo{})))
	}

	info := (*C.VkPhysicalDeviceExternalImageFormatInfo)(preallocatedPointer)

	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO
	info.pNext = next
	info.handleType = C.VkExternalMemoryHandleTypeFlagBits(o.HandleType)

	return preallocatedPointer, nil
}

////

type ExportMemoryAllocateOptions struct {
	HandleTypes ExternalMemoryHandleTypes

	common.NextOptions
}

func (o ExportMemoryAllocateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportMemoryAllocateInfo{})))
	}

	info := (*C.VkExportMemoryAllocateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO
	info.pNext = next
	info.handleTypes = C.VkExternalMemoryHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}

////

type ExternalImageFormatOutData struct {
	ExternalMemoryProperties ExternalMemoryProperties

	common.NextOutData
}

func (o *ExternalImageFormatOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalImageFormatProperties{})))
	}

	info := (*C.VkExternalImageFormatProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalImageFormatOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalImageFormatProperties)(cDataPointer)

	err = (&o.ExternalMemoryProperties).PopulateOutData(unsafe.Pointer(&info.externalMemoryProperties))
	return info.pNext, err
}
