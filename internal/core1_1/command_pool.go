package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanCommandPool struct {
	DeviceDriver      driver.Driver
	CommandPoolHandle driver.VkCommandPool
	DeviceHandle      driver.VkDevice
}
