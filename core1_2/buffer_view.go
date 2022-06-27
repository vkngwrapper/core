package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanBufferView struct {
	core1_1.BufferView
}

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
