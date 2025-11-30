package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanDescriptorUpdateTemplate is an implementation of the DescriptorUpdateTemplate interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorUpdateTemplate struct {
	core1_1.DescriptorUpdateTemplate
}

// PromoteDescriptorUpdateTemplate accepts a DescriptorUpdateTemplate object from any core version. If provided a descriptor update template that supports
// at least core 1.2, it will return a core1_2.DescriptorUpdateTemplate. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanDescriptorUpdateTemplate, even if it is provided a VulkanDescriptorUpdateTemplate from a higher
// core version. Two Vulkan 1.2 compatible DescriptorUpdateTemplate objects with the same DescriptorUpdateTemplate.Handle will
// return the same interface value when passed to this method.
func PromoteDescriptorUpdateTemplate(template core1_1.DescriptorUpdateTemplate) DescriptorUpdateTemplate {
	if template == nil {
		return nil
	}
	if !template.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promoted, alreadyPromoted := template.(DescriptorUpdateTemplate)
	if alreadyPromoted {
		return promoted
	}

	return template.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(template.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanDescriptorUpdateTemplate{template}
		}).(DescriptorUpdateTemplate)
}
