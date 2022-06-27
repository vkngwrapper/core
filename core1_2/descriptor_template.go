package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorUpdateTemplate struct {
	core1_1.DescriptorUpdateTemplate
}

func PromoteDescriptorUpdateTemplate(template core1_1.DescriptorUpdateTemplate) DescriptorSetLayout {
	if !template.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return template.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(template.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanDescriptorUpdateTemplate{template}
		}).(DescriptorSetLayout)
}
