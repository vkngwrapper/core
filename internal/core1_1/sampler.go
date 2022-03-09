package core1_1

import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSampler struct {
	Device        driver.VkDevice
	Driver        driver.Driver
	SamplerHandle driver.VkSampler
}
