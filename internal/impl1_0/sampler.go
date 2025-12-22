package impl1_0

import (
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroySampler(sampler types.Sampler, callbacks *driver.AllocationCallbacks) {
	if sampler.Handle() == 0 {
		panic("sampler was uninitialized")
	}
	v.Driver.VkDestroySampler(sampler.DeviceHandle(), sampler.Handle(), callbacks.Handle())
}
