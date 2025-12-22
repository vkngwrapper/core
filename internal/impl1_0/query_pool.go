package impl1_0

import (
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyQueryPool(queryPool types.QueryPool, callbacks *driver.AllocationCallbacks) {
	if queryPool.Handle() == 0 {
		panic("queryPool was uninitialized")
	}
	v.Driver.VkDestroyQueryPool(queryPool.DeviceHandle(), queryPool.Handle(), callbacks.Handle())
}

func (v *Vulkan) GetQueryPoolResults(queryPool types.QueryPool, firstQuery, queryCount int, results []byte, resultStride int, flags core1_0.QueryResultFlags) (common.VkResult, error) {
	if queryPool.Handle() == 0 {
		return core1_0.VKErrorUnknown, fmt.Errorf("queryPool was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	resultSize := len(results)

	inPointer := arena.Malloc(resultSize)

	res, err := v.Driver.VkGetQueryPoolResults(queryPool.DeviceHandle(), queryPool.Handle(), driver.Uint32(firstQuery), driver.Uint32(queryCount), driver.Size(resultSize), inPointer, driver.VkDeviceSize(resultStride), driver.VkQueryResultFlags(flags))
	if err != nil {
		return res, err
	}

	inBuffer := ([]byte)(unsafe.Slice((*byte)(inPointer), resultSize))
	copy(results, inBuffer)

	return res, nil
}
