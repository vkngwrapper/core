package pipeline

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/resource"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type PipelineHandle C.VkPipeline
type Pipeline struct {
	device C.VkDevice
	handle C.VkPipeline
}

func CreateGraphicsPipelines(allocator cgoalloc.Allocator, device *resource.Device, o []*Options) ([]*Pipeline, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	deviceHandle := (C.VkDevice)(unsafe.Pointer(device.Handle()))
	pipelineCount := len(o)

	pipelineCreateInfosPtr := (*C.VkGraphicsPipelineCreateInfo)(arena.Malloc(pipelineCount * C.sizeof_struct_VkGraphicsPipelineCreateInfo))
	pipelineCreateInfosSlice := ([]C.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pipelineCreateInfosPtr, pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		err := o[i].populate(arena, &pipelineCreateInfosSlice[i])
		if err != nil {
			return nil, core.VKErrorUnknown, err
		}
	}

	pipelinePtr := (*C.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]C.VkPipeline{}))))

	res := core.Result(C.vkCreateGraphicsPipelines(deviceHandle, (C.VkPipelineCache)(C.VK_NULL_HANDLE), C.uint32_t(pipelineCount), pipelineCreateInfosPtr, nil, pipelinePtr))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	var output []*Pipeline
	pipelineSlice := ([]C.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		output = append(output, &Pipeline{device: deviceHandle, handle: pipelineSlice[i]})
	}

	return output, res, nil
}

func (p *Pipeline) Handle() PipelineHandle {
	return PipelineHandle(p.handle)
}

func (p *Pipeline) Destroy() {
	C.vkDestroyPipeline(p.device, p.handle, nil)
}
