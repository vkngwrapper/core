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

const (
	// DeviceQueueCreateProtected specifies that the Device Queue is a protected-capable Queue
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlagBits.html
	DeviceQueueCreateProtected core1_0.DeviceQueueCreateFlags = C.VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT

	// QueueProtected specifies capagbilities of Queue objects in a Queue family
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFlagBits.html
	QueueProtected core1_0.QueueFlags = C.VK_QUEUE_PROTECTED_BIT
)

func init() {
	DeviceQueueCreateProtected.Register("Protected")

	QueueProtected.Register("Protected")
}

////

// DeviceQueueInfo2 specifies the parameters used for Device Queue creation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueInfo2.html
type DeviceQueueInfo2 struct {
	// Flags indicates the flags used to create the Device Queue
	Flags core1_0.DeviceQueueCreateFlags
	// QueueFamilyIndex is the index of the queue family to which the Queue belongs
	QueueFamilyIndex int
	// QueueIndex is the index of the Queue to retrieve from within the set of Queue objects
	// that share both the Queue family and flags specified
	QueueIndex int

	common.NextOptions
}

func (o DeviceQueueInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

////

// DeviceGroupSubmitInfo indicates which PhysicalDevice objects execute Semaphore operations
// and CommandBuffer objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupSubmitInfo.html
type DeviceGroupSubmitInfo struct {
	// WaitSemaphoreDeviceIndices is a slice of Device indices indicating which PhysicalDevice
	// executes the Semaphore wait operation in the corresponding element of SubmitInfo.WaitSemaphores
	WaitSemaphoreDeviceIndices []int
	// CommandBufferDeviceMasks is a slice of Device masks indicating which PhysicalDevice objects
	// execute the CommandBuffer in teh corresponding element of SubmitInfo.CommandBuffers
	CommandBufferDeviceMasks []uint32
	// SignalSemaphoreDeviceIndices is a slice of Device indices indicating which PhysicalDevice
	// executes the Semaphore signal operation in the SubmitInfo.SignalSemaphores
	SignalSemaphoreDeviceIndices []int

	common.NextOptions
}

func (o DeviceGroupSubmitInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

////

// ProtectedSubmitInfo indicates whether the submission is protected
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkProtectedSubmitInfo.html
type ProtectedSubmitInfo struct {
	// ProtectedSubmit specifies whether the batch is protected
	ProtectedSubmit bool

	common.NextOptions
}

func (o ProtectedSubmitInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
