package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
)

type vulkanRenderPass struct {
	driver Driver
	device VkDevice
	handle VkRenderPass
}

func CreateRenderPass(device Device, o *RenderPassOptions) (RenderPass, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var renderPass VkRenderPass

	res, err := device.Driver().VkCreateRenderPass(device.Handle(), (*VkRenderPassCreateInfo)(createInfo), nil, &renderPass)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &vulkanRenderPass{driver: device.Driver(), device: device.Handle(), handle: renderPass}, res, nil
}

func (p *vulkanRenderPass) Handle() VkRenderPass {
	return p.handle
}

func (p *vulkanRenderPass) Destroy() error {
	return p.driver.VkDestroyRenderPass(p.device, p.handle, nil)
}
