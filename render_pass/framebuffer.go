package render_pass

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resources"
	"github.com/CannibalVox/cgoparam"
)

type vulkanFramebuffer struct {
	loader loader.Loader
	device loader.VkDevice
	handle loader.VkFramebuffer
}

func CreateFrameBuffer(device resources.Device, o *FramebufferOptions) (Framebuffer, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var framebuffer loader.VkFramebuffer

	res, err := device.Loader().VkCreateFramebuffer(device.Handle(), (*loader.VkFramebufferCreateInfo)(createInfo), nil, &framebuffer)
	if err != nil {
		return nil, res, err
	}

	return &vulkanFramebuffer{loader: device.Loader(), device: device.Handle(), handle: framebuffer}, res, nil
}

func (b *vulkanFramebuffer) Handle() loader.VkFramebuffer {
	return b.handle
}

func (b *vulkanFramebuffer) Destroy() error {
	return b.loader.VkDestroyFramebuffer(b.device, b.handle, nil)
}
