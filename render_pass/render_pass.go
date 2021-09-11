package render_pass

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

type vulkanRenderPass struct {
	loader loader.Loader
	device loader.VkDevice
	handle loader.VkRenderPass
}

func CreateRenderPass(allocator cgoalloc.Allocator, device resource.Device, o *RenderPassOptions) (RenderPass, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var renderPass loader.VkRenderPass

	res, err := device.Loader().VkCreateRenderPass(device.Handle(), (*loader.VkRenderPassCreateInfo)(createInfo), nil, &renderPass)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &vulkanRenderPass{loader: device.Loader(), device: device.Handle(), handle: renderPass}, res, nil
}

func (p *vulkanRenderPass) Handle() loader.VkRenderPass {
	return p.handle
}

func (p *vulkanRenderPass) Destroy() error {
	return p.loader.VkDestroyRenderPass(p.device, p.handle, nil)
}
