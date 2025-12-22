package impl1_0

import (
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyDescriptorPool(descriptorPool types.DescriptorPool, callbacks *driver.AllocationCallbacks) {
	if descriptorPool.Handle() == 0 {
		panic("descriptorPool was uninitialized")
	}
	v.Driver.VkDestroyDescriptorPool(descriptorPool.DeviceHandle(), descriptorPool.Handle(), callbacks.Handle())
}

func (v *Vulkan) ResetDescriptorPool(descriptorPool types.DescriptorPool, flags core1_0.DescriptorPoolResetFlags) (common.VkResult, error) {
	if descriptorPool.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("descriptorPool was uninitialized")
	}
	return v.Driver.VkResetDescriptorPool(descriptorPool.DeviceHandle(), descriptorPool.Handle(), driver.VkDescriptorPoolResetFlags(flags))
}
