package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanBuffer struct {
	core1_1.Buffer
}

func PromoteBuffer(buffer core1_0.Buffer) core1_2.Buffer {
	if !buffer.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return buffer.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(buffer.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanBuffer{
				Buffer: core1_1.PromoteBuffer(buffer),
			}
		}).(core1_2.Buffer)
}
