package impl1_1

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) TrimCommandPool(commandPool core1_0.CommandPool, flags core1_1.CommandPoolTrimFlags) {
	if !commandPool.Initialized() {
		panic("commandPool was uninitialized")
	}

	v.LoaderObj.VkTrimCommandPool(commandPool.DeviceHandle(),
		commandPool.Handle(),
		loader.VkCommandPoolTrimFlags(flags),
	)
}
