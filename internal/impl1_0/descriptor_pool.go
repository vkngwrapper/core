package impl1_0

import (
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) DestroyDescriptorPool(descriptorPool core.DescriptorPool, callbacks *loader.AllocationCallbacks) {
	if !descriptorPool.Initialized() {
		panic("descriptorPool was uninitialized")
	}
	v.LoaderObj.VkDestroyDescriptorPool(descriptorPool.DeviceHandle(), descriptorPool.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) ResetDescriptorPool(descriptorPool core.DescriptorPool, flags core1_0.DescriptorPoolResetFlags) (common.VkResult, error) {
	if !descriptorPool.Initialized() {
		return core1_0.VKErrorUnknown, errors.New("descriptorPool was uninitialized")
	}
	return v.LoaderObj.VkResetDescriptorPool(descriptorPool.DeviceHandle(), descriptorPool.Handle(), loader.VkDescriptorPoolResetFlags(flags))
}
