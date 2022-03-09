package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorPool struct {
	DeviceDriver         driver.Driver
	DescriptorPoolHandle driver.VkDescriptorPool
	Device               driver.VkDevice
}
