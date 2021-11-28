package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type vulkanShaderModule struct {
	driver Driver
	device VkDevice
	handle VkShaderModule
}

func (m *vulkanShaderModule) Handle() VkShaderModule {
	return m.handle
}

func (m *vulkanShaderModule) Destroy() {
	m.driver.VkDestroyShaderModule(m.device, m.handle, nil)
}
