package impl1_0

import (
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyShaderModule(shaderModule types.ShaderModule, callbacks *driver.AllocationCallbacks) {
	if shaderModule.Handle() == 0 {
		panic("shaderModule is uninitialized")
	}
	v.Driver.VkDestroyShaderModule(shaderModule.DeviceHandle(), shaderModule.Handle(), callbacks.Handle())
}
