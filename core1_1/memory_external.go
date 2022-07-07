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

type ExternalMemoryFeatureFlags int32

var externalMemoryFeaturesMapping = common.NewFlagStringMapping[ExternalMemoryFeatureFlags]()

func (f ExternalMemoryFeatureFlags) Register(str string) {
	externalMemoryFeaturesMapping.Register(f, str)
}

func (f ExternalMemoryFeatureFlags) String() string {
	return externalMemoryFeaturesMapping.FlagsToString(f)
}

////

type ExternalMemoryHandleTypeFlags int32

var externalMemoryHandleTypesMapping = common.NewFlagStringMapping[ExternalMemoryHandleTypeFlags]()

func (f ExternalMemoryHandleTypeFlags) Register(str string) {
	externalMemoryHandleTypesMapping.Register(f, str)
}

func (f ExternalMemoryHandleTypeFlags) String() string {
	return externalMemoryHandleTypesMapping.FlagsToString(f)
}

////

const (
	ExternalMemoryFeatureDedicatedOnly ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT
	ExternalMemoryFeatureExportable    ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT
	ExternalMemoryFeatureImportable    ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT

	ExternalMemoryHandleTypeD3D11Texture    ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT
	ExternalMemoryHandleTypeD3D11TextureKMT ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT
	ExternalMemoryHandleTypeD3D12Heap       ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT
	ExternalMemoryHandleTypeD3D12Resource   ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT
	ExternalMemoryHandleTypeOpaqueFD        ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT
	ExternalMemoryHandleTypeOpaqueWin32     ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT
	ExternalMemoryHandleTypeOpaqueWin32KMT  ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
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
	ExternalMemoryFeatures        ExternalMemoryFeatureFlags
	ExportFromImportedHandleTypes ExternalMemoryHandleTypeFlags
	CompatibleHandleTypes         ExternalMemoryHandleTypeFlags
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
	o.ExternalMemoryFeatures = ExternalMemoryFeatureFlags(info.externalMemoryFeatures)
	o.ExportFromImportedHandleTypes = ExternalMemoryHandleTypeFlags(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalMemoryHandleTypeFlags(info.compatibleHandleTypes)

	return nil
}

////

type PhysicalDeviceExternalBufferInfo struct {
	Flags      core1_0.BufferCreateFlags
	Usage      core1_0.BufferUsageFlags
	HandleType ExternalMemoryHandleTypeFlags

	common.NextOptions
}

func (o PhysicalDeviceExternalBufferInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

type ExternalBufferProperties struct {
	ExternalMemoryProperties ExternalMemoryProperties

	common.NextOutData
}

func (o *ExternalBufferProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalBufferProperties{})))
	}
	info := (*C.VkExternalBufferProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalBufferProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalBufferProperties)(cDataPointer)

	err = (&o.ExternalMemoryProperties).PopulateOutData(unsafe.Pointer(&info.externalMemoryProperties))
	return info.pNext, nil
}

////

type ExternalMemoryBufferCreateInfo struct {
	HandleTypes ExternalMemoryHandleTypeFlags

	common.NextOptions
}

func (o ExternalMemoryBufferCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

type ExternalMemoryImageCreateInfo struct {
	HandleTypes ExternalMemoryHandleTypeFlags

	common.NextOptions
}

func (o ExternalMemoryImageCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

type PhysicalDeviceExternalImageFormatInfo struct {
	HandleType ExternalMemoryHandleTypeFlags

	common.NextOptions
}

func (o PhysicalDeviceExternalImageFormatInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

type ExportMemoryAllocateInfo struct {
	HandleTypes ExternalMemoryHandleTypeFlags

	common.NextOptions
}

func (o ExportMemoryAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

type ExternalImageFormatProperties struct {
	ExternalMemoryProperties ExternalMemoryProperties

	common.NextOutData
}

func (o *ExternalImageFormatProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalImageFormatProperties{})))
	}

	info := (*C.VkExternalImageFormatProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalImageFormatProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalImageFormatProperties)(cDataPointer)

	err = (&o.ExternalMemoryProperties).PopulateOutData(unsafe.Pointer(&info.externalMemoryProperties))
	return info.pNext, err
}
