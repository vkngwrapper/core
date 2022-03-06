package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanShaderModule struct {
	Driver             driver.Driver
	Device             driver.VkDevice
	ShaderModuleHandle driver.VkShaderModule

	MaximumAPIVersion common.APIVersion
}

func (m *VulkanShaderModule) Handle() driver.VkShaderModule {
	return m.ShaderModuleHandle
}

func (m *VulkanShaderModule) Destroy(callbacks *driver.AllocationCallbacks) {
	m.Driver.VkDestroyShaderModule(m.Device, m.ShaderModuleHandle, callbacks.Handle())
}
