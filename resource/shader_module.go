package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type ShaderModule struct {
	loader *loader.Loader
	device loader.VkDevice
	handle loader.VkShaderModule
}

func (m *ShaderModule) Handle() loader.VkShaderModule {
	return m.handle
}

func (m *ShaderModule) Destroy() error {
	return m.loader.VkDestroyShaderModule(m.device, m.handle, nil)
}
