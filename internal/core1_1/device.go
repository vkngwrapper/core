package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	internal_core1_0 "github.com/CannibalVox/VKng/core/internal/core1_0"
)

type VulkanDevice struct {
	internal_core1_0.VulkanDevice
}

func PromoteDevice(device core1_0.Device) core1_1.Device {
	goodDevice, ok := device.(core1_1.Device)
	if ok {
		return goodDevice
	}

	oldVulkanDevice, ok := device.(*internal_core1_0.VulkanDevice)
	if ok && oldVulkanDevice.MaximumAPIVersion.IsAtLeast(common.Vulkan1_0) {
		return &VulkanDevice{*oldVulkanDevice}
	}

	return nil
}
