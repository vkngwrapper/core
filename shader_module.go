package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	driver3 "github.com/CannibalVox/VKng/core/driver"
)

type vulkanShaderModule struct {
	driver driver3.Driver
	device driver3.VkDevice
	handle driver3.VkShaderModule
}

func (m *vulkanShaderModule) Handle() driver3.VkShaderModule {
	return m.handle
}

func (m *vulkanShaderModule) Destroy(callbacks *AllocationCallbacks) {
	m.driver.VkDestroyShaderModule(m.device, m.handle, callbacks.Handle())
}
