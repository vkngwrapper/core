package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanEvent struct {
	EventHandle driver.VkEvent
	Device      driver.VkDevice
	Driver      driver.Driver
}
