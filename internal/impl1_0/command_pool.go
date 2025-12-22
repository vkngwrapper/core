package impl1_0

import (
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyCommandPool(commandPool types.CommandPool, callbacks *driver.AllocationCallbacks) {
	if commandPool.Handle() == 0 {
		panic("commandPool cannot be uninitialized")
	}
	v.Driver.VkDestroyCommandPool(commandPool.DeviceHandle(), commandPool.Handle(), callbacks.Handle())
}

func (v *Vulkan) ResetCommandPool(commandPool types.CommandPool, flags core1_0.CommandPoolResetFlags) (common.VkResult, error) {
	if commandPool.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("commandPool cannot be uninitialized")
	}

	return v.Driver.VkResetCommandPool(commandPool.DeviceHandle(), commandPool.Handle(), driver.VkCommandPoolResetFlags(flags))
}
