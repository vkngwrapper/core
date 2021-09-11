package commands

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resources"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type vulkanCommandPool struct {
	loader loader.Loader
	handle loader.VkCommandPool
	device loader.VkDevice
}

func CreateCommandPool(allocator cgoalloc.Allocator, device resources.Device, o *CommandPoolOptions) (CommandPool, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var cmdPoolHandle loader.VkCommandPool
	res, err := device.Loader().VkCreateCommandPool(device.Handle(), (*loader.VkCommandPoolCreateInfo)(createInfo), nil, &cmdPoolHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanCommandPool{loader: device.Loader(), handle: cmdPoolHandle, device: device.Handle()}, res, nil
}

func (p *vulkanCommandPool) Handle() loader.VkCommandPool {
	return p.handle
}

func (p *vulkanCommandPool) Destroy() error {
	return p.loader.VkDestroyCommandPool(p.device, p.handle, nil)
}

func (p *vulkanCommandPool) DestroyBuffers(allocator cgoalloc.Allocator, buffers []CommandBuffer) error {
	bufferCount := len(buffers)
	if bufferCount == 0 {
		return nil
	}

	destroyPtr := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{})))
	defer allocator.Free(destroyPtr)

	destroySlice := ([]loader.VkCommandBuffer)(unsafe.Slice((*loader.VkCommandBuffer)(destroyPtr), bufferCount))
	for i := 0; i < bufferCount; i++ {
		destroySlice[i] = buffers[i].Handle()
	}

	return p.loader.VkFreeCommandBuffers(p.device, p.handle, loader.Uint32(bufferCount), (*loader.VkCommandBuffer)(destroyPtr))
}
