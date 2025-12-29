package impl1_0

import (
	"fmt"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) DestroyFence(fence core1_0.Fence, callbacks *loader.AllocationCallbacks) {
	if !fence.Initialized() {
		panic("fence was uninitialized")
	}

	v.LoaderObj.VkDestroyFence(fence.DeviceHandle(), fence.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) GetFenceStatus(fence core1_0.Fence) (common.VkResult, error) {
	if !fence.Initialized() {
		return core1_0.VKErrorUnknown, fmt.Errorf("fence was uninitialized")
	}
	return v.LoaderObj.VkGetFenceStatus(fence.DeviceHandle(), fence.Handle())
}
