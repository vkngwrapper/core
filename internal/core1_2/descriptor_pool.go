package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorPool struct {
	core1_1.DescriptorPool
}

func PromoteDescriptorPool(descriptorPool core1_0.DescriptorPool) core1_2.DescriptorPool {
	if !descriptorPool.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return descriptorPool.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(descriptorPool.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanDescriptorPool{core1_1.PromoteDescriptorPool(descriptorPool)}
		}).(core1_2.DescriptorPool)
}
