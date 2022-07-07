package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanBuffer struct {
	core1_1.Buffer
}

func PromoteBuffer(buffer core1_0.Buffer) Buffer {
	if !buffer.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedBuffer := core1_1.PromoteBuffer(buffer)

	return buffer.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(buffer.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanBuffer{
				Buffer: promotedBuffer,
			}
		}).(Buffer)
}
