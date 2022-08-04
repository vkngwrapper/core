package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// ExternalMemoryFeatureFlags specifies features of an external memory handle type
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagBits.html
type ExternalMemoryFeatureFlags int32

var externalMemoryFeaturesMapping = common.NewFlagStringMapping[ExternalMemoryFeatureFlags]()

func (f ExternalMemoryFeatureFlags) Register(str string) {
	externalMemoryFeaturesMapping.Register(f, str)
}

func (f ExternalMemoryFeatureFlags) String() string {
	return externalMemoryFeaturesMapping.FlagsToString(f)
}

////

// ExternalMemoryHandleTypeFlags specifies external memory handle types
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
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
	// ExternalMemoryFeatureDedicatedOnly specifies that Image or Buffer objects created with the
	// specified parameters and handle type must create or import a dedicated allocation for
	// the Image or Buffer object
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagBits.html
	ExternalMemoryFeatureDedicatedOnly ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT
	// ExternalMemoryFeatureExportable specifies that handles of this type can be exported from
	// Vulkan memory objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagBits.html
	ExternalMemoryFeatureExportable ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT
	// ExternalMemoryFeatureImportable specifies that handles of this type can be imported as Vulkan
	// memory objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagBits.html
	ExternalMemoryFeatureImportable ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT

	// ExternalMemoryHandleTypeD3D11Texture specifies an NT handle returned by
	// IDXGIResource1::CreateSharedHandle referring to a Direct3D 10 or 11 texture resource
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeD3D11Texture ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT
	// ExternalMemoryHandleTypeD3D11TextureKMT specifies a global share handle returned by
	// IDXGIResource::GetSharedHandle referring to a Direct3D 10 or 11 texture resource
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeD3D11TextureKMT ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT
	// ExternalMemoryHandleTypeD3D12Heap specifies an NT handle returned by
	// ID3D12Device::CreateSharedHandle referring to a Direct3D 12 heap resource
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeD3D12Heap ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT
	// ExternalMemoryHandleTypeD3D12Resource specifies an NT handle returned by
	// ID3D12Device::CreateSharedHandle referring to a Direct3D 12 committed resource
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeD3D12Resource ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT
	// ExternalMemoryHandleTypeOpaqueFD specifies a POSIX file descriptor handle that has only limited
	// valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeOpaqueFD ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT
	// ExternalMemoryHandleTypeOpaqueWin32 specifies an NT handle that has only limited valid usage
	// outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeOpaqueWin32 ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT
	// ExternalMemoryHandleTypeOpaqueWin32KMT specifies a global share handle that has only
	// limited valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeOpaqueWin32KMT ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
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

// ExternalMemoryProperties specifies external memory handle type capabilities
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryProperties.html
type ExternalMemoryProperties struct {
	// ExternalMemoryFeatures specifies the features of the handle type
	ExternalMemoryFeatures ExternalMemoryFeatureFlags
	// ExportFromImportedHandleTypes specifies which types of imported handle the handle type can
	// be exported from
	ExportFromImportedHandleTypes ExternalMemoryHandleTypeFlags
	// CompatibleHandleTypes specifies handle types which can be specified at the same time as the
	// handle type which creating an Image compatible with external memory
	CompatibleHandleTypes ExternalMemoryHandleTypeFlags
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

// PhysicalDeviceExternalBufferInfo specifies Buffer creation parameters
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalBufferInfo.html
type PhysicalDeviceExternalBufferInfo struct {
	// Flags describes additional parameters of the Buffer, corresponding to BufferCreateInfo.Flags
	Flags core1_0.BufferCreateFlags
	// Usage describes the intended usage of the Buffer, corresponding to BufferCreateInfo.Usage
	Usage core1_0.BufferUsageFlags
	// HandleType specifies the memory handle type that will be used with the memory
	// associated with the Buffer
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

// ExternalBufferProperties specifies supported external handle capabilities
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html
type ExternalBufferProperties struct {
	// ExternalMemoryProperties specifies various capabilities of the external handle type when
	// used with the specified Buffer creation parameters
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

// ExternalMemoryBufferCreateInfo specifies that a Buffer may be backed by external memory
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html
type ExternalMemoryBufferCreateInfo struct {
	// HandleTypes specifies one or more external memory handle types
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

// ExternalMemoryImageCreateInfo specifies that an Image may be backed by external memory
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html
type ExternalMemoryImageCreateInfo struct {
	// HandleTypes specifies one or more external memory handle types
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

// PhysicalDeviceExternalImageFormatInfo specifies external Image creation parameters
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html
type PhysicalDeviceExternalImageFormatInfo struct {
	// HandleType specifies the memory handle type that will be used with the memory associated
	// with the Image
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

// ExportMemoryAllocateInfo specifies exportable handle types for a DeviceMemory object
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html
type ExportMemoryAllocateInfo struct {
	// HandleTypes specifies one or more memory handle types the application can export from
	// the resulting allocation
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

// ExternalImageFormatProperties specifies supported external handle properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatProperties.html
type ExternalImageFormatProperties struct {
	// ExternalMemoryProperties specifies various capabilities of the external handle type when used
	// with the specified Image creation parameters
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
