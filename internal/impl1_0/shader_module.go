package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanShaderModule is an implementation of the ShaderModule interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanShaderModule struct {
	DeviceDriver       driver.Driver
	Device             driver.VkDevice
	ShaderModuleHandle driver.VkShaderModule

	MaximumAPIVersion common.APIVersion
}

func (m *VulkanShaderModule) Handle() driver.VkShaderModule {
	return m.ShaderModuleHandle
}

func (m *VulkanShaderModule) Driver() driver.Driver {
	return m.DeviceDriver
}

func (m *VulkanShaderModule) DeviceHandle() driver.VkDevice {
	return m.Device
}

func (m *VulkanShaderModule) APIVersion() common.APIVersion {
	return m.MaximumAPIVersion
}

func (m *VulkanShaderModule) Destroy(callbacks *driver.AllocationCallbacks) {
	m.DeviceDriver.VkDestroyShaderModule(m.Device, m.ShaderModuleHandle, callbacks.Handle())
}
