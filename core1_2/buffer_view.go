package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

// VulkanBufferView is an implementation of the BufferView interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanBufferView struct {
	core1_1.BufferView
}

// PromoteBufferView accepts a BufferView object from any core version. If provided a buffer view that supports
// at least core 1.2, it will return a core1_2.BufferView. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanBufferView, even if it is provided a VulkanBufferView from a higher
// core version. Two Vulkan 1.2 compatible BufferView objects with the same BufferView.Handle will
// return the same interface value when passed to this method.
func PromoteBufferView(bufferView core1_0.BufferView) BufferView {
	if !bufferView.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedBufferView := core1_1.PromoteBufferView(bufferView)

	return bufferView.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(bufferView.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanBufferView{promotedBufferView}
		}).(BufferView)
}
