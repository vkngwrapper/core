package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyPipelineCache(pipelineCache types.PipelineCache, callbacks *driver.AllocationCallbacks) {
	if pipelineCache.Handle() == 0 {
		panic("pipelineCache was uninitialized")
	}
	v.Driver.VkDestroyPipelineCache(pipelineCache.DeviceHandle(), pipelineCache.Handle(), callbacks.Handle())
}

func (v *Vulkan) GetPipelineCacheData(pipelineCache types.PipelineCache) ([]byte, common.VkResult, error) {
	if pipelineCache.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, fmt.Errorf("pipelineCache was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	cacheSizePtr := arena.Malloc(int(unsafe.Sizeof(C.size_t(0))))
	cacheSize := (*driver.Size)(cacheSizePtr)

	res, err := v.Driver.VkGetPipelineCacheData(pipelineCache.DeviceHandle(), pipelineCache.Handle(), cacheSize, nil)
	if err != nil {
		return nil, res, err
	}

	cacheDataPtr := arena.Malloc(int(*cacheSize))

	res, err = v.Driver.VkGetPipelineCacheData(pipelineCache.DeviceHandle(), pipelineCache.Handle(), cacheSize, cacheDataPtr)
	if err != nil {
		return nil, res, err
	}

	outData := make([]byte, *cacheSize)
	inData := ([]byte)(unsafe.Slice((*byte)(cacheDataPtr), int(*cacheSize)))
	copy(outData, inData)

	return outData, res, nil
}

func (v *Vulkan) MergePipelineCaches(dstCaches types.PipelineCache, srcCaches []types.PipelineCache) (common.VkResult, error) {
	if dstCaches.Handle() == 0 {
		panic("dstCaches was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	srcCount := len(srcCaches)
	srcPtr := (*driver.VkPipelineCache)(arena.Malloc(srcCount * int(unsafe.Sizeof([1]driver.VkPipelineCache{}))))
	srcSlice := ([]driver.VkPipelineCache)(unsafe.Slice(srcPtr, srcCount))

	for i := 0; i < srcCount; i++ {
		if srcCaches[i].Handle() == 0 {
			panic(fmt.Sprintf("elements of srcCaches cannot be uninitialized- element %d is uninitialized", i))
		}
		srcSlice[i] = srcCaches[i].Handle()
	}

	return v.Driver.VkMergePipelineCaches(dstCaches.DeviceHandle(), dstCaches.Handle(), driver.Uint32(srcCount), srcPtr)
}
