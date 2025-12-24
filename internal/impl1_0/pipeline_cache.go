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
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *DeviceVulkanDriver) DestroyPipelineCache(pipelineCache types.PipelineCache, callbacks *loader.AllocationCallbacks) {
	if pipelineCache.Handle() == 0 {
		panic("pipelineCache was uninitialized")
	}
	v.LoaderObj.VkDestroyPipelineCache(pipelineCache.DeviceHandle(), pipelineCache.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) GetPipelineCacheData(pipelineCache types.PipelineCache) ([]byte, common.VkResult, error) {
	if pipelineCache.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, fmt.Errorf("pipelineCache was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	cacheSizePtr := arena.Malloc(int(unsafe.Sizeof(C.size_t(0))))
	cacheSize := (*loader.Size)(cacheSizePtr)

	res, err := v.LoaderObj.VkGetPipelineCacheData(pipelineCache.DeviceHandle(), pipelineCache.Handle(), cacheSize, nil)
	if err != nil {
		return nil, res, err
	}

	cacheDataPtr := arena.Malloc(int(*cacheSize))

	res, err = v.LoaderObj.VkGetPipelineCacheData(pipelineCache.DeviceHandle(), pipelineCache.Handle(), cacheSize, cacheDataPtr)
	if err != nil {
		return nil, res, err
	}

	outData := make([]byte, *cacheSize)
	inData := ([]byte)(unsafe.Slice((*byte)(cacheDataPtr), int(*cacheSize)))
	copy(outData, inData)

	return outData, res, nil
}

func (v *DeviceVulkanDriver) MergePipelineCaches(dstCaches types.PipelineCache, srcCaches ...types.PipelineCache) (common.VkResult, error) {
	if dstCaches.Handle() == 0 {
		panic("dstCaches was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	srcCount := len(srcCaches)
	srcPtr := (*loader.VkPipelineCache)(arena.Malloc(srcCount * int(unsafe.Sizeof([1]loader.VkPipelineCache{}))))
	srcSlice := ([]loader.VkPipelineCache)(unsafe.Slice(srcPtr, srcCount))

	for i := 0; i < srcCount; i++ {
		if srcCaches[i].Handle() == 0 {
			panic(fmt.Sprintf("elements of srcCaches cannot be uninitialized- element %d is uninitialized", i))
		}
		srcSlice[i] = srcCaches[i].Handle()
	}

	return v.LoaderObj.VkMergePipelineCaches(dstCaches.DeviceHandle(), dstCaches.Handle(), loader.Uint32(srcCount), srcPtr)
}
