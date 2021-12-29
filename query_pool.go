package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"encoding/binary"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"reflect"
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

func (p *vulkanQueryPool) PopulateResults(firstQuery, queryCount int, results interface{}, flags common.QueryResultFlags) (VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	resultsSize := binary.Size(results)
	if resultsSize < 0 {
		return VKErrorUnknown, errors.New("could not determine size of results with binary.Size- data must be strongly typed and strongly sized")
	}
	stride := resultsSize
	resultsVal := reflect.ValueOf(results)
	resultsPtr := unsafe.Pointer(&results)

	if resultsVal.Kind() == reflect.Slice {
		idxVal := resultsVal.Index(0)
		stride = binary.Size(idxVal.Interface())
		resultsPtr = unsafe.Pointer(idxVal.UnsafeAddr())
	} else if resultsVal.Kind() == reflect.Ptr {
		resultsPtr = unsafe.Pointer(resultsVal.Elem().UnsafeAddr())
	}

	outBuffer := ([]byte)(unsafe.Slice((*byte)(resultsPtr), resultsSize))
	inPointer := arena.Malloc(resultsSize)

	res, err := p.driver.VkGetQueryPoolResults(p.device, p.handle, Uint32(firstQuery), Uint32(queryCount), Size(resultsSize), inPointer, VkDeviceSize(stride), VkQueryResultFlags(flags))
	if err != nil {
		return res, err
	}

	inBuffer := ([]byte)(unsafe.Slice((*byte)(inPointer), resultsSize))
	copy(outBuffer, inBuffer)

	return res, nil
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
