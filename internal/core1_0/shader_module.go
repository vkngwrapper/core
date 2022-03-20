package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanShaderModule struct {
	Driver             driver.Driver
	Device             driver.VkDevice
	ShaderModuleHandle driver.VkShaderModule

	MaximumAPIVersion common.APIVersion

	ShaderModule1_1 core1_1.ShaderModule
}

func (m *VulkanShaderModule) Handle() driver.VkShaderModule {
	return m.ShaderModuleHandle
}

func (m *VulkanShaderModule) Core1_1() core1_1.ShaderModule {
	return m.ShaderModule1_1
}

func (m *VulkanShaderModule) Destroy(callbacks *driver.AllocationCallbacks) {
	m.Driver.VkDestroyShaderModule(m.Device, m.ShaderModuleHandle, callbacks.Handle())
	m.Driver.ObjectStore().Delete(driver.VulkanHandle(m.ShaderModuleHandle), m)
}
