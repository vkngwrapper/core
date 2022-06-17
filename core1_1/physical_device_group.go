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
	"github.com/cockroachdb/errors"
	"unsafe"
)

type MemoryAllocateFlags int32

var memoryAllocateFlagsMapping = common.NewFlagStringMapping[MemoryAllocateFlags]()

func (f MemoryAllocateFlags) Register(str string) {
	memoryAllocateFlagsMapping.Register(f, str)
}

func (f MemoryAllocateFlags) String() string {
	return memoryAllocateFlagsMapping.FlagsToString(f)
}

////

const (
	MemoryAllocateDeviceMask MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT
)

func init() {
	MemoryAllocateDeviceMask.Register("Device Mask")
}

////

type DeviceGroupDeviceOptions struct {
	PhysicalDevices []core1_0.PhysicalDevice

	common.HaveNext
}

func (o DeviceGroupDeviceOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupDeviceCreateInfoKHR{})))
	}

	if len(o.PhysicalDevices) < 1 {
		return nil, errors.New("must include at least one physical device in DeviceGroupDeviceOptions")
	}

	createInfo := (*C.VkDeviceGroupDeviceCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR
	createInfo.pNext = next

	count := len(o.PhysicalDevices)
	createInfo.physicalDeviceCount = C.uint32_t(count)
	physicalDevicesPtr := (*C.VkPhysicalDevice)(allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkPhysicalDevice{}))))
	physicalDevicesSlice := ([]C.VkPhysicalDevice)(unsafe.Slice(physicalDevicesPtr, count))

	for i := 0; i < count; i++ {
		physicalDevicesSlice[i] = C.VkPhysicalDevice(unsafe.Pointer(o.PhysicalDevices[i].Handle()))
	}
	createInfo.pPhysicalDevices = physicalDevicesPtr
	return preallocatedPointer, nil
}

func (o DeviceGroupDeviceOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkDeviceGroupDeviceCreateInfoKHR)(cDataPointer)
	return createInfo.pNext, nil
}

////

type MemoryAllocateFlagsOptions struct {
	Flags      MemoryAllocateFlags
	DeviceMask uint32

	common.HaveNext
}

func (o MemoryAllocateFlagsOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryAllocateFlagsInfo{})))
	}

	createInfo := (*C.VkMemoryAllocateFlagsInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkMemoryAllocateFlags(o.Flags)
	createInfo.deviceMask = C.uint32_t(o.DeviceMask)

	return preallocatedPointer, nil
}

func (o MemoryAllocateFlagsOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkMemoryAllocateFlagsInfo)(cDataPointer)
	return createInfo.pNext, nil
}

////

type PhysicalDevice16BitStorageFeaturesOptions struct {
	StorageBuffer16BitAccess           bool
	UniformAndStorageBuffer16BitAccess bool
	StoragePushConstant16              bool
	StorageInputOutput16               bool

	common.HaveNext
}

func (o PhysicalDevice16BitStorageFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o PhysicalDevice16BitStorageFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDevice16BitStorageFeatures)(cDataPointer)

	return data.pNext, nil
}