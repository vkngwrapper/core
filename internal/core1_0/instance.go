package internal1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanInstance struct {
	InstanceDriver driver.Driver
	InstanceHandle driver.VkInstance
	MaximumVersion common.APIVersion

	Instance1_1 core1_1.Instance
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

func (i *VulkanInstance) Core1_1() core1_1.Instance {
	return i.Instance1_1
}

func (i *VulkanInstance) Destroy(callbacks *driver.AllocationCallbacks) {
	i.InstanceDriver.VkDestroyInstance(i.InstanceHandle, callbacks.Handle())
	i.InstanceDriver.ObjectStore().Delete(driver.VulkanHandle(i.InstanceHandle), i)
	i.InstanceDriver.Destroy()
}
