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
	v.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(v.BufferViewHandle))
}
