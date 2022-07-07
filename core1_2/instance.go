package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanInstance struct {
	core1_0.Instance
}

func PromoteInstance(instance core1_0.Instance) Instance {
	if !instance.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedInstance := core1_1.PromoteInstance(instance)
	return instance.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(instance.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanInstance{
				Instance: promotedInstance,
			}
		}).(Instance)
}
