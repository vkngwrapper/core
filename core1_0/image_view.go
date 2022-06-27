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

type VulkanImageView struct {
	deviceDriver    driver.Driver
	imageViewHandle driver.VkImageView
	device          driver.VkDevice

	maximumAPIVersion common.APIVersion
}

func (v *VulkanImageView) Handle() driver.VkImageView {
	return v.imageViewHandle
}

func (v *VulkanImageView) Driver() driver.Driver {
	return v.deviceDriver
}

func (v *VulkanImageView) DeviceHandle() driver.VkDevice {
	return v.device
}

func (v *VulkanImageView) APIVersion() common.APIVersion {
	return v.maximumAPIVersion
}

func (v *VulkanImageView) Destroy(callbacks *driver.AllocationCallbacks) {
	v.deviceDriver.VkDestroyImageView(v.device, v.imageViewHandle, callbacks.Handle())
	v.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(v.imageViewHandle))
}
