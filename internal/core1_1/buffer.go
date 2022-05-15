package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanBuffer struct {
	core1_0.Buffer
}

func PromoteBuffer(buffer core1_0.Buffer) core1_1.Buffer {
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
		}).(core1_1.Buffer)
}
