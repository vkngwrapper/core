package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resources"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type vulkanPipeline struct {
	loader loader.Loader
	device loader.VkDevice
	handle loader.VkPipeline
}

func CreateGraphicsPipelines(device resources.Device, o []*Options) ([]Pipeline, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtrUnsafe := arena.Malloc(pipelineCount * C.sizeof_struct_VkGraphicsPipelineCreateInfo)
	pipelineCreateInfosSlice := ([]C.VkGraphicsPipelineCreateInfo)(unsafe.Slice((*C.VkGraphicsPipelineCreateInfo)(pipelineCreateInfosPtrUnsafe), pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		next, err := core.AllocNext(arena, o[i])
		if err != nil {
			return nil, loader.VKErrorUnknown, err
		}

		err = o[i].populate(arena, &pipelineCreateInfosSlice[i], next)
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
		output = append(output, &vulkanPipeline{loader: device.Loader(), device: device.Handle(), handle: pipelineSlice[i]})
	}

	return output, res, nil
}

func (p *vulkanPipeline) Handle() loader.VkPipeline {
	return p.handle
}

func (p *vulkanPipeline) Destroy() error {
	return p.loader.VkDestroyPipeline(p.device, p.handle, nil)
}
