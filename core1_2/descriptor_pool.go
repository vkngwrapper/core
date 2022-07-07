package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanDescriptorPool struct {
	core1_1.DescriptorPool
}

func PromoteDescriptorPool(descriptorPool core1_0.DescriptorPool) DescriptorPool {
	if !descriptorPool.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedDescriptorPool := core1_1.PromoteDescriptorPool(descriptorPool)

	return descriptorPool.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(descriptorPool.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanDescriptorPool{promotedDescriptorPool}
		}).(DescriptorPool)
}
