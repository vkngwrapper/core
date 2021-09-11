package render_pass

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resources"
	"github.com/CannibalVox/cgoparam"
)

type vulkanRenderPass struct {
	loader loader.Loader
	device loader.VkDevice
	handle loader.VkRenderPass
}

func CreateRenderPass(device resources.Device, o *RenderPassOptions) (RenderPass, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

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
