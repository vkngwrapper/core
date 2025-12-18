package core1_2

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanInstance is an implementation of the Instance interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanInstance struct {
	core1_0.Instance
}

// PromoteInstance accepts an Instance object from any core version. If provided an instance that supports
// at least core 1.2, it will return a core1_2.Instance. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanInstance, even if it is provided a VulkanInstance from a higher
// core version. Two Vulkan 1.2 compatible Instance objects with the same Instance.Handle will
// return the same interface value when passed to this method.
func PromoteInstance(instance core1_0.Instance) Instance {
	if instance == nil {
		return nil
	}
	if !instance.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promoted, alreadyPromoted := instance.(Instance)
	if alreadyPromoted {
		return promoted
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
