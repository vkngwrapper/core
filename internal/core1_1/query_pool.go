package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanQueryPool struct {
	Driver          driver.Driver
	QueryPoolHandle driver.VkQueryPool
	Device          driver.VkDevice
}
