package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanBufferView is an implementation of the BufferView interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanBufferView struct {
	DeviceDriver     driver.Driver
	Device           driver.VkDevice
	BufferViewHandle driver.VkBufferView

	MaximumAPIVersion common.APIVersion
}

func (v *VulkanBufferView) Handle() driver.VkBufferView {
	return v.BufferViewHandle
}

func (v *VulkanBufferView) DeviceHandle() driver.VkDevice {
	return v.Device
}

func (v *VulkanBufferView) Driver() driver.Driver {
	return v.DeviceDriver
}

func (v *VulkanBufferView) APIVersion() common.APIVersion {
	return v.MaximumAPIVersion
}

func (v *VulkanBufferView) Destroy(callbacks *driver.AllocationCallbacks) {
	v.DeviceDriver.VkDestroyBufferView(v.Device, v.BufferViewHandle, callbacks.Handle())
}
