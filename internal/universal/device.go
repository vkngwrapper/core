package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0/options"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type VulkanDevice struct {
	driver driver.Driver
	handle driver.VkDevice
}

func (d *VulkanDevice) Driver() driver.Driver {
	return d.driver
}

func (d *VulkanDevice) Handle() driver.VkDevice {
	return d.handle
}

func (d *VulkanDevice) Destroy(callbacks *driver.AllocationCallbacks) {
	d.driver.VkDestroyDevice(d.handle, callbacks.Handle())
}

func (d *VulkanDevice) WaitForIdle() (common.VkResult, error) {
	return d.driver.VkDeviceWaitIdle(d.handle)
}

func (d *VulkanDevice) WaitForFences(waitForAll bool, timeout time.Duration, fences []iface.Fence) (common.VkResult, error) {
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

func (d *VulkanDevice) ResetFences(fences []iface.Fence) (common.VkResult, error) {
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

func (d *VulkanDevice) UpdateDescriptorSets(writes []options.WriteDescriptorSetOptions, copies []options.CopyDescriptorSetOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	writeCount := len(writes)
	copyCount := len(copies)

	var err error
	var writePtr *C.VkWriteDescriptorSet
	var copyPtr *C.VkCopyDescriptorSet

	if writeCount > 0 {
		writePtr, err = core.AllocOptionSlice[C.VkWriteDescriptorSet, options.WriteDescriptorSetOptions](arena, writes)
		if err != nil {
			return err
		}
	}

	if copyCount > 0 {
		copyPtr, err = core.AllocOptionSlice[C.VkCopyDescriptorSet, options.CopyDescriptorSetOptions](arena, copies)
		if err != nil {
			return err
		}
	}

	d.driver.VkUpdateDescriptorSets(d.handle, driver.Uint32(writeCount), (*driver.VkWriteDescriptorSet)(unsafe.Pointer(writePtr)), driver.Uint32(copyCount), (*driver.VkCopyDescriptorSet)(unsafe.Pointer(copyPtr)))
	return nil
}

func (d *VulkanDevice) FlushMappedMemoryRanges(ranges []options.MappedMemoryRange) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos, err := core.AllocOptionSlice[C.VkMappedMemoryRange, options.MappedMemoryRange](arena, ranges)
	if err != nil {
		return common.VKErrorUnknown, err
	}

	return d.driver.VkFlushMappedMemoryRanges(d.handle, driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (d *VulkanDevice) InvalidateMappedMemoryRanges(ranges []options.MappedMemoryRange) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos, err := core.AllocOptionSlice[C.VkMappedMemoryRange, options.MappedMemoryRange](arena, ranges)
	if err != nil {
		return common.VKErrorUnknown, err
	}

	return d.driver.VkInvalidateMappedMemoryRanges(d.handle, driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}
