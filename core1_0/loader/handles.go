package loader

import (
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/VKng/core/internal/universal"
)

func CreateImageFromHandles(handle driver.VkImage, device driver.VkDevice, driver driver.Driver) iface.Image {
	return universal.CreateImageFromHandles(handle, device, driver)
}
