package impl1_1

import (
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) DestroySamplerYcbcrConversion(conversion core1_1.SamplerYcbcrConversion, allocator *loader.AllocationCallbacks) {
	v.LoaderObj.VkDestroySamplerYcbcrConversion(conversion.DeviceHandle(), conversion.Handle(), allocator.Handle())
}
