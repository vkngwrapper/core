package core1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorPool struct {
	core1_0.DescriptorPool
}

func PromoteDescriptorPool(descriptorPool core1_0.DescriptorPool) DescriptorPool {
	if !descriptorPool.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return descriptorPool.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(descriptorPool.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanDescriptorPool{descriptorPool}
		}).(DescriptorPool)
}
