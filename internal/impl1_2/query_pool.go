package impl1_2

import (
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *DeviceVulkanDriver) ResetQueryPool(queryPool types.QueryPool, firstQuery, queryCount int) {
	if queryPool.Handle() == 0 {
		panic("queryPool cannot be uninitialized")
	}

	v.LoaderObj.VkResetQueryPool(queryPool.DeviceHandle(), queryPool.Handle(), loader.Uint32(firstQuery), loader.Uint32(queryCount))
}
