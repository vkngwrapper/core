package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// SemaphoreImportFlags specifies additional parameters of Semaphore payload import
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagBits.html
type SemaphoreImportFlags int32

var semaphoreImportFlagsMapping = common.NewFlagStringMapping[SemaphoreImportFlags]()

func (f SemaphoreImportFlags) Register(str string) {
	semaphoreImportFlagsMapping.Register(f, str)
}

func (f SemaphoreImportFlags) String() string {
	return semaphoreImportFlagsMapping.FlagsToString(f)
}

////

// ExternalSemaphoreFeatureFlags describes features of an external Semaphore handle type
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlagBits.html
type ExternalSemaphoreFeatureFlags int32

var externalSemaphoreFeaturesMapping = common.NewFlagStringMapping[ExternalSemaphoreFeatureFlags]()

func (f ExternalSemaphoreFeatureFlags) Register(str string) {
	externalSemaphoreFeaturesMapping.Register(f, str)
}

func (f ExternalSemaphoreFeatureFlags) String() string {
	return externalSemaphoreFeaturesMapping.FlagsToString(f)
}

////

// ExternalSemaphoreHandleTypeFlags is a bitmask of valid external Semaphore handle types
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
type ExternalSemaphoreHandleTypeFlags int32

var externalSemaphoreHandleTypesMapping = common.NewFlagStringMapping[ExternalSemaphoreHandleTypeFlags]()

func (f ExternalSemaphoreHandleTypeFlags) Register(str string) {
	externalSemaphoreHandleTypesMapping.Register(f, str)
}

func (f ExternalSemaphoreHandleTypeFlags) String() string {
	return externalSemaphoreHandleTypesMapping.FlagsToString(f)
}

////

const (
	// ExternalSemaphoreFeatureExportable specifies that handles of this type can be exported
	// from Vulkan Semaphore objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlagBits.html
	ExternalSemaphoreFeatureExportable ExternalSemaphoreFeatureFlags = C.VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT
	// ExternalSemaphoreFeatureImportable specifies that handles of this type can be imported
	// as Vulkan Semaphore objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlagBits.html
	ExternalSemaphoreFeatureImportable ExternalSemaphoreFeatureFlags = C.VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT

	// ExternalSemaphoreHandleTypeOpaqueFD specifies a POSIX file descriptor handle that has
	// only limited valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
	ExternalSemaphoreHandleTypeOpaqueFD ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT
	// ExternalSemaphoreHandleTypeOpaqueWin32 specifies an NT handle that has only limited
	// valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
	ExternalSemaphoreHandleTypeOpaqueWin32 ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT
	// ExternalSemaphoreHandleTypeOpaqueWin32KMT specifies a global share handle that has only
	// limited valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
	ExternalSemaphoreHandleTypeOpaqueWin32KMT ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	// ExternalSemaphoreHandleTypeD3D12Fence specifies an NT handle returned by
	// ID3D12Device::CreateSharedHandle referring to a Direct3D 12 fence, or
	// ID3D11Device5::CreateFence referring to a Direct3D 11 fence
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
	ExternalSemaphoreHandleTypeD3D12Fence ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT
	// ExternalSemaphoreHandleTypeSyncFD specifies a POSIX file descriptor handle to a Linux Sync
	// File or Android Fence object. It can be used with any native API accepting a valid sync file or
	// Fence as input
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
	ExternalSemaphoreHandleTypeSyncFD ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT

	// SemaphoreImportTemporary specifies that the Semaphore payload will be imported only
	// temporarily, regardless of the permanence of the handle type
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagBits.html
	SemaphoreImportTemporary SemaphoreImportFlags = C.VK_SEMAPHORE_IMPORT_TEMPORARY_BIT
)

func init() {
	ExternalSemaphoreFeatureExportable.Register("Exportable")
	ExternalSemaphoreFeatureImportable.Register("Importable")

	ExternalSemaphoreHandleTypeOpaqueFD.Register("Opaque File Descriptor")
	ExternalSemaphoreHandleTypeOpaqueWin32.Register("Opaque Win32 Handle")
	ExternalSemaphoreHandleTypeOpaqueWin32KMT.Register("Opaque Win32 Handle (Kernel Mode)")
	ExternalSemaphoreHandleTypeD3D12Fence.Register("D3D Fence")
	ExternalSemaphoreHandleTypeSyncFD.Register("Sync File Descriptor")

	SemaphoreImportTemporary.Register("Temporary")
}

////

// PhysicalDeviceExternalSemaphoreInfo specifies Semaphore creation parameters
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html
type PhysicalDeviceExternalSemaphoreInfo struct {
	// HandleType specifies the external Semaphore handle type for which capabilities will
	// be returned
	HandleType ExternalSemaphoreHandleTypeFlags

	common.NextOptions
}

func (o PhysicalDeviceExternalSemaphoreInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalSemaphoreInfo{})))
	}

	info := (*C.VkPhysicalDeviceExternalSemaphoreInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO
	info.pNext = next
	info.handleType = C.VkExternalSemaphoreHandleTypeFlagBits(o.HandleType)

	return preallocatedPointer, nil
}

////

// ExternalSemaphoreProperties describes supported external Semaphore handle features
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html
type ExternalSemaphoreProperties struct {
	// ExportFromImportedHandleTypes specifies which types of imported handle HandleType can
	// be exported from
	ExportFromImportedHandleTypes ExternalSemaphoreHandleTypeFlags
	// CompatibleHandleTypes specifies handle types which can be specified at the same time as
	// HandleType when creating a Semaphore
	CompatibleHandleTypes ExternalSemaphoreHandleTypeFlags
	// ExternalSemaphoreFeatures describes the features of HandleType
	ExternalSemaphoreFeatures ExternalSemaphoreFeatureFlags

	common.NextOutData
}

func (o *ExternalSemaphoreProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalSemaphoreProperties{})))
	}

	info := (*C.VkExternalSemaphoreProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalSemaphoreProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalSemaphoreProperties)(cDataPointer)

	o.ExportFromImportedHandleTypes = ExternalSemaphoreHandleTypeFlags(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalSemaphoreHandleTypeFlags(info.compatibleHandleTypes)
	o.ExternalSemaphoreFeatures = ExternalSemaphoreFeatureFlags(info.externalSemaphoreFeatures)

	return info.pNext, nil
}

////

// ExportSemaphoreCreateInfo specifies handle types that can be exported from a Semaphore
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreCreateInfo.html
type ExportSemaphoreCreateInfo struct {
	// HandleTypes specifies one or more Semaphore handle types the application can export
	// from the resulting Semaphore
	HandleTypes ExternalSemaphoreHandleTypeFlags

	common.NextOptions
}

func (o ExportSemaphoreCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportSemaphoreCreateInfo{})))
	}

	info := (*C.VkExportSemaphoreCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO
	info.pNext = next
	info.handleTypes = C.VkExternalSemaphoreHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}
