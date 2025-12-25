package impl1_1

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) DestroySamplerYcbcrConversion(conversion core.SamplerYcbcrConversion, allocator *loader.AllocationCallbacks) {
	v.LoaderObj.VkDestroySamplerYcbcrConversion(conversion.DeviceHandle(), conversion.Handle(), allocator.Handle())
}
