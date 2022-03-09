package core1_1

import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorSetLayout struct {
	Driver                    driver.Driver
	Device                    driver.VkDevice
	DescriptorSetLayoutHandle driver.VkDescriptorSetLayout
}
