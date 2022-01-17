package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type vulkanDevice struct {
	driver driver.Driver
	handle driver.VkDevice
}

func (d *vulkanDevice) Driver() driver.Driver {
	return d.driver
}

func (d *vulkanDevice) Handle() driver.VkDevice {
	return d.handle
}

func (d *vulkanDevice) Destroy(callbacks *AllocationCallbacks) {
	d.driver.VkDestroyDevice(d.handle, callbacks.Handle())
}

func (d *vulkanDevice) GetQueue(queueFamilyIndex int, queueIndex int) Queue {
	var queueHandle driver.VkQueue

	d.driver.VkGetDeviceQueue(d.handle, driver.Uint32(queueFamilyIndex), driver.Uint32(queueIndex), &queueHandle)

	return &vulkanQueue{driver: d.driver, handle: queueHandle}
}

func (d *vulkanDevice) WaitForIdle() (common.VkResult, error) {
	return d.driver.VkDeviceWaitIdle(d.handle)
}

func (d *vulkanDevice) WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (common.VkResult, error) {
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

	return d.driver.VkWaitForFences(d.handle, driver.Uint32(fenceCount), fencePtr, driver.VkBool32(waitAllConst), driver.Uint64(common.TimeoutNanoseconds(timeout)))
}

func (d *vulkanDevice) ResetFences(fences []Fence) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	return d.driver.VkResetFences(d.handle, driver.Uint32(fenceCount), fencePtr)
}

func (d *vulkanDevice) AllocateMemory(allocationCallbacks *AllocationCallbacks, o *DeviceMemoryOptions) (DeviceMemory, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var deviceMemory driver.VkDeviceMemory

	res, err := d.driver.VkAllocateMemory(d.handle, (*driver.VkMemoryAllocateInfo)(createInfo), nil, &deviceMemory)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDeviceMemory{
		driver: d.driver,
		device: d.handle,
		handle: deviceMemory,
		size:   o.AllocationSize,
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

			_, err = writes[i].populate(arena, &(writeSlice[i]), next)
			if err != nil {
				return err
			}
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

			_, err = copies[i].populate(arena, &(copySlice[i]), next)
			if err != nil {
				return err
			}
		}
	}

	d.driver.VkUpdateDescriptorSets(d.handle, driver.Uint32(writeCount), (*driver.VkWriteDescriptorSet)(writePtr), driver.Uint32(copyCount), (*driver.VkCopyDescriptorSet)(copyPtr))
	return nil
}

type MappedMemoryRange struct {
	Memory DeviceMemory
	Offset int
	Size   int

	common.HaveNext
}

func (r *MappedMemoryRange) populate(mappedRange *C.VkMappedMemoryRange, next unsafe.Pointer) error {
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = next
	mappedRange.memory = C.VkDeviceMemory(unsafe.Pointer(r.Memory.Handle()))
	mappedRange.offset = C.VkDeviceSize(r.Offset)
	mappedRange.size = C.VkDeviceSize(r.Size)

	return nil
}

func (r *MappedMemoryRange) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkMappedMemoryRange)(allocator.Malloc(C.sizeof_struct_VkMappedMemoryRange))
	err := r.populate(createInfo, next)
	return unsafe.Pointer(createInfo), err
}

func (d *vulkanDevice) FlushMappedMemoryRanges(ranges []*MappedMemoryRange) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos := (*C.VkMappedMemoryRange)(arena.Malloc(rangeCount * C.sizeof_struct_VkMappedMemoryRange))
	createInfoSlice := ([]C.VkMappedMemoryRange)(unsafe.Slice(createInfos, rangeCount))

	for rangeIndex, memRange := range ranges {
		next, err := common.AllocNext(arena, memRange)
		if err != nil {
			return common.VKErrorUnknown, err
		}

		err = memRange.populate(&createInfoSlice[rangeIndex], next)
		if err != nil {
			return common.VKErrorUnknown, err
		}
	}

	return d.driver.VkFlushMappedMemoryRanges(d.handle, driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (d *vulkanDevice) InvalidateMappedMemoryRanges(ranges []*MappedMemoryRange) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos := (*C.VkMappedMemoryRange)(arena.Malloc(rangeCount * C.sizeof_struct_VkMappedMemoryRange))
	createInfoSlice := ([]C.VkMappedMemoryRange)(unsafe.Slice(createInfos, rangeCount))

	for rangeIndex, memRange := range ranges {
		next, err := common.AllocNext(arena, memRange)
		if err != nil {
			return common.VKErrorUnknown, err
		}

		err = memRange.populate(&createInfoSlice[rangeIndex], next)
		if err != nil {
			return common.VKErrorUnknown, err
		}
	}

	return d.driver.VkInvalidateMappedMemoryRanges(d.handle, driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}
