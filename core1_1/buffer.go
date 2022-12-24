package core1_1

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanBuffer is an implementation of the Buffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanBuffer struct {
	core1_0.Buffer
}

// PromoteBuffer accepts a Buffer object from any core version. If provided a buffer that supports
// at least core 1.1, it will return a core1_1.Buffer. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanBuffer, even if it is provided a VulkanBuffer from a higher
// core version. Two Vulkan 1.1 compatible Buffer objects with the same Buffer.Handle will
// return the same interface value when passed to this method.
func PromoteBuffer(buffer core1_0.Buffer) Buffer {
	if buffer == nil {
		return nil
	}

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
