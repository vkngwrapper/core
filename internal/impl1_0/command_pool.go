package impl1_0

import (
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *DeviceVulkanDriver) DestroyCommandPool(commandPool types.CommandPool, callbacks *loader.AllocationCallbacks) {
	if commandPool.Handle() == 0 {
		panic("commandPool cannot be uninitialized")
	}
	v.LoaderObj.VkDestroyCommandPool(commandPool.DeviceHandle(), commandPool.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) ResetCommandPool(commandPool types.CommandPool, flags core1_0.CommandPoolResetFlags) (common.VkResult, error) {
	if commandPool.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("commandPool cannot be uninitialized")
	}

	return v.LoaderObj.VkResetCommandPool(commandPool.DeviceHandle(), commandPool.Handle(), loader.VkCommandPoolResetFlags(flags))
}
