package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type ShaderModuleHandle C.VkShaderModule
type ShaderModule struct {
	device C.VkDevice
	handle C.VkShaderModule
}

func (m *ShaderModule) Handle() ShaderModuleHandle {
	return ShaderModuleHandle(m.handle)
}

func (m *ShaderModule) Destroy() {
	C.vkDestroyShaderModule(m.device, m.handle, nil)
}
