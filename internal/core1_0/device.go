package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type VulkanDevice struct {
	DeviceDriver driver.Driver
	DeviceHandle driver.VkDevice

	MaximumAPIVersion common.APIVersion
}

func (d *VulkanDevice) Driver() driver.Driver {
	return d.DeviceDriver
}

func (d *VulkanDevice) Handle() driver.VkDevice {
	return d.DeviceHandle
}

func (d *VulkanDevice) APIVersion() common.APIVersion {
	return d.MaximumAPIVersion
}

func (d *VulkanDevice) Destroy(callbacks *driver.AllocationCallbacks) {
	d.DeviceDriver.VkDestroyDevice(d.DeviceHandle, callbacks.Handle())
}

func (d *VulkanDevice) WaitForIdle() (common.VkResult, error) {
	return d.DeviceDriver.VkDeviceWaitIdle(d.DeviceHandle)
}

func (d *VulkanDevice) WaitForFences(waitForAll bool, timeout time.Duration, fences []core1_0.Fence) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	waitAllConst := C.VK_FALSE
	if waitForAll {
		waitAllConst = C.VK_TRUE
	}

	return d.DeviceDriver.VkWaitForFences(d.DeviceHandle, driver.Uint32(fenceCount), fencePtr, driver.VkBool32(waitAllConst), driver.Uint64(common.TimeoutNanoseconds(timeout)))
}

func (d *VulkanDevice) ResetFences(fences []core1_0.Fence) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	return d.DeviceDriver.VkResetFences(d.DeviceHandle, driver.Uint32(fenceCount), fencePtr)
}

func (d *VulkanDevice) UpdateDescriptorSets(writes []core1_0.WriteDescriptorSetOptions, copies []core1_0.CopyDescriptorSetOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	writeCount := len(writes)
	copyCount := len(copies)

	var err error
	var writePtr *C.VkWriteDescriptorSet
	var copyPtr *C.VkCopyDescriptorSet

	if writeCount > 0 {
		writePtr, err = core.AllocOptionSlice[C.VkWriteDescriptorSet, core1_0.WriteDescriptorSetOptions](arena, writes)
		if err != nil {
			return err
		}
	}

	if copyCount > 0 {
		copyPtr, err = core.AllocOptionSlice[C.VkCopyDescriptorSet, core1_0.CopyDescriptorSetOptions](arena, copies)
		if err != nil {
			return err
		}
	}

	d.DeviceDriver.VkUpdateDescriptorSets(d.DeviceHandle, driver.Uint32(writeCount), (*driver.VkWriteDescriptorSet)(unsafe.Pointer(writePtr)), driver.Uint32(copyCount), (*driver.VkCopyDescriptorSet)(unsafe.Pointer(copyPtr)))
	return nil
}

func (d *VulkanDevice) FlushMappedMemoryRanges(ranges []core1_0.MappedMemoryRange) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos, err := core.AllocOptionSlice[C.VkMappedMemoryRange, core1_0.MappedMemoryRange](arena, ranges)
	if err != nil {
		return common.VKErrorUnknown, err
	}

	return d.DeviceDriver.VkFlushMappedMemoryRanges(d.DeviceHandle, driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (d *VulkanDevice) InvalidateMappedMemoryRanges(ranges []core1_0.MappedMemoryRange) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos, err := core.AllocOptionSlice[C.VkMappedMemoryRange, core1_0.MappedMemoryRange](arena, ranges)
	if err != nil {
		return common.VKErrorUnknown, err
	}

	return d.DeviceDriver.VkInvalidateMappedMemoryRanges(d.DeviceHandle, driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (d *VulkanDevice) GetQueue(queueFamilyIndex int, queueIndex int) core1_0.Queue {

	var queueHandle driver.VkQueue

	d.DeviceDriver.VkGetDeviceQueue(d.DeviceHandle, driver.Uint32(queueFamilyIndex), driver.Uint32(queueIndex), &queueHandle)

	return &VulkanQueue{
		DeviceDriver: d.DeviceDriver,
		QueueHandle:  queueHandle,

		MaximumAPIVersion: d.MaximumAPIVersion,
	}
}

func (d *VulkanDevice) AllocateMemory(allocationCallbacks *driver.AllocationCallbacks, o *core1_0.DeviceMemoryOptions) (core1_0.DeviceMemory, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var deviceMemoryHandle driver.VkDeviceMemory

	deviceDriver := d.DeviceDriver
	deviceHandle := d.DeviceHandle

	res, err := deviceDriver.VkAllocateMemory(deviceHandle, (*driver.VkMemoryAllocateInfo)(createInfo), allocationCallbacks.Handle(), &deviceMemoryHandle)
	if err != nil {
		return nil, res, err
	}

	return &VulkanDeviceMemory{
		DeviceDriver:       deviceDriver,
		Device:             deviceHandle,
		DeviceMemoryHandle: deviceMemoryHandle,

		MaximumAPIVersion: d.MaximumAPIVersion,
		size:              o.AllocationSize,
	}, res, nil
}

func (d *VulkanDevice) FreeMemory(deviceMemory core1_0.DeviceMemory, allocationCallbacks *driver.AllocationCallbacks) {
	// This is really only here for a kind of API symmetry
	freeDeviceMemory(deviceMemory, allocationCallbacks)
}
