package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanRenderPass struct {
	Driver           driver.Driver
	Device           driver.VkDevice
	RenderPassHandle driver.VkRenderPass
}