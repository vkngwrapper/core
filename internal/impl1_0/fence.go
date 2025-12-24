package impl1_0

import (
	"fmt"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *DeviceVulkanDriver) DestroyFence(fence types.Fence, callbacks *loader.AllocationCallbacks) {
	if fence.Handle() == 0 {
		panic("fence was uninitialized")
	}

	v.LoaderObj.VkDestroyFence(fence.DeviceHandle(), fence.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) GetFenceStatus(fence types.Fence) (common.VkResult, error) {
	if fence.Handle() == 0 {
		return core1_0.VKErrorUnknown, fmt.Errorf("fence was uninitialized")
	}
	return v.LoaderObj.VkGetFenceStatus(fence.DeviceHandle(), fence.Handle())
}
