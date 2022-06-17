package internal1_0

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
	m.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(m.ShaderModuleHandle))
}
