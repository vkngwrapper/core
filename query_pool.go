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

type vulkanQueryPool struct {
	driver Driver
	handle VkQueryPool
	device VkDevice
}

func (p *vulkanQueryPool) Handle() VkQueryPool {
	return p.handle
}

func (p *vulkanQueryPool) Destroy() {
	p.driver.VkDestroyQueryPool(p.device, p.handle, nil)
}

func (p *vulkanQueryPool) PopulateResults(firstQuery, queryCount int, resultSize, resultStride int, flags common.QueryResultFlags) ([]byte, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outBuffer := make([]byte, resultSize)

	inPointer := arena.Malloc(resultSize)

	res, err := p.driver.VkGetQueryPoolResults(p.device, p.handle, Uint32(firstQuery), Uint32(queryCount), Size(resultSize), inPointer, VkDeviceSize(resultStride), VkQueryResultFlags(flags))
	if err != nil {
		return nil, res, err
	}

	inBuffer := ([]byte)(unsafe.Slice((*byte)(inPointer), resultSize))
	copy(outBuffer, inBuffer)

	return outBuffer, res, nil
}

type QueryPoolOptions struct {
	QueryType          common.QueryType
	QueryCount         int
	PipelineStatistics common.PipelineStatistics

	common.HaveNext
}

func (o *QueryPoolOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkQueryPoolCreateInfo)(allocator.Malloc(C.sizeof_struct_VkQueryPoolCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = 0
	createInfo.queryType = C.VkQueryType(o.QueryType)
	createInfo.queryCount = C.uint32_t(o.QueryCount)
	createInfo.pipelineStatistics = C.VkQueryPipelineStatisticFlags(o.PipelineStatistics)

	return unsafe.Pointer(createInfo), nil
}
