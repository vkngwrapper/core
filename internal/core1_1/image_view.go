package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanImageView struct {
	Driver          driver.Driver
	ImageViewHandle driver.VkImageView
	Device          driver.VkDevice
}
