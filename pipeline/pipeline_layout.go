package pipeline

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

type PipelineLayout struct {
	loader *loader.Loader
	device loader.VkDevice
	handle loader.VkPipelineLayout
}

func (l *PipelineLayout) Handle() loader.VkPipelineLayout {
	return l.handle
}

func (l *PipelineLayout) Destroy() error {
	return l.loader.VkDestroyPipelineLayout(l.device, l.handle, nil)
}

func CreatePipelineLayout(allocator cgoalloc.Allocator, device *resource.Device, o *PipelineLayoutOptions) (*PipelineLayout, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var pipelineLayout loader.VkPipelineLayout
	res, err := device.Loader().VkCreatePipelineLayout(device.Handle(), (*loader.VkPipelineLayoutCreateInfo)(createInfo), nil, &pipelineLayout)
	if err != nil {
		return nil, res, err
	}

	return &PipelineLayout{loader: device.Loader(), handle: pipelineLayout, device: device.Handle()}, res, nil
}
