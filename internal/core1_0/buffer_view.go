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

type VulkanBufferView struct {
	Driver           driver.Driver
	Device           driver.VkDevice
	BufferViewHandle driver.VkBufferView

	BufferView1_1 core1_1.BufferView

	MaximumAPIVersion common.APIVersion
}

func (v *VulkanBufferView) Handle() driver.VkBufferView {
	return v.BufferViewHandle
}

func (v *VulkanBufferView) Destroy(callbacks *driver.AllocationCallbacks) {
	v.Driver.VkDestroyBufferView(v.Device, v.BufferViewHandle, callbacks.Handle())
}

func (v *VulkanBufferView) Core1_1() core1_1.BufferView {
	return v.BufferView1_1
}
