package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type SemaphoreImportFlags int32

var semaphoreImportFlagsMapping = common.NewFlagStringMapping[SemaphoreImportFlags]()

func (f SemaphoreImportFlags) Register(str string) {
	semaphoreImportFlagsMapping.Register(f, str)
}

func (f SemaphoreImportFlags) String() string {
	return semaphoreImportFlagsMapping.FlagsToString(f)
}

////

type ExternalSemaphoreFeatureFlags int32

var externalSemaphoreFeaturesMapping = common.NewFlagStringMapping[ExternalSemaphoreFeatureFlags]()

func (f ExternalSemaphoreFeatureFlags) Register(str string) {
	externalSemaphoreFeaturesMapping.Register(f, str)
}

func (f ExternalSemaphoreFeatureFlags) String() string {
	return externalSemaphoreFeaturesMapping.FlagsToString(f)
}

////

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
	ExternalSemaphoreFeatureExportable ExternalSemaphoreFeatureFlags = C.VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT
	ExternalSemaphoreFeatureImportable ExternalSemaphoreFeatureFlags = C.VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT

	ExternalSemaphoreHandleTypeOpaqueFD       ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT
	ExternalSemaphoreHandleTypeOpaqueWin32    ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT
	ExternalSemaphoreHandleTypeOpaqueWin32KMT ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	ExternalSemaphoreHandleTypeD3D12Fence     ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT
	ExternalSemaphoreHandleTypeSyncFD         ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT

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

type PhysicalDeviceExternalSemaphoreInfo struct {
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

type ExternalSemaphoreProperties struct {
	ExportFromImportedHandleTypes ExternalSemaphoreHandleTypeFlags
	CompatibleHandleTypes         ExternalSemaphoreHandleTypeFlags
	ExternalSemaphoreFeatures     ExternalSemaphoreFeatureFlags

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

type ExportSemaphoreCreateInfo struct {
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
