package impl1_1

import (
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *DeviceVulkanDriver) DestroySamplerYcbcrConversion(conversion types.SamplerYcbcrConversion, allocator *loader.AllocationCallbacks) {
	v.LoaderObj.VkDestroySamplerYcbcrConversion(conversion.DeviceHandle(), conversion.Handle(), allocator.Handle())
}
