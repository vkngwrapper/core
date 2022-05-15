package internal1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanInstance struct {
	InstanceDriver driver.Driver
	InstanceHandle driver.VkInstance
	MaximumVersion common.APIVersion
}

func (i *VulkanInstance) Driver() driver.Driver {
	return i.InstanceDriver
}

func (i *VulkanInstance) Handle() driver.VkInstance {
	return i.InstanceHandle
}

func (i *VulkanInstance) APIVersion() common.APIVersion {
	return i.MaximumVersion
}

func (i *VulkanInstance) Destroy(callbacks *driver.AllocationCallbacks) {
	i.InstanceDriver.VkDestroyInstance(i.InstanceHandle, callbacks.Handle())
	i.InstanceDriver.ObjectStore().Delete(driver.VulkanHandle(i.InstanceHandle))
	i.InstanceDriver.Destroy()
}
