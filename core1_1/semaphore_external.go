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

type ExternalSemaphoreFeatures int32

var externalSemaphoreFeaturesMapping = common.NewFlagStringMapping[ExternalSemaphoreFeatures]()

func (f ExternalSemaphoreFeatures) Register(str string) {
	externalSemaphoreFeaturesMapping.Register(f, str)
}

func (f ExternalSemaphoreFeatures) String() string {
	return externalSemaphoreFeaturesMapping.FlagsToString(f)
}

////

type ExternalSemaphoreHandleTypes int32

var externalSemaphoreHandleTypesMapping = common.NewFlagStringMapping[ExternalSemaphoreHandleTypes]()

func (f ExternalSemaphoreHandleTypes) Register(str string) {
	externalSemaphoreHandleTypesMapping.Register(f, str)
}

func (f ExternalSemaphoreHandleTypes) String() string {
	return externalSemaphoreHandleTypesMapping.FlagsToString(f)
}

////

const (
	ExternalSemaphoreFeatureExportable ExternalSemaphoreFeatures = C.VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT
	ExternalSemaphoreFeatureImportable ExternalSemaphoreFeatures = C.VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT

	ExternalSemaphoreHandleTypeOpaqueFD       ExternalSemaphoreHandleTypes = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT
	ExternalSemaphoreHandleTypeOpaqueWin32    ExternalSemaphoreHandleTypes = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT
	ExternalSemaphoreHandleTypeOpaqueWin32KMT ExternalSemaphoreHandleTypes = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	ExternalSemaphoreHandleTypeD3D12Fence     ExternalSemaphoreHandleTypes = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT
	ExternalSemaphoreHandleTypeSyncFD         ExternalSemaphoreHandleTypes = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT

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

type ExternalSemaphoreOptions struct {
	HandleType ExternalSemaphoreHandleTypes

	common.NextOptions
}

func (o ExternalSemaphoreOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

type ExternalSemaphoreOutData struct {
	ExportFromImportedHandleTypes ExternalSemaphoreHandleTypes
	CompatibleHandleTypes         ExternalSemaphoreHandleTypes
	ExternalSemaphoreFeatures     ExternalSemaphoreFeatures

	common.NextOutData
}

func (o *ExternalSemaphoreOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalSemaphoreProperties{})))
	}

	info := (*C.VkExternalSemaphoreProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalSemaphoreOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalSemaphoreProperties)(cDataPointer)

	o.ExportFromImportedHandleTypes = ExternalSemaphoreHandleTypes(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalSemaphoreHandleTypes(info.compatibleHandleTypes)
	o.ExternalSemaphoreFeatures = ExternalSemaphoreFeatures(info.externalSemaphoreFeatures)

	return info.pNext, nil
}

////

type ExportSemaphoreOptions struct {
	HandleTypes ExternalSemaphoreHandleTypes

	common.NextOptions
}

func (o ExportSemaphoreOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportSemaphoreCreateInfo{})))
	}

	info := (*C.VkExportSemaphoreCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO
	info.pNext = next
	info.handleTypes = C.VkExternalSemaphoreHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}
