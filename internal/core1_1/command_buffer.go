package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanCommandBuffer struct {
	DeviceDriver        driver.Driver
	Device              driver.VkDevice
	CommandPool         driver.VkCommandPool
	CommandBufferHandle driver.VkCommandBuffer
}
