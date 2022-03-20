package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanPhysicalDevice struct {
	InstanceDriver       driver.Driver
	PhysicalDeviceHandle driver.VkPhysicalDevice
}
