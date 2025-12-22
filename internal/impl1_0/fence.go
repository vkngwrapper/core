package impl1_0

import (
	"fmt"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyFence(fence types.Fence, callbacks *driver.AllocationCallbacks) {
	if fence.Handle() == 0 {
		panic("fence was uninitialized")
	}

	v.Driver.VkDestroyFence(fence.DeviceHandle(), fence.Handle(), callbacks.Handle())
}

func (v *Vulkan) GetFenceStatus(fence types.Fence) (common.VkResult, error) {
	if fence.Handle() == 0 {
		return core1_0.VKErrorUnknown, fmt.Errorf("fence was uninitialized")
	}
	return v.Driver.VkGetFenceStatus(fence.DeviceHandle(), fence.Handle())
}
