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
	Driver           driver.Driver
	Device           driver.VkDevice
	BufferViewHandle driver.VkBufferView

	MaximumAPIVersion common.APIVersion
}

func (v *VulkanBufferView) Handle() driver.VkBufferView {
	return v.BufferViewHandle
}

func (v *VulkanBufferView) Destroy(callbacks *driver.AllocationCallbacks) {
	v.Driver.VkDestroyBufferView(v.Device, v.BufferViewHandle, callbacks.Handle())
	v.Driver.ObjectStore().Delete(driver.VulkanHandle(v.BufferViewHandle), v)
}
