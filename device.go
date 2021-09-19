package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type DeviceHandle C.VkDevice
type vulkanDevice struct {
	driver Driver
	handle VkDevice
}

func (d *vulkanDevice) Driver() Driver {
	return d.driver
}

func (d *vulkanDevice) Handle() VkDevice {
	return d.handle
}

func (d *vulkanDevice) Destroy() error {
	return d.driver.VkDestroyDevice(d.handle, nil)
}

func (d *vulkanDevice) GetQueue(queueFamilyIndex int, queueIndex int) (Queue, error) {
	var queueHandle VkQueue

	err := d.driver.VkGetDeviceQueue(d.handle, Uint32(queueFamilyIndex), Uint32(queueIndex), &queueHandle)
	if err != nil {
		return nil, err
	}

	return &vulkanQueue{driver: d.driver, handle: queueHandle}, nil
}

func (d *vulkanDevice) WaitForIdle() (VkResult, error) {
	return d.driver.VkDeviceWaitIdle(d.handle)
}

func (d *vulkanDevice) WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*VkFence)(fenceUnsafePtr)

	fenceSlice := ([]VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	waitAllConst := C.VK_FALSE
	if waitForAll {
		waitAllConst = C.VK_TRUE
	}

	return d.driver.VkWaitForFences(d.handle, Uint32(fenceCount), fencePtr, VkBool32(waitAllConst), Uint64(common.TimeoutNanoseconds(timeout)))
}

func (d *vulkanDevice) ResetFences(fences []Fence) (VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*VkFence)(fenceUnsafePtr)
	fenceSlice := ([]VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	return d.driver.VkResetFences(d.handle, Uint32(fenceCount), fencePtr)
}

func (d *vulkanDevice) AllocateMemory(o *DeviceMemoryOptions) (DeviceMemory, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var deviceMemory VkDeviceMemory

	res, err := d.driver.VkAllocateMemory(d.handle, (*VkMemoryAllocateInfo)(createInfo), nil, &deviceMemory)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDeviceMemory{
		driver: d.driver,
		device: d.handle,
		handle: deviceMemory,
	}, res, nil
}

func (d *vulkanDevice) UpdateDescriptorSets(writes []WriteDescriptorSetOptions, copies []CopyDescriptorSetOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	writeCount := len(writes)
	copyCount := len(copies)

	var writePtr unsafe.Pointer
	var copyPtr unsafe.Pointer

	if writeCount > 0 {
		writePtr = arena.Malloc(writeCount * C.sizeof_struct_VkWriteDescriptorSet)
		writeSlice := ([]C.VkWriteDescriptorSet)(unsafe.Slice((*C.VkWriteDescriptorSet)(writePtr), writeCount))
		for i := 0; i < writeCount; i++ {
			next, err := common.AllocNext(arena, writes[i])
			if err != nil {
				return err
			}

			writes[i].populate(arena, &(writeSlice[i]), next)
		}
	}

	if copyCount > 0 {
		copyPtr = arena.Malloc(copyCount * C.sizeof_struct_VkCopyDescriptorSet)
		copySlice := ([]C.VkCopyDescriptorSet)(unsafe.Slice((*C.VkCopyDescriptorSet)(copyPtr), copyCount))
		for i := 0; i < copyCount; i++ {
			next, err := common.AllocNext(arena, copies[i])
			if err != nil {
				return err
			}

			copies[i].populate(arena, &(copySlice[i]), next)
		}
	}

	return d.driver.VkUpdateDescriptorSets(d.handle, Uint32(writeCount), (*VkWriteDescriptorSet)(writePtr), Uint32(copyCount), (*VkCopyDescriptorSet)(copyPtr))
}