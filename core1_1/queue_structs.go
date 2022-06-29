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

const (
	DeviceQueueCreateProtected core1_0.DeviceQueueCreateFlags = C.VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT

	QueueProtected core1_0.QueueFlags = C.VK_QUEUE_PROTECTED_BIT
)

func init() {
	DeviceQueueCreateProtected.Register("Protected")

	QueueProtected.Register("Protected")
}

////

type DeviceQueueOptions struct {
	Flags            core1_0.DeviceQueueCreateFlags
	QueueFamilyIndex int
	QueueIndex       int

	common.HaveNext
}

func (o DeviceQueueOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDeviceQueueInfo2)
	}

	info := (*C.VkDeviceQueueInfo2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2
	info.pNext = next
	info.flags = C.VkDeviceQueueCreateFlags(o.Flags)
	info.queueFamilyIndex = C.uint32_t(o.QueueFamilyIndex)
	info.queueIndex = C.uint32_t(o.QueueIndex)

	return preallocatedPointer, nil
}

func (o DeviceQueueOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDeviceQueueInfo2)(cDataPointer)
	return info.pNext, nil
}

////

type DeviceGroupSubmitOptions struct {
	WaitSemaphoreDeviceIndices   []int
	CommandBufferDeviceMasks     []uint32
	SignalSemaphoreDeviceIndices []int

	common.HaveNext
}

func (o DeviceGroupSubmitOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupSubmitInfo{})))
	}

	info := (*C.VkDeviceGroupSubmitInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO
	info.pNext = next

	count := len(o.WaitSemaphoreDeviceIndices)
	info.waitSemaphoreCount = C.uint32_t(count)
	info.pWaitSemaphoreDeviceIndices = nil

	if count > 0 {
		indices := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		indexSlice := ([]C.uint32_t)(unsafe.Slice(indices, count))

		for i := 0; i < count; i++ {
			indexSlice[i] = C.uint32_t(o.WaitSemaphoreDeviceIndices[i])
		}
		info.pWaitSemaphoreDeviceIndices = indices
	}

	count = len(o.CommandBufferDeviceMasks)
	info.commandBufferCount = C.uint32_t(count)
	info.pCommandBufferDeviceMasks = nil

	if count > 0 {
		masks := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		maskSlice := ([]C.uint32_t)(unsafe.Slice(masks, count))

		for i := 0; i < count; i++ {
			maskSlice[i] = C.uint32_t(o.CommandBufferDeviceMasks[i])
		}
		info.pCommandBufferDeviceMasks = masks
	}

	count = len(o.SignalSemaphoreDeviceIndices)
	info.signalSemaphoreCount = C.uint32_t(count)
	info.pSignalSemaphoreDeviceIndices = nil

	if count > 0 {
		indices := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		indexSlice := ([]C.uint32_t)(unsafe.Slice(indices, count))

		for i := 0; i < count; i++ {
			indexSlice[i] = C.uint32_t(o.SignalSemaphoreDeviceIndices[i])
		}
		info.pSignalSemaphoreDeviceIndices = indices
	}

	return preallocatedPointer, nil
}

func (o DeviceGroupSubmitOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDeviceGroupSubmitInfo)(cDataPointer)
	return info.pNext, nil
}

////

type ProtectedSubmitOptions struct {
	ProtectedSubmit bool

	common.HaveNext
}

func (o ProtectedSubmitOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkProtectedSubmitInfo)
	}

	info := (*C.VkProtectedSubmitInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO
	info.pNext = next
	info.protectedSubmit = C.VkBool32(0)

	if o.ProtectedSubmit {
		info.protectedSubmit = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o ProtectedSubmitOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkProtectedSubmitInfo)(cDataPointer)
	return info.pNext, nil
}
