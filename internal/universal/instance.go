package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanInstance struct {
	driver driver.Driver
	handle driver.VkInstance
}

func (i *VulkanInstance) Driver() driver.Driver {
	return i.driver
}

func (i *VulkanInstance) Handle() driver.VkInstance {
	return i.handle
}

func (i *VulkanInstance) Destroy(callbacks *driver.AllocationCallbacks) {
	i.driver.VkDestroyInstance(i.handle, callbacks.Handle())
	i.driver.Destroy()
}
