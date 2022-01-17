package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	driver3 "github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type vulkanQueryPool struct {
	driver driver3.Driver
	handle driver3.VkQueryPool
	device driver3.VkDevice
}

func (p *vulkanQueryPool) Handle() driver3.VkQueryPool {
	return p.handle
}

func (p *vulkanQueryPool) Destroy(callbacks *AllocationCallbacks) {
	p.driver.VkDestroyQueryPool(p.device, p.handle, callbacks.Handle())
}

func (p *vulkanQueryPool) PopulateResults(firstQuery, queryCount int, resultSize, resultStride int, flags common.QueryResultFlags) ([]byte, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outBuffer := make([]byte, resultSize)

	inPointer := arena.Malloc(resultSize)

	res, err := p.driver.VkGetQueryPoolResults(p.device, p.handle, driver3.Uint32(firstQuery), driver3.Uint32(queryCount), driver3.Size(resultSize), inPointer, driver3.VkDeviceSize(resultStride), driver3.VkQueryResultFlags(flags))
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
