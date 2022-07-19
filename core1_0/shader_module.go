package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
)

// VulkanShaderModule is an implementation of the ShaderModule interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanShaderModule struct {
	deviceDriver       driver.Driver
	device             driver.VkDevice
	shaderModuleHandle driver.VkShaderModule

	maximumAPIVersion common.APIVersion
}

func (m *VulkanShaderModule) Handle() driver.VkShaderModule {
	return m.shaderModuleHandle
}

func (m *VulkanShaderModule) Driver() driver.Driver {
	return m.deviceDriver
}

func (m *VulkanShaderModule) DeviceHandle() driver.VkDevice {
	return m.device
}

func (m *VulkanShaderModule) APIVersion() common.APIVersion {
	return m.maximumAPIVersion
}

func (m *VulkanShaderModule) Destroy(callbacks *driver.AllocationCallbacks) {
	m.deviceDriver.VkDestroyShaderModule(m.device, m.shaderModuleHandle, callbacks.Handle())
	m.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(m.shaderModuleHandle))
}
