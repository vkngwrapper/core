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

// FenceImportFlags specifies additional parameters of a Fence payload import
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagBits.html
type FenceImportFlags int32

var fenceImportFlagsMapping = common.NewFlagStringMapping[FenceImportFlags]()

func (f FenceImportFlags) Register(str string) {
	fenceImportFlagsMapping.Register(f, str)
}

func (f FenceImportFlags) String() string {
	return fenceImportFlagsMapping.FlagsToString(f)
}

////

// ExternalFenceFeatureFlags describes features of an external Fence handle type
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceFeatureFlagBits.html
type ExternalFenceFeatureFlags int32

var externalFenceFeaturesMapping = common.NewFlagStringMapping[ExternalFenceFeatureFlags]()

func (f ExternalFenceFeatureFlags) Register(str string) {
	externalFenceFeaturesMapping.Register(f, str)
}

func (f ExternalFenceFeatureFlags) String() string {
	return externalFenceFeaturesMapping.FlagsToString(f)
}

////

// ExternalFenceHandleTypeFlags is a bitmask of valid external Fence handle types
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html
type ExternalFenceHandleTypeFlags int32

var externalFenceHandleTypesMapping = common.NewFlagStringMapping[ExternalFenceHandleTypeFlags]()

func (f ExternalFenceHandleTypeFlags) Register(str string) {
	externalFenceHandleTypesMapping.Register(f, str)
}

func (f ExternalFenceHandleTypeFlags) String() string {
	return externalFenceHandleTypesMapping.FlagsToString(f)
}

////

const (
	// ExternalFenceFeatureExportable specifies handles of this type can be exported from Vulkan
	// Fence objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceFeatureFlagBits.html
	ExternalFenceFeatureExportable ExternalFenceFeatureFlags = C.VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT
	// ExternalFenceFeatureImportable specifies handles of this type can be imported to Vulkan Fence
	// objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceFeatureFlagBits.html
	ExternalFenceFeatureImportable ExternalFenceFeatureFlags = C.VK_EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT

	// ExternalFenceHandleTypeOpaqueFD specifies a POSIX file descriptor handle that has only limited
	// valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html
	ExternalFenceHandleTypeOpaqueFD ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT
	// ExternalFenceHandleTypeOpaqueWin32 specifies an NT handle that has only limited valid usage
	// outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html
	ExternalFenceHandleTypeOpaqueWin32 ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT
	// ExternalFenceHandleTypeOpaqueWin32KMT specifies a global share handle that has only limited
	// usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html
	ExternalFenceHandleTypeOpaqueWin32KMT ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	// ExternalFenceHandleTypeSyncFD specifies a POSIX file descriptor handle to a Linux Sync File
	// or Android Fence
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html
	ExternalFenceHandleTypeSyncFD ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT

	// FenceImportTemporary specifies that the Fence payload will be imported only temporarily,
	// regardless of the permanence of HandleType
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagBits.html
	FenceImportTemporary FenceImportFlags = C.VK_FENCE_IMPORT_TEMPORARY_BIT
)

func init() {
	ExternalFenceFeatureExportable.Register("Exportable")
	ExternalFenceFeatureImportable.Register("Importable")

	ExternalFenceHandleTypeOpaqueFD.Register("Opaque File Descriptor")
	ExternalFenceHandleTypeOpaqueWin32.Register("Opaque Win32")
	ExternalFenceHandleTypeOpaqueWin32KMT.Register("Opaque Win32 Kernel-Mode")
	ExternalFenceHandleTypeSyncFD.Register("Sync File Descriptor")

	FenceImportTemporary.Register("Temporary")
}

// PhysicalDeviceExternalFenceInfo specifies Fence creation parameters
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFenceInfo.html
type PhysicalDeviceExternalFenceInfo struct {
	// HandleType specifies an external Fence handle type for which capabilities will be
	// returned
	HandleType ExternalFenceHandleTypeFlags

	common.NextOptions
}

func (o PhysicalDeviceExternalFenceInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalFenceInfo{})))
	}
	info := (*C.VkPhysicalDeviceExternalFenceInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO
	info.pNext = next
	info.handleType = C.VkExternalFenceHandleTypeFlagBits(o.HandleType)

	return preallocatedPointer, nil
}

////

// ExternalFenceProperties describes supported external Fence handle features
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html
type ExternalFenceProperties struct {
	ExportFromImportedHandleTypes ExternalFenceHandleTypeFlags
	CompatibleHandleTypes         ExternalFenceHandleTypeFlags
	ExternalFenceFeatures         ExternalFenceFeatureFlags

	common.NextOutData
}

func (o *ExternalFenceProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalFenceProperties{})))
	}

	info := (*C.VkExternalFenceProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalFenceProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalFenceProperties)(cDataPointer)

	o.ExportFromImportedHandleTypes = ExternalFenceHandleTypeFlags(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalFenceHandleTypeFlags(info.compatibleHandleTypes)
	o.ExternalFenceFeatures = ExternalFenceFeatureFlags(info.externalFenceFeatures)

	return info.pNext, nil
}

////

// ExportFenceCreateInfo specifies handle types that can be exported from a Fence
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportFenceCreateInfo.html
type ExportFenceCreateInfo struct {
	// HandleTypes specifies one or more Fence handle types the application can export from
	// the resulting Fence
	HandleTypes ExternalFenceHandleTypeFlags

	common.NextOptions
}

func (o ExportFenceCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportFenceCreateInfo{})))
	}

	info := (*C.VkExportFenceCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO
	info.pNext = next
	info.handleTypes = C.VkExternalFenceHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}
