package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDeviceMemory struct {
	DeviceDriver       driver.Driver
	Device             driver.VkDevice
	DeviceMemoryHandle driver.VkDeviceMemory

	Size int
}
