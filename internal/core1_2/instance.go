package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanInstance struct {
	core1_0.Instance
}

func PromoteInstance(instance core1_0.Instance) core1_1.Instance {
	if !instance.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return instance.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(instance.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanInstance{
				Instance: core1_1.PromoteInstance(instance),
			}
		}).(core1_2.Instance)
}
