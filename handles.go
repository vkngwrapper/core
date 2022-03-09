package core

import (
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	internal1_0 "github.com/CannibalVox/VKng/core/internal/core1_0"
)

func CreateImageFromHandles(handle driver.VkImage, device driver.VkDevice, driver driver.Driver) core1_0.Image {
	return &internal1_0.VulkanImage{ImageHandle: handle, Device: device, Driver: driver}
}
