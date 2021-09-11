package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type VulkanShaderModule struct {
	loader *loader.Loader
	device loader.VkDevice
	handle loader.VkShaderModule
}

func (m *VulkanShaderModule) Handle() loader.VkShaderModule {
	return m.handle
}

func (m *VulkanShaderModule) Destroy() error {
	return m.loader.VkDestroyShaderModule(m.device, m.handle, nil)
}
