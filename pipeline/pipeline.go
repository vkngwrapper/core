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
	"unsafe"
)

type VulkanPipeline struct {
	loader *loader.Loader
	device loader.VkDevice
	handle loader.VkPipeline
}

func CreateGraphicsPipelines(allocator cgoalloc.Allocator, device resource.Device, o []*Options) ([]Pipeline, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	pipelineCount := len(o)

	pipelineCreateInfosPtrUnsafe := arena.Malloc(pipelineCount * C.sizeof_struct_VkGraphicsPipelineCreateInfo)
	pipelineCreateInfosSlice := ([]C.VkGraphicsPipelineCreateInfo)(unsafe.Slice((*C.VkGraphicsPipelineCreateInfo)(pipelineCreateInfosPtrUnsafe), pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		err := o[i].populate(arena, &pipelineCreateInfosSlice[i])
		if err != nil {
			return nil, loader.VKErrorUnknown, err
		}
	}

	pipelinePtr := (*loader.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]loader.VkPipeline{}))))

	res, err := device.Loader().VkCreateGraphicsPipelines(device.Handle(), nil, loader.Uint32(pipelineCount), (*loader.VkGraphicsPipelineCreateInfo)(pipelineCreateInfosPtrUnsafe), nil, pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []Pipeline
	pipelineSlice := ([]loader.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		output = append(output, &VulkanPipeline{loader: device.Loader(), device: device.Handle(), handle: pipelineSlice[i]})
	}

	return output, res, nil
}

func (p *VulkanPipeline) Handle() loader.VkPipeline {
	return p.handle
}

func (p *VulkanPipeline) Destroy() error {
	return p.loader.VkDestroyPipeline(p.device, p.handle, nil)
}
