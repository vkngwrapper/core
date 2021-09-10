package commands

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resource"
	"github.com/CannibalVox/cgoalloc"
)

type CommandPool struct {
	loader *loader.Loader
	handle loader.VkCommandPool
	device loader.VkDevice
}

func CreateCommandPool(allocator cgoalloc.Allocator, device *resource.Device, o *CommandPoolOptions) (*CommandPool, loader.VkResult, error) {
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

	return &CommandPool{loader: device.Loader(), handle: cmdPoolHandle, device: device.Handle()}, res, nil
}

func (p *CommandPool) Handle() loader.VkCommandPool {
	return p.handle
}

func (p *CommandPool) Destroy() error {
	return p.loader.VkDestroyCommandPool(p.device, p.handle, nil)
}
