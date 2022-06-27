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

type VulkanBufferView struct {
	deviceDriver     driver.Driver
	device           driver.VkDevice
	bufferViewHandle driver.VkBufferView

	maximumAPIVersion common.APIVersion
}

func (v *VulkanBufferView) Handle() driver.VkBufferView {
	return v.bufferViewHandle
}

func (v *VulkanBufferView) DeviceHandle() driver.VkDevice {
	return v.device
}

func (v *VulkanBufferView) Driver() driver.Driver {
	return v.deviceDriver
}

func (v *VulkanBufferView) APIVersion() common.APIVersion {
	return v.maximumAPIVersion
}

func (v *VulkanBufferView) Destroy(callbacks *driver.AllocationCallbacks) {
	v.deviceDriver.VkDestroyBufferView(v.device, v.bufferViewHandle, callbacks.Handle())
	v.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(v.bufferViewHandle))
}
