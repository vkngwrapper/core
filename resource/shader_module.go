package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type vulkanShaderModule struct {
	loader *loader.Loader
	device loader.VkDevice
	handle loader.VkShaderModule
}

func (m *vulkanShaderModule) Handle() loader.VkShaderModule {
	return m.handle
}

func (m *vulkanShaderModule) Destroy() error {
	return m.loader.VkDestroyShaderModule(m.device, m.handle, nil)
}
