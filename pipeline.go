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

type vulkanPipeline struct {
	driver Driver
	device VkDevice
	handle VkPipeline
}

func CreateGraphicsPipelines(device Device, o []*Options) ([]Pipeline, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtrUnsafe := arena.Malloc(pipelineCount * C.sizeof_struct_VkGraphicsPipelineCreateInfo)
	pipelineCreateInfosSlice := ([]C.VkGraphicsPipelineCreateInfo)(unsafe.Slice((*C.VkGraphicsPipelineCreateInfo)(pipelineCreateInfosPtrUnsafe), pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		next, err := common.AllocNext(arena, o[i])
		if err != nil {
			return nil, VKErrorUnknown, err
		}

		err = o[i].populate(arena, &pipelineCreateInfosSlice[i], next)
		if err != nil {
			return nil, VKErrorUnknown, err
		}
	}

	pipelinePtr := (*VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]VkPipeline{}))))

	res, err := device.Driver().VkCreateGraphicsPipelines(device.Handle(), nil, Uint32(pipelineCount), (*VkGraphicsPipelineCreateInfo)(pipelineCreateInfosPtrUnsafe), nil, pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []Pipeline
	pipelineSlice := ([]VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		output = append(output, &vulkanPipeline{driver: device.Driver(), device: device.Handle(), handle: pipelineSlice[i]})
	}

	return output, res, nil
}

func (p *vulkanPipeline) Handle() VkPipeline {
	return p.handle
}

func (p *vulkanPipeline) Destroy() error {
	return p.driver.VkDestroyPipeline(p.device, p.handle, nil)
}
