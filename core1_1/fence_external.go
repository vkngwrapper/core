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

type FenceImportFlags int32

var fenceImportFlagsMapping = common.NewFlagStringMapping[FenceImportFlags]()

func (f FenceImportFlags) Register(str string) {
	fenceImportFlagsMapping.Register(f, str)
}

func (f FenceImportFlags) String() string {
	return fenceImportFlagsMapping.FlagsToString(f)
}

////

type ExternalFenceFeatures int32

var externalFenceFeaturesMapping = common.NewFlagStringMapping[ExternalFenceFeatures]()

func (f ExternalFenceFeatures) Register(str string) {
	externalFenceFeaturesMapping.Register(f, str)
}

func (f ExternalFenceFeatures) String() string {
	return externalFenceFeaturesMapping.FlagsToString(f)
}

////

type ExternalFenceHandleTypes int32

var externalFenceHandleTypesMapping = common.NewFlagStringMapping[ExternalFenceHandleTypes]()

func (f ExternalFenceHandleTypes) Register(str string) {
	externalFenceHandleTypesMapping.Register(f, str)
}

func (f ExternalFenceHandleTypes) String() string {
	return externalFenceHandleTypesMapping.FlagsToString(f)
}

////

const (
	ExternalFenceFeatureExportable ExternalFenceFeatures = C.VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT
	ExternalFenceFeatureImportable ExternalFenceFeatures = C.VK_EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT

	ExternalFenceHandleTypeOpaqueFD       ExternalFenceHandleTypes = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT
	ExternalFenceHandleTypeOpaqueWin32    ExternalFenceHandleTypes = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT
	ExternalFenceHandleTypeOpaqueWin32KMT ExternalFenceHandleTypes = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	ExternalFenceHandleTypeSyncFD         ExternalFenceHandleTypes = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT

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

type ExternalFenceOptions struct {
	HandleType ExternalFenceHandleTypes

	common.HaveNext
}

func (o ExternalFenceOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalFenceInfo{})))
	}
	info := (*C.VkPhysicalDeviceExternalFenceInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO
	info.pNext = next
	info.handleType = C.VkExternalFenceHandleTypeFlagBits(o.HandleType)

	return preallocatedPointer, nil
}

func (o ExternalFenceOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceExternalFenceInfo)(cDataPointer)
	return info.pNext, nil
}

////

type ExternalFenceOutData struct {
	ExportFromImportedHandleTypes ExternalFenceHandleTypes
	CompatibleHandleTypes         ExternalFenceHandleTypes
	ExternalFenceFeatures         ExternalFenceFeatures

	common.HaveNext
}

func (o *ExternalFenceOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalFenceProperties{})))
	}

	info := (*C.VkExternalFenceProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalFenceOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalFenceProperties)(cDataPointer)

	o.ExportFromImportedHandleTypes = ExternalFenceHandleTypes(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalFenceHandleTypes(info.compatibleHandleTypes)
	o.ExternalFenceFeatures = ExternalFenceFeatures(info.externalFenceFeatures)

	return info.pNext, nil
}

////

type ExportFenceOptions struct {
	HandleTypes ExternalFenceHandleTypes

	common.HaveNext
}

func (o ExportFenceOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportFenceCreateInfo{})))
	}

	info := (*C.VkExportFenceCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO
	info.pNext = next
	info.handleTypes = C.VkExternalFenceHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}

func (o ExportFenceOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExportFenceCreateInfo)(cDataPointer)
	return info.pNext, nil
}
