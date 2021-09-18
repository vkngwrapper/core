package core

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

type vulkanCommandPool struct {
	driver Driver
	handle VkCommandPool
	device VkDevice
}

func CreateCommandPool(device Device, o *CommandPoolOptions) (CommandPool, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var cmdPoolHandle VkCommandPool
	res, err := device.Driver().VkCreateCommandPool(device.Handle(), (*VkCommandPoolCreateInfo)(createInfo), nil, &cmdPoolHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanCommandPool{driver: device.Driver(), handle: cmdPoolHandle, device: device.Handle()}, res, nil
}

func (p *vulkanCommandPool) Handle() VkCommandPool {
	return p.handle
}

func (p *vulkanCommandPool) Destroy() error {
	return p.driver.VkDestroyCommandPool(p.device, p.handle, nil)
}

func (p *vulkanCommandPool) DestroyBuffers(buffers []CommandBuffer) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	bufferCount := len(buffers)
	if bufferCount == 0 {
		return nil
	}

	destroyPtr := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{})))
	destroySlice := ([]VkCommandBuffer)(unsafe.Slice((*VkCommandBuffer)(destroyPtr), bufferCount))
	for i := 0; i < bufferCount; i++ {
		destroySlice[i] = buffers[i].Handle()
	}

	return p.driver.VkFreeCommandBuffers(p.device, p.handle, Uint32(bufferCount), (*VkCommandBuffer)(destroyPtr))
}
