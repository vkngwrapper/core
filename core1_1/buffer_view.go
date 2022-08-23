package core1_1

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanBufferView is an implementation of the BufferView interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanBufferView struct {
	core1_0.BufferView
}

// PromoteBufferView accepts a BufferView object from any core version. If provided a buffer view that supports
// at least core 1.1, it will return a core1_1.BufferView. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanBufferView, even if it is provided a VulkanBufferView from a higher
// core version. Two Vulkan 1.1 compatible BufferView objects with the same BufferView.Handle will
// return the same interface value when passed to this method.
func PromoteBufferView(bufferView core1_0.BufferView) BufferView {
	if !bufferView.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return bufferView.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(bufferView.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanBufferView{bufferView}
		}).(BufferView)
}
