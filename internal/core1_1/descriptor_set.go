package core1_1

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorSet struct {
	DescriptorSetHandle driver.VkDescriptorSet
	DeviceDriver        driver.Driver
	Device              driver.VkDevice
	DescriptorPool      driver.VkDescriptorPool
}
