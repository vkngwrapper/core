package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanShaderModule struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkShaderModule
}

func (m *VulkanShaderModule) Handle() driver.VkShaderModule {
	return m.handle
}

func (m *VulkanShaderModule) Destroy(callbacks *driver.AllocationCallbacks) {
	m.driver.VkDestroyShaderModule(m.device, m.handle, callbacks.Handle())
}
