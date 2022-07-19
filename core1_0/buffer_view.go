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

// VulkanBufferView is an implementation of the BufferView interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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
