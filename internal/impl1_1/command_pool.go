package impl1_1

import (
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *DeviceVulkanDriver) TrimCommandPool(commandPool types.CommandPool, flags core1_1.CommandPoolTrimFlags) {
	if commandPool.Handle() == 0 {
		panic("commandPool was uninitialized")
	}
	
	v.LoaderObj.VkTrimCommandPool(commandPool.DeviceHandle(),
		commandPool.Handle(),
		loader.VkCommandPoolTrimFlags(flags),
	)
}
