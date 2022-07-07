package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanBuffer struct {
	core1_0.Buffer
}

func PromoteBuffer(buffer core1_0.Buffer) Buffer {
	if !buffer.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return buffer.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(buffer.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanBuffer{
				Buffer: buffer,
			}
		}).(Buffer)
}
