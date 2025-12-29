package impl1_2

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) ResetQueryPool(queryPool core1_0.QueryPool, firstQuery, queryCount int) {
	if !queryPool.Initialized() {
		panic("queryPool cannot be uninitialized")
	}

	v.LoaderObj.VkResetQueryPool(queryPool.DeviceHandle(), queryPool.Handle(), loader.Uint32(firstQuery), loader.Uint32(queryCount))
}
